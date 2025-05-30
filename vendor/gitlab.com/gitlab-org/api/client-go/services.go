//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type (
	ServicesServiceInterface interface {
		ListServices(pid interface{}, options ...RequestOptionFunc) ([]*Service, *Response, error)
		GetCustomIssueTrackerService(pid interface{}, options ...RequestOptionFunc) (*CustomIssueTrackerService, *Response, error)
		SetCustomIssueTrackerService(pid interface{}, opt *SetCustomIssueTrackerServiceOptions, options ...RequestOptionFunc) (*CustomIssueTrackerService, *Response, error)
		DeleteCustomIssueTrackerService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetDataDogService(pid interface{}, options ...RequestOptionFunc) (*DataDogService, *Response, error)
		SetDataDogService(pid interface{}, opt *SetDataDogServiceOptions, options ...RequestOptionFunc) (*DataDogService, *Response, error)
		DeleteDataDogService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetDiscordService(pid interface{}, options ...RequestOptionFunc) (*DiscordService, *Response, error)
		SetDiscordService(pid interface{}, opt *SetDiscordServiceOptions, options ...RequestOptionFunc) (*DiscordService, *Response, error)
		DeleteDiscordService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetDroneCIService(pid interface{}, options ...RequestOptionFunc) (*DroneCIService, *Response, error)
		SetDroneCIService(pid interface{}, opt *SetDroneCIServiceOptions, options ...RequestOptionFunc) (*DroneCIService, *Response, error)
		DeleteDroneCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetEmailsOnPushService(pid interface{}, options ...RequestOptionFunc) (*EmailsOnPushService, *Response, error)
		SetEmailsOnPushService(pid interface{}, opt *SetEmailsOnPushServiceOptions, options ...RequestOptionFunc) (*EmailsOnPushService, *Response, error)
		DeleteEmailsOnPushService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetExternalWikiService(pid interface{}, options ...RequestOptionFunc) (*ExternalWikiService, *Response, error)
		SetExternalWikiService(pid interface{}, opt *SetExternalWikiServiceOptions, options ...RequestOptionFunc) (*ExternalWikiService, *Response, error)
		DeleteExternalWikiService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetGithubService(pid interface{}, options ...RequestOptionFunc) (*GithubService, *Response, error)
		SetGithubService(pid interface{}, opt *SetGithubServiceOptions, options ...RequestOptionFunc) (*GithubService, *Response, error)
		DeleteGithubService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetHarborService(pid interface{}, options ...RequestOptionFunc) (*HarborService, *Response, error)
		SetHarborService(pid interface{}, opt *SetHarborServiceOptions, options ...RequestOptionFunc) (*HarborService, *Response, error)
		DeleteHarborService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetSlackApplication(pid interface{}, options ...RequestOptionFunc) (*SlackApplication, *Response, error)
		SetSlackApplication(pid interface{}, opt *SetSlackApplicationOptions, options ...RequestOptionFunc) (*SlackApplication, *Response, error)
		DisableSlackApplication(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		SetGitLabCIService(pid interface{}, opt *SetGitLabCIServiceOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteGitLabCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		SetHipChatService(pid interface{}, opt *SetHipChatServiceOptions, options ...RequestOptionFunc) (*Response, error)
		DeleteHipChatService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetJenkinsCIService(pid interface{}, options ...RequestOptionFunc) (*JenkinsCIService, *Response, error)
		SetJenkinsCIService(pid interface{}, opt *SetJenkinsCIServiceOptions, options ...RequestOptionFunc) (*JenkinsCIService, *Response, error)
		DeleteJenkinsCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetJiraService(pid interface{}, options ...RequestOptionFunc) (*JiraService, *Response, error)
		SetJiraService(pid interface{}, opt *SetJiraServiceOptions, options ...RequestOptionFunc) (*JiraService, *Response, error)
		DeleteJiraService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetMattermostService(pid interface{}, options ...RequestOptionFunc) (*MattermostService, *Response, error)
		SetMattermostService(pid interface{}, opt *SetMattermostServiceOptions, options ...RequestOptionFunc) (*MattermostService, *Response, error)
		DeleteMattermostService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetMattermostSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*MattermostSlashCommandsService, *Response, error)
		SetMattermostSlashCommandsService(pid interface{}, opt *SetMattermostSlashCommandsServiceOptions, options ...RequestOptionFunc) (*MattermostSlashCommandsService, *Response, error)
		DeleteMattermostSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetMicrosoftTeamsService(pid interface{}, options ...RequestOptionFunc) (*MicrosoftTeamsService, *Response, error)
		SetMicrosoftTeamsService(pid interface{}, opt *SetMicrosoftTeamsServiceOptions, options ...RequestOptionFunc) (*MicrosoftTeamsService, *Response, error)
		DeleteMicrosoftTeamsService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetPipelinesEmailService(pid interface{}, options ...RequestOptionFunc) (*PipelinesEmailService, *Response, error)
		SetPipelinesEmailService(pid interface{}, opt *SetPipelinesEmailServiceOptions, options ...RequestOptionFunc) (*PipelinesEmailService, *Response, error)
		DeletePipelinesEmailService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetPrometheusService(pid interface{}, options ...RequestOptionFunc) (*PrometheusService, *Response, error)
		SetPrometheusService(pid interface{}, opt *SetPrometheusServiceOptions, options ...RequestOptionFunc) (*PrometheusService, *Response, error)
		DeletePrometheusService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetRedmineService(pid interface{}, options ...RequestOptionFunc) (*RedmineService, *Response, error)
		SetRedmineService(pid interface{}, opt *SetRedmineServiceOptions, options ...RequestOptionFunc) (*RedmineService, *Response, error)
		DeleteRedmineService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetSlackService(pid interface{}, options ...RequestOptionFunc) (*SlackService, *Response, error)
		SetSlackService(pid interface{}, opt *SetSlackServiceOptions, options ...RequestOptionFunc) (*SlackService, *Response, error)
		DeleteSlackService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetSlackSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*SlackSlashCommandsService, *Response, error)
		SetSlackSlashCommandsService(pid interface{}, opt *SetSlackSlashCommandsServiceOptions, options ...RequestOptionFunc) (*SlackSlashCommandsService, *Response, error)
		DeleteSlackSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetTelegramService(pid interface{}, options ...RequestOptionFunc) (*TelegramService, *Response, error)
		SetTelegramService(pid interface{}, opt *SetTelegramServiceOptions, options ...RequestOptionFunc) (*TelegramService, *Response, error)
		DeleteTelegramService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
		GetYouTrackService(pid interface{}, options ...RequestOptionFunc) (*YouTrackService, *Response, error)
		SetYouTrackService(pid interface{}, opt *SetYouTrackServiceOptions, options ...RequestOptionFunc) (*YouTrackService, *Response, error)
		DeleteYouTrackService(pid interface{}, options ...RequestOptionFunc) (*Response, error)
	}

	// ServicesService handles communication with the services related methods of
	// the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/integrations.html
	ServicesService struct {
		client *Client
	}
)

var _ ServicesServiceInterface = (*ServicesService)(nil)

// Service represents a GitLab service.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/integrations.html
type Service struct {
	ID                             int        `json:"id"`
	Title                          string     `json:"title"`
	Slug                           string     `json:"slug"`
	CreatedAt                      *time.Time `json:"created_at"`
	UpdatedAt                      *time.Time `json:"updated_at"`
	Active                         bool       `json:"active"`
	AlertEvents                    bool       `json:"alert_events"`
	CommitEvents                   bool       `json:"commit_events"`
	ConfidentialIssuesEvents       bool       `json:"confidential_issues_events"`
	ConfidentialNoteEvents         bool       `json:"confidential_note_events"`
	DeploymentEvents               bool       `json:"deployment_events"`
	GroupConfidentialMentionEvents bool       `json:"group_confidential_mention_events"`
	GroupMentionEvents             bool       `json:"group_mention_events"`
	IncidentEvents                 bool       `json:"incident_events"`
	IssuesEvents                   bool       `json:"issues_events"`
	JobEvents                      bool       `json:"job_events"`
	MergeRequestsEvents            bool       `json:"merge_requests_events"`
	NoteEvents                     bool       `json:"note_events"`
	PipelineEvents                 bool       `json:"pipeline_events"`
	PushEvents                     bool       `json:"push_events"`
	TagPushEvents                  bool       `json:"tag_push_events"`
	VulnerabilityEvents            bool       `json:"vulnerability_events"`
	WikiPageEvents                 bool       `json:"wiki_page_events"`
	CommentOnEventEnabled          bool       `json:"comment_on_event_enabled"`
	Inherited                      bool       `json:"inherited"`
}

// ListServices gets a list of all active services.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/integrations.html#list-all-active-integrations
func (s *ServicesService) ListServices(pid interface{}, options ...RequestOptionFunc) ([]*Service, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var svcs []*Service
	resp, err := s.client.Do(req, &svcs)
	if err != nil {
		return nil, resp, err
	}

	return svcs, resp, nil
}

// CustomIssueTrackerService represents Custom Issue Tracker service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker
type CustomIssueTrackerService struct {
	Service
	Properties *CustomIssueTrackerServiceProperties `json:"properties"`
}

// CustomIssueTrackerServiceProperties represents Custom Issue Tracker specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker
type CustomIssueTrackerServiceProperties struct {
	ProjectURL  string `json:"project_url,omitempty"`
	IssuesURL   string `json:"issues_url,omitempty"`
	NewIssueURL string `json:"new_issue_url,omitempty"`
}

// GetCustomIssueTrackerService gets Custom Issue Tracker service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-custom-issue-tracker-settings
func (s *ServicesService) GetCustomIssueTrackerService(pid interface{}, options ...RequestOptionFunc) (*CustomIssueTrackerService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/custom-issue-tracker", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(CustomIssueTrackerService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetCustomIssueTrackerServiceOptions represents the available SetCustomIssueTrackerService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-a-custom-issue-tracker
type SetCustomIssueTrackerServiceOptions struct {
	NewIssueURL *string `url:"new_issue_url,omitempty" json:"new_issue_url,omitempty"`
	IssuesURL   *string `url:"issues_url,omitempty" json:"issues_url,omitempty"`
	ProjectURL  *string `url:"project_url,omitempty" json:"project_url,omitempty"`
}

// SetCustomIssueTrackerService sets Custom Issue Tracker service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-a-custom-issue-tracker
func (s *ServicesService) SetCustomIssueTrackerService(pid interface{}, opt *SetCustomIssueTrackerServiceOptions, options ...RequestOptionFunc) (*CustomIssueTrackerService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/custom-issue-tracker", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(CustomIssueTrackerService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteCustomIssueTrackerService deletes Custom Issue Tracker service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-a-custom-issue-tracker
func (s *ServicesService) DeleteCustomIssueTrackerService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/custom-issue-tracker", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DataDogService represents DataDog service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#datadog
type DataDogService struct {
	Service
	Properties *DataDogServiceProperties `json:"properties"`
}

// DataDogServiceProperties represents DataDog specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#datadog
type DataDogServiceProperties struct {
	APIURL             string `url:"api_url,omitempty" json:"api_url,omitempty"`
	DataDogEnv         string `url:"datadog_env,omitempty" json:"datadog_env,omitempty"`
	DataDogService     string `url:"datadog_service,omitempty" json:"datadog_service,omitempty"`
	DataDogSite        string `url:"datadog_site,omitempty" json:"datadog_site,omitempty"`
	DataDogTags        string `url:"datadog_tags,omitempty" json:"datadog_tags,omitempty"`
	ArchiveTraceEvents bool   `url:"archive_trace_events,omitempty" json:"archive_trace_events,omitempty"`
}

// GetDataDogService gets DataDog service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-datadog-settings
func (s *ServicesService) GetDataDogService(pid interface{}, options ...RequestOptionFunc) (*DataDogService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/datadog", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DataDogService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetDataDogServiceOptions represents the available SetDataDogService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-datadog
type SetDataDogServiceOptions struct {
	APIKey             *string `url:"api_key,omitempty" json:"api_key,omitempty"`
	APIURL             *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	DataDogEnv         *string `url:"datadog_env,omitempty" json:"datadog_env,omitempty"`
	DataDogService     *string `url:"datadog_service,omitempty" json:"datadog_service,omitempty"`
	DataDogSite        *string `url:"datadog_site,omitempty" json:"datadog_site,omitempty"`
	DataDogTags        *string `url:"datadog_tags,omitempty" json:"datadog_tags,omitempty"`
	ArchiveTraceEvents *bool   `url:"archive_trace_events,omitempty" json:"archive_trace_events,omitempty"`
}

// SetDataDogService sets DataDog service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-datadog
func (s *ServicesService) SetDataDogService(pid interface{}, opt *SetDataDogServiceOptions, options ...RequestOptionFunc) (*DataDogService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/datadog", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DataDogService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteDataDogService deletes the DataDog service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-datadog
func (s *ServicesService) DeleteDataDogService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/datadog", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DiscordService represents Discord service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#discord-notifications
type DiscordService struct {
	Service
	Properties *DiscordServiceProperties `json:"properties"`
}

// DiscordServiceProperties represents Discord specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#discord-notifications
type DiscordServiceProperties struct {
	BranchesToBeNotified      string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	NotifyOnlyBrokenPipelines bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
}

// GetDiscordService gets Discord service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-discord-notifications-settings
func (s *ServicesService) GetDiscordService(pid interface{}, options ...RequestOptionFunc) (*DiscordService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/discord", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DiscordService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetDiscordServiceOptions represents the available SetDiscordService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-discord-notifications
type SetDiscordServiceOptions struct {
	WebHook                          *string `url:"webhook,omitempty" json:"webhook,omitempty"`
	BranchesToBeNotified             *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	ConfidentialIssuesEvents         *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	ConfidentialIssuesChannel        *string `url:"confidential_issue_channel,omitempty" json:"confidential_issue_channel,omitempty"`
	ConfidentialNoteEvents           *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	ConfidentialNoteChannel          *string `url:"confidential_note_channel,omitempty" json:"confidential_note_channel,omitempty"`
	DeploymentEvents                 *bool   `url:"deployment_events,omitempty" json:"deployment_events,omitempty"`
	DeploymentChannel                *string `url:"deployment_channel,omitempty" json:"deployment_channel,omitempty"`
	GroupConfidentialMentionsEvents  *bool   `url:"group_confidential_mentions_events,omitempty" json:"group_confidential_mentions_events,omitempty"`
	GroupConfidentialMentionsChannel *string `url:"group_confidential_mentions_channel,omitempty" json:"group_confidential_mentions_channel,omitempty"`
	GroupMentionsEvents              *bool   `url:"group_mentions_events,omitempty" json:"group_mentions_events,omitempty"`
	GroupMentionsChannel             *string `url:"group_mentions_channel,omitempty" json:"group_mentions_channel,omitempty"`
	IssuesEvents                     *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	IssueChannel                     *string `url:"issue_channel,omitempty" json:"issue_channel,omitempty"`
	MergeRequestsEvents              *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	MergeRequestChannel              *string `url:"merge_request_channel,omitempty" json:"merge_request_channel,omitempty"`
	NoteEvents                       *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	NoteChannel                      *string `url:"note_channel,omitempty" json:"note_channel,omitempty"`
	NotifyOnlyBrokenPipelines        *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	PipelineEvents                   *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	PipelineChannel                  *string `url:"pipeline_channel,omitempty" json:"pipeline_channel,omitempty"`
	PushEvents                       *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	PushChannel                      *string `url:"push_channel,omitempty" json:"push_channel,omitempty"`
	TagPushEvents                    *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	TagPushChannel                   *string `url:"tag_push_channel,omitempty" json:"tag_push_channel,omitempty"`
	WikiPageEvents                   *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
	WikiPageChannel                  *string `url:"wiki_page_channel,omitempty" json:"wiki_page_channel,omitempty"`
}

// SetDiscordService sets Discord service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-discord-notifications
func (s *ServicesService) SetDiscordService(pid interface{}, opt *SetDiscordServiceOptions, options ...RequestOptionFunc) (*DiscordService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/discord", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DiscordService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// DeleteDiscordService deletes Discord service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-discord-notifications
func (s *ServicesService) DeleteDiscordService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/discord", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DroneCIService represents Drone CI service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#drone
type DroneCIService struct {
	Service
	Properties *DroneCIServiceProperties `json:"properties"`
}

// DroneCIServiceProperties represents Drone CI specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#drone
type DroneCIServiceProperties struct {
	DroneURL              string `json:"drone_url"`
	EnableSSLVerification bool   `json:"enable_ssl_verification"`
}

// GetDroneCIService gets Drone CI service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-drone-settings
func (s *ServicesService) GetDroneCIService(pid interface{}, options ...RequestOptionFunc) (*DroneCIService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/drone-ci", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DroneCIService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetDroneCIServiceOptions represents the available SetDroneCIService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-drone
type SetDroneCIServiceOptions struct {
	Token                 *string `url:"token,omitempty" json:"token,omitempty"`
	DroneURL              *string `url:"drone_url,omitempty" json:"drone_url,omitempty"`
	EnableSSLVerification *bool   `url:"enable_ssl_verification,omitempty" json:"enable_ssl_verification,omitempty"`
	PushEvents            *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	MergeRequestsEvents   *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	TagPushEvents         *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
}

// SetDroneCIService sets Drone CI service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-drone
func (s *ServicesService) SetDroneCIService(pid interface{}, opt *SetDroneCIServiceOptions, options ...RequestOptionFunc) (*DroneCIService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/drone-ci", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(DroneCIService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteDroneCIService deletes Drone CI service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-drone
func (s *ServicesService) DeleteDroneCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/drone-ci", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// EmailsOnPushService represents Emails on Push service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#emails-on-push
type EmailsOnPushService struct {
	Service
	Properties *EmailsOnPushServiceProperties `json:"properties"`
}

// EmailsOnPushServiceProperties represents Emails on Push specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#emails-on-push
type EmailsOnPushServiceProperties struct {
	Recipients             string `json:"recipients"`
	DisableDiffs           bool   `json:"disable_diffs"`
	SendFromCommitterEmail bool   `json:"send_from_committer_email"`
	PushEvents             bool   `json:"push_events"`
	TagPushEvents          bool   `json:"tag_push_events"`
	BranchesToBeNotified   string `json:"branches_to_be_notified"`
}

// GetEmailsOnPushService gets Emails on Push service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-emails-on-push-integration-settings
func (s *ServicesService) GetEmailsOnPushService(pid interface{}, options ...RequestOptionFunc) (*EmailsOnPushService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/emails-on-push", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(EmailsOnPushService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetEmailsOnPushServiceOptions represents the available SetEmailsOnPushService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-emails-on-push
type SetEmailsOnPushServiceOptions struct {
	Recipients             *string `url:"recipients,omitempty" json:"recipients,omitempty"`
	DisableDiffs           *bool   `url:"disable_diffs,omitempty" json:"disable_diffs,omitempty"`
	SendFromCommitterEmail *bool   `url:"send_from_committer_email,omitempty" json:"send_from_committer_email,omitempty"`
	PushEvents             *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	TagPushEvents          *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	BranchesToBeNotified   *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
}

// SetEmailsOnPushService sets Emails on Push service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-emails-on-push
func (s *ServicesService) SetEmailsOnPushService(pid interface{}, opt *SetEmailsOnPushServiceOptions, options ...RequestOptionFunc) (*EmailsOnPushService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/emails-on-push", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(EmailsOnPushService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteEmailsOnPushService deletes Emails on Push service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-emails-on-push
func (s *ServicesService) DeleteEmailsOnPushService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/emails-on-push", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ExternalWikiService represents External Wiki service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#external-wiki
type ExternalWikiService struct {
	Service
	Properties *ExternalWikiServiceProperties `json:"properties"`
}

// ExternalWikiServiceProperties represents External Wiki specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#external-wiki
type ExternalWikiServiceProperties struct {
	ExternalWikiURL string `json:"external_wiki_url"`
}

// GetExternalWikiService gets External Wiki service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-external-wiki-settings
func (s *ServicesService) GetExternalWikiService(pid interface{}, options ...RequestOptionFunc) (*ExternalWikiService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/external-wiki", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(ExternalWikiService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetExternalWikiServiceOptions represents the available SetExternalWikiService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-an-external-wiki
type SetExternalWikiServiceOptions struct {
	ExternalWikiURL *string `url:"external_wiki_url,omitempty" json:"external_wiki_url,omitempty"`
}

// SetExternalWikiService sets External Wiki service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-an-external-wiki
func (s *ServicesService) SetExternalWikiService(pid interface{}, opt *SetExternalWikiServiceOptions, options ...RequestOptionFunc) (*ExternalWikiService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/external-wiki", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(ExternalWikiService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteExternalWikiService deletes External Wiki service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-an-external-wiki
func (s *ServicesService) DeleteExternalWikiService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/external-wiki", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// GithubService represents Github service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#github
type GithubService struct {
	Service
	Properties *GithubServiceProperties `json:"properties"`
}

// GithubServiceProperties represents Github specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#github
type GithubServiceProperties struct {
	RepositoryURL string `json:"repository_url"`
	StaticContext bool   `json:"static_context"`
}

// GetGithubService gets Github service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-github-settings
func (s *ServicesService) GetGithubService(pid interface{}, options ...RequestOptionFunc) (*GithubService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/github", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(GithubService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetGithubServiceOptions represents the available SetGithubService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-github
type SetGithubServiceOptions struct {
	Token         *string `url:"token,omitempty" json:"token,omitempty"`
	RepositoryURL *string `url:"repository_url,omitempty" json:"repository_url,omitempty"`
	StaticContext *bool   `url:"static_context,omitempty" json:"static_context,omitempty"`
}

// SetGithubService sets Github service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-github
func (s *ServicesService) SetGithubService(pid interface{}, opt *SetGithubServiceOptions, options ...RequestOptionFunc) (*GithubService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/github", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(GithubService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteGithubService deletes Github service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-github
func (s *ServicesService) DeleteGithubService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/github", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// HarborService represents the Harbor service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#harbor
type HarborService struct {
	Service
	Properties *HarborServiceProperties `json:"properties"`
}

// HarborServiceProperties represents Harbor specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#harbor
type HarborServiceProperties struct {
	URL                  string `json:"url"`
	ProjectName          string `json:"project_name"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	UseInheritedSettings bool   `json:"use_inherited_settings"`
}

// GetHarborService gets Harbor service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-harbor-settings
func (s *ServicesService) GetHarborService(pid interface{}, options ...RequestOptionFunc) (*HarborService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/harbor", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(HarborService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetHarborServiceOptions represents the available SetHarborService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-harbor
type SetHarborServiceOptions struct {
	URL                  *string `url:"url,omitempty" json:"url,omitempty"`
	ProjectName          *string `url:"project_name,omitempty" json:"project_name,omitempty"`
	Username             *string `url:"username,omitempty" json:"username,omitempty"`
	Password             *string `url:"password,omitempty" json:"password,omitempty"`
	UseInheritedSettings *bool   `url:"use_inherited_settings,omitempty" json:"use_inherited_settings,omitempty"`
}

// SetHarborService sets Harbor service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-harbor
func (s *ServicesService) SetHarborService(pid interface{}, opt *SetHarborServiceOptions, options ...RequestOptionFunc) (*HarborService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/harbor", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(HarborService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteHarborService deletes Harbor service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-harbor
func (s *ServicesService) DeleteHarborService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/harbor", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SlackApplication represents GitLab for slack application settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#gitlab-for-slack-app
type SlackApplication struct {
	Service
	Properties *SlackApplicationProperties `json:"properties"`
}

// SlackApplicationProperties represents GitLab for slack application specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#gitlab-for-slack-app
type SlackApplicationProperties struct {
	Channel                    string `json:"channel"`
	NotifyOnlyBrokenPipelines  bool   `json:"notify_only_broken_pipelines"`
	BranchesToBeNotified       string `json:"branches_to_be_notified"`
	LabelsToBeNotified         string `json:"labels_to_be_notified"`
	LabelsToBeNotifiedBehavior string `json:"labels_to_be_notified_behavior"`
	PushChannel                string `json:"push_channel"`
	IssueChannel               string `json:"issue_channel"`
	ConfidentialIssueChannel   string `json:"confidential_issue_channel"`
	MergeRequestChannel        string `json:"merge_request_channel"`
	NoteChannel                string `json:"note_channel"`
	ConfidentialNoteChannel    string `json:"confidential_note_channel"`
	TagPushChannel             string `json:"tag_push_channel"`
	PipelineChannel            string `json:"pipeline_channel"`
	WikiPageChannel            string `json:"wiki_page_channel"`
	DeploymentChannel          string `json:"deployment_channel"`
	IncidentChannel            string `json:"incident_channel"`
	VulnerabilityChannel       string `json:"vulnerability_channel"`
	AlertChannel               string `json:"alert_channel"`

	// Deprecated: This parameter has been replaced with BranchesToBeNotified.
	NotifyOnlyDefaultBranch bool `json:"notify_only_default_branch"`
}

// GetSlackApplication gets the GitLab for Slack app integration settings for a
// project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-gitlab-for-slack-app-settings
func (s *ServicesService) GetSlackApplication(pid interface{}, options ...RequestOptionFunc) (*SlackApplication, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/gitlab-slack-application", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackApplication)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetSlackApplicationOptions represents the available SetSlackApplication()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-gitlab-for-slack-app
type SetSlackApplicationOptions struct {
	Channel                    *string `url:"channel,omitempty" json:"channel,omitempty"`
	NotifyOnlyBrokenPipelines  *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	BranchesToBeNotified       *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	AlertEvents                *bool   `url:"alert_events,omitempty" json:"alert_events,omitempty"`
	IssuesEvents               *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	ConfidentialIssuesEvents   *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	MergeRequestsEvents        *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	NoteEvents                 *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	ConfidentialNoteEvents     *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	DeploymentEvents           *bool   `url:"deployment_events,omitempty" json:"deployment_events,omitempty"`
	IncidentsEvents            *bool   `url:"incidents_events,omitempty" json:"incidents_events,omitempty"`
	PipelineEvents             *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	PushEvents                 *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	TagPushEvents              *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	VulnerabilityEvents        *bool   `url:"vulnerability_events,omitempty" json:"vulnerability_events,omitempty"`
	WikiPageEvents             *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
	LabelsToBeNotified         *string `url:"labels_to_be_notified,omitempty" json:"labels_to_be_notified,omitempty"`
	LabelsToBeNotifiedBehavior *string `url:"labels_to_be_notified_behavior,omitempty" json:"labels_to_be_notified_behavior,omitempty"`
	PushChannel                *string `url:"push_channel,omitempty" json:"push_channel,omitempty"`
	IssueChannel               *string `url:"issue_channel,omitempty" json:"issue_channel,omitempty"`
	ConfidentialIssueChannel   *string `url:"confidential_issue_channel,omitempty" json:"confidential_issue_channel,omitempty"`
	MergeRequestChannel        *string `url:"merge_request_channel,omitempty" json:"merge_request_channel,omitempty"`
	NoteChannel                *string `url:"note_channel,omitempty" json:"note_channel,omitempty"`
	ConfidentialNoteChannel    *string `url:"confidential_note_channel,omitempty" json:"confidential_note_channel,omitempty"`
	TagPushChannel             *string `url:"tag_push_channel,omitempty" json:"tag_push_channel,omitempty"`
	PipelineChannel            *string `url:"pipeline_channel,omitempty" json:"pipeline_channel,omitempty"`
	WikiPageChannel            *string `url:"wiki_page_channel,omitempty" json:"wiki_page_channel,omitempty"`
	DeploymentChannel          *string `url:"deployment_channel,omitempty" json:"deployment_channel,omitempty"`
	IncidentChannel            *string `url:"incident_channel,omitempty" json:"incident_channel,omitempty"`
	VulnerabilityChannel       *string `url:"vulnerability_channel,omitempty" json:"vulnerability_channel,omitempty"`
	AlertChannel               *string `url:"alert_channel,omitempty" json:"alert_channel,omitempty"`
	UseInheritedSettings       *bool   `url:"use_inherited_settings,omitempty" json:"use_inherited_settings,omitempty"`

	// Deprecated: This parameter has been replaced with BranchesToBeNotified.
	NotifyOnlyDefaultBranch *bool `url:"notify_only_default_branch,omitempty" json:"notify_only_default_branch,omitempty"`
}

// SetSlackApplication update the GitLab for Slack app integration for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-gitlab-for-slack-app
func (s *ServicesService) SetSlackApplication(pid interface{}, opt *SetSlackApplicationOptions, options ...RequestOptionFunc) (*SlackApplication, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/gitlab-slack-application", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackApplication)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DisableSlackApplication disable the GitLab for Slack app integration for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-gitlab-for-slack-app
func (s *ServicesService) DisableSlackApplication(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/gitlab-slack-application", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SetGitLabCIServiceOptions represents the available SetGitLabCIService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-gitlab-ci-service
type SetGitLabCIServiceOptions struct {
	Token      *string `url:"token,omitempty" json:"token,omitempty"`
	ProjectURL *string `url:"project_url,omitempty" json:"project_url,omitempty"`
}

// SetGitLabCIService sets GitLab CI service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-gitlab-ci-service
func (s *ServicesService) SetGitLabCIService(pid interface{}, opt *SetGitLabCIServiceOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/gitlab-ci", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteGitLabCIService deletes GitLab CI service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-gitlab-ci-service
func (s *ServicesService) DeleteGitLabCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/gitlab-ci", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SetHipChatServiceOptions represents the available SetHipChatService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-hipchat-service
type SetHipChatServiceOptions struct {
	Token *string `url:"token,omitempty" json:"token,omitempty" `
	Room  *string `url:"room,omitempty" json:"room,omitempty"`
}

// SetHipChatService sets HipChat service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-hipchat-service
func (s *ServicesService) SetHipChatService(pid interface{}, opt *SetHipChatServiceOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/hipchat", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteHipChatService deletes HipChat service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-hipchat-service
func (s *ServicesService) DeleteHipChatService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/hipchat", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// JenkinsCIService represents Jenkins CI service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#jenkins
type JenkinsCIService struct {
	Service
	Properties *JenkinsCIServiceProperties `json:"properties"`
}

// JenkinsCIServiceProperties represents Jenkins CI specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#jenkins
type JenkinsCIServiceProperties struct {
	URL                   string `json:"jenkins_url"`
	EnableSSLVerification bool   `json:"enable_ssl_verification"`
	ProjectName           string `json:"project_name"`
	Username              string `json:"username"`
}

// GetJenkinsCIService gets Jenkins CI service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-jenkins-settings
func (s *ServicesService) GetJenkinsCIService(pid interface{}, options ...RequestOptionFunc) (*JenkinsCIService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/jenkins", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(JenkinsCIService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetJenkinsCIServiceOptions represents the available SetJenkinsCIService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#jenkins
type SetJenkinsCIServiceOptions struct {
	URL                   *string `url:"jenkins_url,omitempty" json:"jenkins_url,omitempty"`
	EnableSSLVerification *bool   `url:"enable_ssl_verification,omitempty" json:"enable_ssl_verification,omitempty"`
	ProjectName           *string `url:"project_name,omitempty" json:"project_name,omitempty"`
	Username              *string `url:"username,omitempty" json:"username,omitempty"`
	Password              *string `url:"password,omitempty" json:"password,omitempty"`
	PushEvents            *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	MergeRequestsEvents   *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	TagPushEvents         *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
}

// SetJenkinsCIService sets Jenkins service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-jenkins
func (s *ServicesService) SetJenkinsCIService(pid interface{}, opt *SetJenkinsCIServiceOptions, options ...RequestOptionFunc) (*JenkinsCIService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/jenkins", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(JenkinsCIService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteJenkinsCIService deletes Jenkins CI service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-jenkins
func (s *ServicesService) DeleteJenkinsCIService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/jenkins", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// JiraService represents Jira service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#jira
type JiraService struct {
	Service
	Properties *JiraServiceProperties `json:"properties"`
}

// JiraServiceProperties represents Jira specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#jira
type JiraServiceProperties struct {
	URL                          string   `json:"url"`
	APIURL                       string   `json:"api_url"`
	Username                     string   `json:"username" `
	Password                     string   `json:"password" `
	Active                       bool     `json:"active"`
	JiraAuthType                 int      `json:"jira_auth_type"`
	JiraIssuePrefix              string   `json:"jira_issue_prefix"`
	JiraIssueRegex               string   `json:"jira_issue_regex"`
	JiraIssueTransitionAutomatic bool     `json:"jira_issue_transition_automatic"`
	JiraIssueTransitionID        string   `json:"jira_issue_transition_id"`
	CommitEvents                 bool     `json:"commit_events"`
	MergeRequestsEvents          bool     `json:"merge_requests_events"`
	CommentOnEventEnabled        bool     `json:"comment_on_event_enabled"`
	IssuesEnabled                bool     `json:"issues_enabled"`
	ProjectKeys                  []string `json:"project_keys" `
	UseInheritedSettings         bool     `json:"use_inherited_settings"`

	// Deprecated: This parameter was removed in GitLab 17.0
	ProjectKey string `json:"project_key" `
}

// UnmarshalJSON decodes the Jira Service Properties.
//
// This allows support of JiraIssueTransitionID for both type string (>11.9) and float64 (<11.9)
func (p *JiraServiceProperties) UnmarshalJSON(b []byte) error {
	type Alias JiraServiceProperties
	raw := struct {
		*Alias
		JiraIssueTransitionID interface{} `json:"jira_issue_transition_id"`
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	switch id := raw.JiraIssueTransitionID.(type) {
	case nil:
		// No action needed.
	case string:
		p.JiraIssueTransitionID = id
	case float64:
		p.JiraIssueTransitionID = strconv.Itoa(int(id))
	default:
		return fmt.Errorf("failed to unmarshal JiraTransitionID of type: %T", id)
	}

	return nil
}

// GetJiraService gets Jira service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-jira-service-settings
func (s *ServicesService) GetJiraService(pid interface{}, options ...RequestOptionFunc) (*JiraService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/jira", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(JiraService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetJiraServiceOptions represents the available SetJiraService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-jira-service
type SetJiraServiceOptions struct {
	URL                          *string   `url:"url,omitempty" json:"url,omitempty"`
	APIURL                       *string   `url:"api_url,omitempty" json:"api_url,omitempty"`
	Username                     *string   `url:"username,omitempty" json:"username,omitempty" `
	Password                     *string   `url:"password,omitempty" json:"password,omitempty" `
	Active                       *bool     `url:"active,omitempty" json:"active,omitempty"`
	JiraAuthType                 *int      `url:"jira_auth_type,omitempty" json:"jira_auth_type,omitempty"`
	JiraIssuePrefix              *string   `url:"jira_issue_prefix,omitempty" json:"jira_issue_prefix,omitempty"`
	JiraIssueRegex               *string   `url:"jira_issue_regex,omitempty" json:"jira_issue_regex,omitempty"`
	JiraIssueTransitionAutomatic *bool     `url:"jira_issue_transition_automatic,omitempty" json:"jira_issue_transition_automatic,omitempty"`
	JiraIssueTransitionID        *string   `url:"jira_issue_transition_id,omitempty" json:"jira_issue_transition_id,omitempty"`
	CommitEvents                 *bool     `url:"commit_events,omitempty" json:"commit_events,omitempty"`
	MergeRequestsEvents          *bool     `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	CommentOnEventEnabled        *bool     `url:"comment_on_event_enabled,omitempty" json:"comment_on_event_enabled,omitempty"`
	IssuesEnabled                *bool     `url:"issues_enabled,omitempty" json:"issues_enabled,omitempty"`
	ProjectKeys                  *[]string `url:"project_keys,comma,omitempty" json:"project_keys,omitempty" `
	UseInheritedSettings         *bool     `url:"use_inherited_settings,omitempty" json:"use_inherited_settings,omitempty"`

	// Deprecated: This parameter was removed in GitLab 17.0
	ProjectKey *string `url:"project_key,omitempty" json:"project_key,omitempty" `
}

// SetJiraService sets Jira service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-jira-service
func (s *ServicesService) SetJiraService(pid interface{}, opt *SetJiraServiceOptions, options ...RequestOptionFunc) (*JiraService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/jira", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(JiraService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteJiraService deletes Jira service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-jira-service
func (s *ServicesService) DeleteJiraService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/jira", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// MattermostService represents Mattermost service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#mattermost-notifications
type MattermostService struct {
	Service
	Properties *MattermostServiceProperties `json:"properties"`
}

// MattermostServiceProperties represents Mattermost specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#mattermost-notifications
type MattermostServiceProperties struct {
	WebHook                   string    `json:"webhook"`
	Username                  string    `json:"username"`
	Channel                   string    `json:"channel"`
	NotifyOnlyBrokenPipelines BoolValue `json:"notify_only_broken_pipelines"`
	BranchesToBeNotified      string    `json:"branches_to_be_notified"`
	ConfidentialIssueChannel  string    `json:"confidential_issue_channel"`
	ConfidentialNoteChannel   string    `json:"confidential_note_channel"`
	IssueChannel              string    `json:"issue_channel"`
	MergeRequestChannel       string    `json:"merge_request_channel"`
	NoteChannel               string    `json:"note_channel"`
	TagPushChannel            string    `json:"tag_push_channel"`
	PipelineChannel           string    `json:"pipeline_channel"`
	PushChannel               string    `json:"push_channel"`
	VulnerabilityChannel      string    `json:"vulnerability_channel"`
	WikiPageChannel           string    `json:"wiki_page_channel"`
}

// GetMattermostService gets Mattermost service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-slack-service-settings
func (s *ServicesService) GetMattermostService(pid interface{}, options ...RequestOptionFunc) (*MattermostService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MattermostService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetMattermostServiceOptions represents the available SetMattermostService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-mattermost-notifications-service
type SetMattermostServiceOptions struct {
	WebHook                   *string `url:"webhook,omitempty" json:"webhook,omitempty"`
	Username                  *string `url:"username,omitempty" json:"username,omitempty"`
	Channel                   *string `url:"channel,omitempty" json:"channel,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	PushEvents                *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	IssuesEvents              *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	ConfidentialIssuesEvents  *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	MergeRequestsEvents       *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	TagPushEvents             *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	NoteEvents                *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	ConfidentialNoteChannel   *string `url:"confidential_note_channel,omitempty" json:"confidential_note_channel,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	WikiPageEvents            *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
	PushChannel               *string `url:"push_channel,omitempty" json:"push_channel,omitempty"`
	IssueChannel              *string `url:"issue_channel,omitempty" json:"issue_channel,omitempty"`
	ConfidentialIssueChannel  *string `url:"confidential_issue_channel,omitempty" json:"confidential_issue_channel,omitempty"`
	MergeRequestChannel       *string `url:"merge_request_channel,omitempty" json:"merge_request_channel,omitempty"`
	NoteChannel               *string `url:"note_channel,omitempty" json:"note_channel,omitempty"`
	ConfidentialNoteEvents    *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	TagPushChannel            *string `url:"tag_push_channel,omitempty" json:"tag_push_channel,omitempty"`
	PipelineChannel           *string `url:"pipeline_channel,omitempty" json:"pipeline_channel,omitempty"`
	WikiPageChannel           *string `url:"wiki_page_channel,omitempty" json:"wiki_page_channel,omitempty"`
}

// SetMattermostService sets Mattermost service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-mattermost-notifications-service
func (s *ServicesService) SetMattermostService(pid interface{}, opt *SetMattermostServiceOptions, options ...RequestOptionFunc) (*MattermostService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MattermostService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteMattermostService deletes Mattermost service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-mattermost-notifications-service
func (s *ServicesService) DeleteMattermostService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// MattermostSlashCommandsService represents Mattermost slash commands settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#mattermost-slash-commands
type MattermostSlashCommandsService struct {
	Service
	Properties *MattermostSlashCommandsProperties `json:"properties"`
}

// MattermostSlashCommandsProperties represents Mattermost slash commands specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#mattermost-slash-commands
type MattermostSlashCommandsProperties struct {
	Token    string `json:"token"`
	Username string `json:"username,omitempty"`
}

// GetMattermostSlashCommandsService gets Slack Mattermost commands service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-mattermost-slash-command-integration-settings
func (s *ServicesService) GetMattermostSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*MattermostSlashCommandsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MattermostSlashCommandsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetMattermostSlashCommandsServiceOptions represents the available SetSlackSlashCommandsService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-mattermost-slash-command-integration-settings
type SetMattermostSlashCommandsServiceOptions struct {
	Token    *string `url:"token,omitempty" json:"token,omitempty"`
	Username *string `url:"username,omitempty" json:"username,omitempty"`
}

// SetMattermostSlashCommandsService sets Mattermost slash commands service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-mattermost-slash-command-integration
func (s *ServicesService) SetMattermostSlashCommandsService(pid interface{}, opt *SetMattermostSlashCommandsServiceOptions, options ...RequestOptionFunc) (*MattermostSlashCommandsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MattermostSlashCommandsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteMattermostSlashCommandsService deletes Mattermost slash commands service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-mattermost-slash-command-integration
func (s *ServicesService) DeleteMattermostSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/mattermost-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// MicrosoftTeamsService represents Microsoft Teams service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#microsoft-teams
type MicrosoftTeamsService struct {
	Service
	Properties *MicrosoftTeamsServiceProperties `json:"properties"`
}

// MicrosoftTeamsServiceProperties represents Microsoft Teams specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#microsoft-teams
type MicrosoftTeamsServiceProperties struct {
	WebHook                   string    `json:"webhook"`
	NotifyOnlyBrokenPipelines BoolValue `json:"notify_only_broken_pipelines"`
	BranchesToBeNotified      string    `json:"branches_to_be_notified"`
	IssuesEvents              BoolValue `json:"issues_events"`
	ConfidentialIssuesEvents  BoolValue `json:"confidential_issues_events"`
	MergeRequestsEvents       BoolValue `json:"merge_requests_events"`
	TagPushEvents             BoolValue `json:"tag_push_events"`
	NoteEvents                BoolValue `json:"note_events"`
	ConfidentialNoteEvents    BoolValue `json:"confidential_note_events"`
	PipelineEvents            BoolValue `json:"pipeline_events"`
	WikiPageEvents            BoolValue `json:"wiki_page_events"`
}

// GetMicrosoftTeamsService gets MicrosoftTeams service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-microsoft-teams-service-settings
func (s *ServicesService) GetMicrosoftTeamsService(pid interface{}, options ...RequestOptionFunc) (*MicrosoftTeamsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/microsoft-teams", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MicrosoftTeamsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetMicrosoftTeamsServiceOptions represents the available SetMicrosoftTeamsService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#create-edit-microsoft-teams-service
type SetMicrosoftTeamsServiceOptions struct {
	WebHook                   *string `url:"webhook,omitempty" json:"webhook,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	PushEvents                *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	IssuesEvents              *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	ConfidentialIssuesEvents  *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	MergeRequestsEvents       *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	TagPushEvents             *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	NoteEvents                *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	ConfidentialNoteEvents    *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	WikiPageEvents            *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
}

// SetMicrosoftTeamsService sets Microsoft Teams service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#create-edit-microsoft-teams-service
func (s *ServicesService) SetMicrosoftTeamsService(pid interface{}, opt *SetMicrosoftTeamsServiceOptions, options ...RequestOptionFunc) (*MicrosoftTeamsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/microsoft-teams", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(MicrosoftTeamsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteMicrosoftTeamsService deletes Microsoft Teams service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-microsoft-teams-service
func (s *ServicesService) DeleteMicrosoftTeamsService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/microsoft-teams", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// PipelinesEmailService represents Pipelines Email service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails
type PipelinesEmailService struct {
	Service
	Properties *PipelinesEmailProperties `json:"properties"`
}

// PipelinesEmailProperties represents PipelinesEmail specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails
type PipelinesEmailProperties struct {
	Recipients                string    `json:"recipients"`
	NotifyOnlyBrokenPipelines BoolValue `json:"notify_only_broken_pipelines"`
	NotifyOnlyDefaultBranch   BoolValue `json:"notify_only_default_branch"`
	BranchesToBeNotified      string    `json:"branches_to_be_notified"`
}

// GetPipelinesEmailService gets Pipelines Email service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-pipeline-emails-service-settings
func (s *ServicesService) GetPipelinesEmailService(pid interface{}, options ...RequestOptionFunc) (*PipelinesEmailService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(PipelinesEmailService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetPipelinesEmailServiceOptions represents the available
// SetPipelinesEmailService() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails
type SetPipelinesEmailServiceOptions struct {
	Recipients                *string `url:"recipients,omitempty" json:"recipients,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	NotifyOnlyDefaultBranch   *bool   `url:"notify_only_default_branch,omitempty" json:"notify_only_default_branch,omitempty"`
	AddPusher                 *bool   `url:"add_pusher,omitempty" json:"add_pusher,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
}

// SetPipelinesEmailService sets Pipelines Email service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails
func (s *ServicesService) SetPipelinesEmailService(pid interface{}, opt *SetPipelinesEmailServiceOptions, options ...RequestOptionFunc) (*PipelinesEmailService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(PipelinesEmailService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeletePipelinesEmailService deletes Pipelines Email service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-pipeline-emails-service
func (s *ServicesService) DeletePipelinesEmailService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// PrometheusService represents Prometheus service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#prometheus
type PrometheusService struct {
	Service
	Properties *PrometheusServiceProperties `json:"properties"`
}

// PrometheusServiceProperties represents Prometheus specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#prometheus
type PrometheusServiceProperties struct {
	APIURL                      string `json:"api_url"`
	GoogleIAPAudienceClientID   string `json:"google_iap_audience_client_id"`
	GoogleIAPServiceAccountJSON string `json:"google_iap_service_account_json"`
}

// GetPrometheusService gets Prometheus service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-prometheus-service-settings
func (s *ServicesService) GetPrometheusService(pid interface{}, options ...RequestOptionFunc) (*PrometheusService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/prometheus", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(PrometheusService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetPrometheusServiceOptions represents the available SetPrometheusService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-prometheus-service
type SetPrometheusServiceOptions struct {
	APIURL                      *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	GoogleIAPAudienceClientID   *string `url:"google_iap_audience_client_id,omitempty" json:"google_iap_audience_client_id,omitempty"`
	GoogleIAPServiceAccountJSON *string `url:"google_iap_service_account_json,omitempty" json:"google_iap_service_account_json,omitempty"`
}

// SetPrometheusService sets Prometheus service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-prometheus-service
func (s *ServicesService) SetPrometheusService(pid interface{}, opt *SetPrometheusServiceOptions, options ...RequestOptionFunc) (*PrometheusService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/prometheus", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(PrometheusService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeletePrometheusService deletes Prometheus service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-prometheus-service
func (s *ServicesService) DeletePrometheusService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/prometheus", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RedmineService represents the Redmine service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#redmine
type RedmineService struct {
	Service
	Properties *RedmineServiceProperties `json:"properties"`
}

// RedmineServiceProperties represents Redmine specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#redmine
type RedmineServiceProperties struct {
	NewIssueURL          string    `json:"new_issue_url"`
	ProjectURL           string    `json:"project_url"`
	IssuesURL            string    `json:"issues_url"`
	UseInheritedSettings BoolValue `json:"use_inherited_settings"`
}

// GetRedmineService gets Redmine service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-redmine-settings
func (s *ServicesService) GetRedmineService(pid interface{}, options ...RequestOptionFunc) (*RedmineService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/redmine", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(RedmineService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetRedmineServiceOptions represents the available SetRedmineService().
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-redmine
type SetRedmineServiceOptions struct {
	NewIssueURL          *string `url:"new_issue_url,omitempty" json:"new_issue_url,omitempty"`
	ProjectURL           *string `url:"project_url,omitempty" json:"project_url,omitempty"`
	IssuesURL            *string `url:"issues_url,omitempty" json:"issues_url,omitempty"`
	UseInheritedSettings *bool   `url:"use_inherited_settings,omitempty" json:"use_inherited_settings,omitempty"`
}

// SetRedmineService sets Redmine service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-redmine
func (s *ServicesService) SetRedmineService(pid interface{}, opt *SetRedmineServiceOptions, options ...RequestOptionFunc) (*RedmineService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/redmine", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(RedmineService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteRedmineService deletes Redmine service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-redmine
func (s *ServicesService) DeleteRedmineService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/integrations/redmine", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SlackService represents Slack service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#slack
type SlackService struct {
	Service
	Properties *SlackServiceProperties `json:"properties"`
}

// SlackServiceProperties represents Slack specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#slack
type SlackServiceProperties struct {
	WebHook                   string    `json:"webhook"`
	Username                  string    `json:"username"`
	Channel                   string    `json:"channel"`
	NotifyOnlyBrokenPipelines BoolValue `json:"notify_only_broken_pipelines"`
	NotifyOnlyDefaultBranch   BoolValue `json:"notify_only_default_branch"`
	BranchesToBeNotified      string    `json:"branches_to_be_notified"`
	AlertChannel              string    `json:"alert_channel"`
	ConfidentialIssueChannel  string    `json:"confidential_issue_channel"`
	ConfidentialNoteChannel   string    `json:"confidential_note_channel"`
	DeploymentChannel         string    `json:"deployment_channel"`
	IssueChannel              string    `json:"issue_channel"`
	MergeRequestChannel       string    `json:"merge_request_channel"`
	NoteChannel               string    `json:"note_channel"`
	TagPushChannel            string    `json:"tag_push_channel"`
	PipelineChannel           string    `json:"pipeline_channel"`
	PushChannel               string    `json:"push_channel"`
	VulnerabilityChannel      string    `json:"vulnerability_channel"`
	WikiPageChannel           string    `json:"wiki_page_channel"`
}

// GetSlackService gets Slack service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-slack-service-settings
func (s *ServicesService) GetSlackService(pid interface{}, options ...RequestOptionFunc) (*SlackService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetSlackServiceOptions represents the available SetSlackService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-slack-service
type SetSlackServiceOptions struct {
	WebHook                   *string `url:"webhook,omitempty" json:"webhook,omitempty"`
	Username                  *string `url:"username,omitempty" json:"username,omitempty"`
	Channel                   *string `url:"channel,omitempty" json:"channel,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	NotifyOnlyDefaultBranch   *bool   `url:"notify_only_default_branch,omitempty" json:"notify_only_default_branch,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	AlertChannel              *string `url:"alert_channel,omitempty" json:"alert_channel,omitempty"`
	AlertEvents               *bool   `url:"alert_events,omitempty" json:"alert_events,omitempty"`
	ConfidentialIssueChannel  *string `url:"confidential_issue_channel,omitempty" json:"confidential_issue_channel,omitempty"`
	ConfidentialIssuesEvents  *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	ConfidentialNoteChannel   *string `url:"confidential_note_channel,omitempty" json:"confidential_note_channel,omitempty"`
	ConfidentialNoteEvents    *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	DeploymentChannel         *string `url:"deployment_channel,omitempty" json:"deployment_channel,omitempty"`
	DeploymentEvents          *bool   `url:"deployment_events,omitempty" json:"deployment_events,omitempty"`
	IssueChannel              *string `url:"issue_channel,omitempty" json:"issue_channel,omitempty"`
	IssuesEvents              *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	MergeRequestChannel       *string `url:"merge_request_channel,omitempty" json:"merge_request_channel,omitempty"`
	MergeRequestsEvents       *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	NoteChannel               *string `url:"note_channel,omitempty" json:"note_channel,omitempty"`
	NoteEvents                *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	PipelineChannel           *string `url:"pipeline_channel,omitempty" json:"pipeline_channel,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	PushChannel               *string `url:"push_channel,omitempty" json:"push_channel,omitempty"`
	PushEvents                *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	TagPushChannel            *string `url:"tag_push_channel,omitempty" json:"tag_push_channel,omitempty"`
	TagPushEvents             *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	WikiPageChannel           *string `url:"wiki_page_channel,omitempty" json:"wiki_page_channel,omitempty"`
	WikiPageEvents            *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
}

// SetSlackService sets Slack service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#edit-slack-service
func (s *ServicesService) SetSlackService(pid interface{}, opt *SetSlackServiceOptions, options ...RequestOptionFunc) (*SlackService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteSlackService deletes Slack service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-slack-service
func (s *ServicesService) DeleteSlackService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SlackSlashCommandsService represents Slack slash commands settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#slack-slash-commands
type SlackSlashCommandsService struct {
	Service
	Properties *SlackSlashCommandsProperties `json:"properties"`
}

// SlackSlashCommandsProperties represents Slack slash commands specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#slack-slash-commands
type SlackSlashCommandsProperties struct {
	Token string `json:"token"`
}

// GetSlackSlashCommandsService gets Slack slash commands service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-slack-slash-command-integration-settings
func (s *ServicesService) GetSlackSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*SlackSlashCommandsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackSlashCommandsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetSlackSlashCommandsServiceOptions represents the available SetSlackSlashCommandsService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-slack-slash-command-service
type SetSlackSlashCommandsServiceOptions struct {
	Token *string `url:"token,omitempty" json:"token,omitempty"`
}

// SetSlackSlashCommandsService sets Slack slash commands service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/13.12/ee/api/integrations.html#createedit-slack-slash-command-service
func (s *ServicesService) SetSlackSlashCommandsService(pid interface{}, opt *SetSlackSlashCommandsServiceOptions, options ...RequestOptionFunc) (*SlackSlashCommandsService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(SlackSlashCommandsService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteSlackSlashCommandsService deletes Slack slash commands service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/13.12/ee/api/integrations.html#delete-slack-slash-command-service
func (s *ServicesService) DeleteSlackSlashCommandsService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/slack-slash-commands", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// TelegramService represents Telegram service settings.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#telegram
type TelegramService struct {
	Service
	Properties *TelegramServiceProperties `json:"properties"`
}

// TelegramServiceProperties represents Telegram specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-telegram
type TelegramServiceProperties struct {
	Room                      string `json:"room"`
	NotifyOnlyBrokenPipelines bool   `json:"notify_only_broken_pipelines"`
	BranchesToBeNotified      string `json:"branches_to_be_notified"`
}

// GetTelegramService gets MicrosoftTeams service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-telegram-settings
func (s *ServicesService) GetTelegramService(pid interface{}, options ...RequestOptionFunc) (*TelegramService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/telegram", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(TelegramService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetTelegramServiceOptions represents the available SetTelegramService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-telegram
type SetTelegramServiceOptions struct {
	Token                     *string `url:"token,omitempty" json:"token,omitempty"`
	Room                      *string `url:"room,omitempty" json:"room,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	PushEvents                *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
	IssuesEvents              *bool   `url:"issues_events,omitempty" json:"issues_events,omitempty"`
	ConfidentialIssuesEvents  *bool   `url:"confidential_issues_events,omitempty" json:"confidential_issues_events,omitempty"`
	MergeRequestsEvents       *bool   `url:"merge_requests_events,omitempty" json:"merge_requests_events,omitempty"`
	TagPushEvents             *bool   `url:"tag_push_events,omitempty" json:"tag_push_events,omitempty"`
	NoteEvents                *bool   `url:"note_events,omitempty" json:"note_events,omitempty"`
	ConfidentialNoteEvents    *bool   `url:"confidential_note_events,omitempty" json:"confidential_note_events,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
	WikiPageEvents            *bool   `url:"wiki_page_events,omitempty" json:"wiki_page_events,omitempty"`
}

// SetTelegramService sets Telegram service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#set-up-telegram
func (s *ServicesService) SetTelegramService(pid interface{}, opt *SetTelegramServiceOptions, options ...RequestOptionFunc) (*TelegramService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/telegram", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(TelegramService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteTelegramService deletes Telegram service for project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#disable-telegram
func (s *ServicesService) DeleteTelegramService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/telegram", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// YouTrackService represents YouTrack service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#youtrack
type YouTrackService struct {
	Service
	Properties *YouTrackServiceProperties `json:"properties"`
}

// YouTrackServiceProperties represents YouTrack specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#youtrack
type YouTrackServiceProperties struct {
	IssuesURL   string `json:"issues_url"`
	ProjectURL  string `json:"project_url"`
	Description string `json:"description"`
	PushEvents  bool   `json:"push_events"`
}

// GetYouTrackService gets YouTrack service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#get-youtrack-service-settings
func (s *ServicesService) GetYouTrackService(pid interface{}, options ...RequestOptionFunc) (*YouTrackService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/youtrack", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(YouTrackService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, nil
}

// SetYouTrackServiceOptions represents the available SetYouTrackService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-youtrack-service
type SetYouTrackServiceOptions struct {
	IssuesURL   *string `url:"issues_url,omitempty" json:"issues_url,omitempty"`
	ProjectURL  *string `url:"project_url,omitempty" json:"project_url,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	PushEvents  *bool   `url:"push_events,omitempty" json:"push_events,omitempty"`
}

// SetYouTrackService sets YouTrack service for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#createedit-youtrack-service
func (s *ServicesService) SetYouTrackService(pid interface{}, opt *SetYouTrackServiceOptions, options ...RequestOptionFunc) (*YouTrackService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/youtrack", PathEscape(project))

	svc := new(YouTrackService)
	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, nil, err
	}

	return svc, resp, nil
}

// DeleteYouTrackService deletes YouTrack service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/integrations.html#delete-youtrack-service
func (s *ServicesService) DeleteYouTrackService(pid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/youtrack", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
