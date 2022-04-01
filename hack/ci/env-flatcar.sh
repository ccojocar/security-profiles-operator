# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Setup the Go environment variables
export PATH="/opt/go/bin:$PATH"
export GOPATH="/opt/go"
export GOBIN="$GOPATH/bin"

# Add the binaries to the path
export PATH="/opt/bin:$PATH"

# Added for faster debugging with lower verbosity
export KUBECONFIG=/etc/kubernetes/admin.conf
alias k=kubectl
