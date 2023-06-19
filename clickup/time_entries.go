package clickup

import (
	"context"
	"encoding/json"
	"fmt"
)

type TimeEntriesService service

type GetTimeEntriesResponse struct {
	TimeEntries []TimeEntry `json:"data"`
}

type TimeEntry struct {
	ID           string      `json:"id,omitempty"`
	Task         Task        `json:"task,omitempty"`
	WID          string      `json:"wid,omitempty"` // TeamSpace ID
	User         User        `json:"user,omitempty"`
	Billable     bool        `json:"billable,omitempty"`
	Duration     json.Number `json:"duration,omitempty"`
	Start        *Date       `json:"start,omitempty"`
	End          *Date       `json:"end,omitempty"`
	At           *Date       `json:"at,omitempty"`
	Description  string      `json:"description,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	Source       string      `json:"source,omitempty"` // clickup
	TaskLocation struct {
		ListID     string `json:"list_id,omitempty"`
		FolderID   string `json:"folder_id,omitempty"`
		SpaceID    string `json:"space_id,omitempty"`
		ListName   string `json:"list_name,omitempty"`
		FolderName string `json:"folder_name,omitempty"`
		SpaceName  string `json:"space_name,omitempty"`
	} `json:"task_location,omitempty"`
	TaskTags []Tag  `json:"task_tags,omitempty"`
	TaskURL  string `json:"task_url,omitempty"`
}

type GetTimeEntriesOptions struct {
	StartDate            *Date `url:"start_date,omitempty"`
	EndDate              *Date `url:"end_date,omitempty"`
	Assignee             int   `url:"assignee,omitempty"`
	IncludeTaskTags      bool  `url:"include_task_tags,omitempty"`
	IncludelocationNames bool  `url:"includelocation_names,omitempty"`
	SpaceID              int   `url:"space_id,omitempty"`
	FolderID             int   `url:"folder_id,omitempty"`
	ListID               int   `url:"list_id,omitempty"`
	TaskID               int   `url:"task_id,omitempty"`
	CustomTaskIDs        bool  `url:"custom_task_ids,omitempty"`
	TeamID               int   `url:"team_id,omitempty"`
}

func (s *TimeEntriesService) GetTimeEntries(ctx context.Context, teamID string, opts *GetTimeEntriesOptions) ([]TimeEntry, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries", teamID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetTimeEntriesResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.TimeEntries, resp, nil
}

type GetCurrentTimeEntryOptions struct {
	Assignee int `url:"assignee,omitempty"`
}

type GetCurrentTimeEntryResponse struct {
	TimeEntry *TimeEntry `json:"data"`
}

func (s *TimeEntriesService) GetCurrentTimeEntry(ctx context.Context, teamID string, opts *GetTimeEntriesOptions) (*TimeEntry, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries/current", teamID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetCurrentTimeEntryResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.TimeEntry, resp, nil
}

// StopTimer stop a timer that's currently running for the authenticated user.
func (s *TimeEntriesService) StopTimer(ctx context.Context, teamID string) (*TimeEntry, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries/stop", teamID)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetCurrentTimeEntryResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.TimeEntry, resp, nil
}

type StartTimerRequest struct {
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
	TaskID      string `json:"tid,omitempty"`
	Billable    bool
}

// StartTimer Start a timer for the authenticated user.
func (s *TimeEntriesService) StartTimer(ctx context.Context, teamID string, reqUp *StartTimerRequest) (*TimeEntry, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries/start", teamID)

	req, err := s.client.NewRequest("POST", u, reqUp)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetCurrentTimeEntryResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.TimeEntry, resp, nil
}

type TimeEntryUpdateRequest struct {
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
	TagAction   string `json:"tag_action,omitempty"` // Accessible tag actions are ["replace", "add", "remove"]
	Start       *Date  `json:"start,omitempty"`
	End         *Date  `json:"end,omitempty"`
	TaskID      string `json:"tid,omitempty"`
	Billable    bool
	Duration    int
}

// StopTimer stop a timer that's currently running for the authenticated user.
func (s *TimeEntriesService) UpdateTimeEntry(ctx context.Context, teamID, timeEntryID string, upReq *TimeEntryUpdateRequest) (*TimeEntry, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries/%s", teamID, timeEntryID)

	req, err := s.client.NewRequest("PUT", u, upReq)
	if err != nil {
		return nil, nil, err
	}

	// gtr := new(GetCurrentTimeEntryResponse)
	var respObj = struct{}{}
	resp, err := s.client.Do(ctx, req, &respObj)
	if err != nil {
		return nil, resp, err
	}

	return nil, resp, nil
}
