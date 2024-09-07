package yaygo

import (
	"encoding/json"
	"net/http"
)

type ThreadApi struct {
	s *Session
}

func newThreadApi(s *Session) *ThreadApi {
	return &ThreadApi{
		s: s,
	}
}



type AddPostToThreadParams struct {
	PostID   int64 `json:"post_id"`
	ThreadID int64 `json:"thread_id"`
}

func (t *ThreadAPI) AddPostToThread(params *AddPostToThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPut, fmt.Sprintf("v3/posts/%d/move_to_thread/%d", params.PostID, params.ThreadID), nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ConvertPostToThreadParams struct {
	PostID     int64   `json:"post_id"`
	ThreadTitle *string `json:"thread_title,omitempty"`
	ThreadIcon  *string `json:"thread_icon_filename,omitempty"`
}

func (t *ThreadAPI) ConvertPostToThread(params *ConvertPostToThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPost, fmt.Sprintf("v3/posts/%d/move_to_thread", params.PostID), nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateThreadParams struct {
	Request *CreateGroupThreadRequest `json:"request"`
}

func (t *ThreadAPI) CreateThread(params *CreateThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPost, "v1/threads", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupThreadListParams struct {
	GroupID   int64   `json:"group_id"`
	From      *string `json:"from,omitempty"`
	JoinStatus *string `json:"join_status,omitempty"`
}

func (t *ThreadAPI) GetGroupThreadList(params *GetGroupThreadListParams) (st *GroupThreadListResponse, err error) {
	resp, err := t.s.request(http.MethodGet, "v1/threads", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetJoinedStatusesParams struct {
	IDs []int64 `json:"ids"`
}

func (t *ThreadAPI) GetJoinedStatuses(params *GetJoinedStatusesParams) (st map[string]string, err error) {
	resp, err := t.s.request(http.MethodGet, "v1/threads/joined_statuses", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetThreadParams struct {
	ThreadID int64 `json:"id"`
}

func (t *ThreadAPI) GetThread(params *GetThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodGet, fmt.Sprintf("v1/threads/%d", params.ThreadID), nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetThreadPostsParams struct {
	ThreadID int64   `json:"id"`
	PostType *string `json:"post_type,omitempty"`
	Number   *int    `json:"number,omitempty"`
	From     *int64  `json:"from,omitempty"`
}

func (t *ThreadAPI) GetThreadPosts(params *GetThreadPostsParams) (st *PostsResponse, err error) {
	resp, err := t.s.request(http.MethodGet, fmt.Sprintf("v1/threads/%d/posts", params.ThreadID), params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type JoinThreadParams struct {
	ThreadID int64 `json:"thread_id"`
	UserID   int64 `json:"id"`
}

func (t *ThreadAPI) JoinThread(params *JoinThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodPost, fmt.Sprintf("v1/threads/%d/members/%d", params.ThreadID, params.UserID), nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type LeaveThreadParams struct {
	ThreadID int64 `json:"thread_id"`
	UserID   int64 `json:"id"`
}

func (t *ThreadAPI) LeaveThread(params *LeaveThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodDelete, fmt.Sprintf("v1/threads/%d/members/%d", params.ThreadID, params.UserID), nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveThreadParams struct {
	ThreadID int64 `json:"id"`
}

func (t *ThreadAPI) RemoveThread(params *RemoveThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodDelete, fmt.Sprintf("v1/threads/%d", params.ThreadID), nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateThreadParams struct {
	ThreadID    int64   `json:"id"`
	ThreadTitle string  `json:"thread_title"`
	ThreadIcon  *string `json:"thread_icon_filename,omitempty"`
}

func (t *ThreadAPI) UpdateThread(params *UpdateThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodPut, fmt.Sprintf("v1/threads/%d", params.ThreadID), nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}