package priorityqueue

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/btree"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/clock"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/internal/metrics"
)

// AddOpts describes the options for adding items to the queue.
type AddOpts struct {
	After       time.Duration
	RateLimited bool
	Priority    int
}

// PriorityQueue is a priority queue for a controller. It
// internally de-duplicates all items that are added to
// it. It will use the max of the passed priorities and the
// min of possible durations.
type PriorityQueue[T comparable] interface {
	workqueue.TypedRateLimitingInterface[T]
	AddWithOpts(o AddOpts, Items ...T)
	GetWithPriority() (item T, priority int, shutdown bool)
}

// Opts contains the options for a PriorityQueue.
type Opts[T comparable] struct {
	// Ratelimiter is being used when AddRateLimited is called. Defaults to a per-item exponential backoff
	// limiter with an initial delay of five milliseconds and a max delay of 1000 seconds.
	RateLimiter    workqueue.TypedRateLimiter[T]
	MetricProvider workqueue.MetricsProvider
}

// Opt allows to configure a PriorityQueue.
type Opt[T comparable] func(*Opts[T])

// New constructs a new PriorityQueue.
func New[T comparable](name string, o ...Opt[T]) PriorityQueue[T] {
	opts := &Opts[T]{}
	for _, f := range o {
		f(opts)
	}

	if opts.RateLimiter == nil {
		opts.RateLimiter = workqueue.NewTypedItemExponentialFailureRateLimiter[T](5*time.Millisecond, 1000*time.Second)
	}

	if opts.MetricProvider == nil {
		opts.MetricProvider = metrics.WorkqueueMetricsProvider{}
	}

	pq := &priorityqueue[T]{
		items:       map[T]*item[T]{},
		queue:       btree.NewG(32, less[T]),
		becameReady: sets.Set[T]{},
		metrics:     newQueueMetrics[T](opts.MetricProvider, name, clock.RealClock{}),
		// itemOrWaiterAdded indicates that an item or
		// waiter was added. It must be buffered, because
		// if we currently process items we can't tell
		// if that included the new item/waiter.
		itemOrWaiterAdded: make(chan struct{}, 1),
		rateLimiter:       opts.RateLimiter,
		locked:            sets.Set[T]{},
		done:              make(chan struct{}),
		get:               make(chan item[T]),
		now:               time.Now,
		tick:              time.Tick,
	}

	go pq.spin()
	if _, ok := pq.metrics.(noMetrics[T]); !ok {
		go pq.updateUnfinishedWorkLoop()
	}

	return pq
}

type priorityqueue[T comparable] struct {
	// lock has to be acquired for any access any of items, queue, addedCounter
	// or becameReady
	lock  sync.Mutex
	items map[T]*item[T]
	queue bTree[*item[T]]

	// addedCounter is a counter of elements added, we need it
	// because unixNano is not guaranteed to be unique.
	addedCounter uint64

	// becameReady holds items that are in the queue, were added
	// with non-zero after and became ready. We need it to call the
	// metrics add exactly once for them.
	becameReady sets.Set[T]
	metrics     queueMetrics[T]

	itemOrWaiterAdded chan struct{}

	rateLimiter workqueue.TypedRateLimiter[T]

	// locked contains the keys we handed out through Get() and that haven't
	// yet been returned through Done().
	locked     sets.Set[T]
	lockedLock sync.RWMutex

	shutdown atomic.Bool
	done     chan struct{}

	get chan item[T]

	// waiters is the number of routines blocked in Get, we use it to determine
	// if we can push items.
	waiters atomic.Int64

	// Configurable for testing
	now  func() time.Time
	tick func(time.Duration) <-chan time.Time
}

func (w *priorityqueue[T]) AddWithOpts(o AddOpts, items ...T) {
	w.lock.Lock()
	defer w.lock.Unlock()

	for _, key := range items {
		if o.RateLimited {
			after := w.rateLimiter.When(key)
			if o.After == 0 || after < o.After {
				o.After = after
			}
		}

		var readyAt *time.Time
		if o.After > 0 {
			readyAt = ptr.To(w.now().Add(o.After))
			w.metrics.retry()
		}
		if _, ok := w.items[key]; !ok {
			item := &item[T]{
				key:          key,
				addedCounter: w.addedCounter,
				priority:     o.Priority,
				readyAt:      readyAt,
			}
			w.items[key] = item
			w.queue.ReplaceOrInsert(item)
			if item.readyAt == nil {
				w.metrics.add(key)
			}
			w.addedCounter++
			continue
		}

		// The b-tree de-duplicates based on ordering and any change here
		// will affect the order - Just delete and re-add.
		item, _ := w.queue.Delete(w.items[key])
		if o.Priority > item.priority {
			item.priority = o.Priority
		}

		if item.readyAt != nil && (readyAt == nil || readyAt.Before(*item.readyAt)) {
			if readyAt == nil {
				w.metrics.add(key)
			}
			item.readyAt = readyAt
		}

		w.queue.ReplaceOrInsert(item)
	}

	if len(items) > 0 {
		w.notifyItemOrWaiterAdded()
	}
}

func (w *priorityqueue[T]) notifyItemOrWaiterAdded() {
	select {
	case w.itemOrWaiterAdded <- struct{}{}:
	default:
	}
}

func (w *priorityqueue[T]) spin() {
	blockForever := make(chan time.Time)
	var nextReady <-chan time.Time
	nextReady = blockForever

	for {
		select {
		case <-w.done:
			return
		case <-w.itemOrWaiterAdded:
		case <-nextReady:
		}

		nextReady = blockForever

		func() {
			w.lock.Lock()
			defer w.lock.Unlock()

			w.lockedLock.Lock()
			defer w.lockedLock.Unlock()

			// manipulating the tree from within Ascend might lead to panics, so
			// track what we want to delete and do it after we are done ascending.
			var toDelete []*item[T]
			w.queue.Ascend(func(item *item[T]) bool {
				if item.readyAt != nil {
					if readyAt := item.readyAt.Sub(w.now()); readyAt > 0 {
						nextReady = w.tick(readyAt)
						return false
					}
					if !w.becameReady.Has(item.key) {
						w.metrics.add(item.key)
						w.becameReady.Insert(item.key)
					}
				}

				if w.waiters.Load() == 0 {
					// Have to keep iterating here to ensure we update metrics
					// for further items that became ready and set nextReady.
					return true
				}

				// Item is locked, we can not hand it out
				if w.locked.Has(item.key) {
					return true
				}

				w.metrics.get(item.key)
				w.locked.Insert(item.key)
				w.waiters.Add(-1)
				delete(w.items, item.key)
				toDelete = append(toDelete, item)
				w.becameReady.Delete(item.key)
				w.get <- *item

				return true
			})

			for _, item := range toDelete {
				w.queue.Delete(item)
			}
		}()
	}
}

func (w *priorityqueue[T]) Add(item T) {
	w.AddWithOpts(AddOpts{}, item)
}

func (w *priorityqueue[T]) AddAfter(item T, after time.Duration) {
	w.AddWithOpts(AddOpts{After: after}, item)
}

func (w *priorityqueue[T]) AddRateLimited(item T) {
	w.AddWithOpts(AddOpts{RateLimited: true}, item)
}

func (w *priorityqueue[T]) GetWithPriority() (_ T, priority int, shutdown bool) {
	w.waiters.Add(1)

	w.notifyItemOrWaiterAdded()
	item := <-w.get

	return item.key, item.priority, w.shutdown.Load()
}

func (w *priorityqueue[T]) Get() (item T, shutdown bool) {
	key, _, shutdown := w.GetWithPriority()
	return key, shutdown
}

func (w *priorityqueue[T]) Forget(item T) {
	w.rateLimiter.Forget(item)
}

func (w *priorityqueue[T]) NumRequeues(item T) int {
	return w.rateLimiter.NumRequeues(item)
}

func (w *priorityqueue[T]) ShuttingDown() bool {
	return w.shutdown.Load()
}

func (w *priorityqueue[T]) Done(item T) {
	w.lockedLock.Lock()
	defer w.lockedLock.Unlock()
	w.locked.Delete(item)
	w.metrics.done(item)
	w.notifyItemOrWaiterAdded()
}

func (w *priorityqueue[T]) ShutDown() {
	w.shutdown.Store(true)
	close(w.done)
}

// ShutDownWithDrain just calls ShutDown, as the draining
// functionality is not used by controller-runtime.
func (w *priorityqueue[T]) ShutDownWithDrain() {
	w.ShutDown()
}

// Len returns the number of items that are ready to be
// picked up. It does not include items that are not yet
// ready.
func (w *priorityqueue[T]) Len() int {
	w.lock.Lock()
	defer w.lock.Unlock()

	var result int
	w.queue.Ascend(func(item *item[T]) bool {
		if item.readyAt == nil || item.readyAt.Compare(w.now()) <= 0 {
			result++
			return true
		}
		return false
	})

	return result
}

func less[T comparable](a, b *item[T]) bool {
	if a.readyAt == nil && b.readyAt != nil {
		return true
	}
	if b.readyAt == nil && a.readyAt != nil {
		return false
	}
	if a.readyAt != nil && b.readyAt != nil && !a.readyAt.Equal(*b.readyAt) {
		return a.readyAt.Before(*b.readyAt)
	}
	if a.priority != b.priority {
		return a.priority > b.priority
	}

	return a.addedCounter < b.addedCounter
}

type item[T comparable] struct {
	key          T
	addedCounter uint64
	priority     int
	readyAt      *time.Time
}

func (w *priorityqueue[T]) updateUnfinishedWorkLoop() {
	t := time.NewTicker(500 * time.Millisecond) // borrowed from workqueue: https://github.com/kubernetes/kubernetes/blob/67a807bf142c7a2a5ecfdb2a5d24b4cdea4cc79c/staging/src/k8s.io/client-go/util/workqueue/queue.go#L182
	defer t.Stop()
	for range t.C {
		if w.shutdown.Load() {
			return
		}
		w.metrics.updateUnfinishedWork()
	}
}

type bTree[T any] interface {
	ReplaceOrInsert(item T) (_ T, _ bool)
	Delete(item T) (T, bool)
	Ascend(iterator btree.ItemIteratorG[T])
}