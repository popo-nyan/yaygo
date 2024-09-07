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
	PostID   int64 `json:"post_id,omitempty"`
	ThreadID int64 `json:"thread_id,omitempty"`
}

func (t *ThreadAPI) AddPostToThread(params *AddPostToThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPut, EndpointChatRoomsV1 + "add_post_to_thread", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ConvertPostToThreadParams struct {
	PostID     int64   `json:"post_id,omitempty"`
	ThreadTitle *string `json:"thread_title,omitempty"`
	ThreadIcon  *string `json:"thread_icon_filename,omitempty"`
}

func (t *ThreadAPI) ConvertPostToThread(params *ConvertPostToThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPost, EndpointChatRoomsV1 + "convert_post_to_thread", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateThreadParams struct {
	Request *CreateGroupThreadRequest `json:"request,omitempty"`
}

func (t *ThreadAPI) CreateThread(params *CreateThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodPost, EndpointChatRoomsV1 + "create_thread", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupThreadListParams struct {
	GroupID   int64   `json:"group_id,omitempty"`
	From      *string `json:"from,omitempty"`
	JoinStatus *string `json:"join_status,omitempty"`
}

func (t *ThreadAPI) GetGroupThreadList(params *GetGroupThreadListParams) (st *GroupThreadListResponse, err error) {
	resp, err := t.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_group_thread_list", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetJoinedStatusesParams struct {
	IDs []int64 `json:"ids,omitempty"`
}

func (t *ThreadAPI) GetJoinedStatuses(params *GetJoinedStatusesParams) (st map[string]string, err error) {
	resp, err := t.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_joined_statuses", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetThreadParams struct {
	ThreadID int64 `json:"id,omitempty"`
}

func (t *ThreadAPI) GetThread(params *GetThreadParams) (st *ThreadInfo, err error) {
	resp, err := t.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_thread", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetThreadPostsParams struct {
	ThreadID int64   `json:"id,omitempty"`
	PostType *string `json:"post_type,omitempty"`
	Number   *int    `json:"number,omitempty"`
	From     *int64  `json:"from,omitempty"`
}

func (t *ThreadAPI) GetThreadPosts(params *GetThreadPostsParams) (st *PostsResponse, err error) {
	resp, err := t.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_thread_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type JoinThreadParams struct {
	ThreadID int64 `json:"thread_id,omitempty"`
	UserID   int64 `json:"id,omitempty"`
}

func (t *ThreadAPI) JoinThread(params *JoinThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodPost, EndpointChatRoomsV1 + "join_thread", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type LeaveThreadParams struct {
	ThreadID int64 `json:"thread_id,omitempty"`
	UserID   int64 `json:"id,omitempty"`
}

func (t *ThreadAPI) LeaveThread(params *LeaveThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodDelete, EndpointChatRoomsV1 + "leave_thread", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveThreadParams struct {
	ThreadID int64 `json:"id,omitempty"`
}

func (t *ThreadAPI) RemoveThread(params *RemoveThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodDelete, EndpointChatRoomsV1 + "remove_thread", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateThreadParams struct {
	ThreadID    int64   `json:"id,omitempty"`
	ThreadTitle string  `json:"thread_title,omitempty"`
	ThreadIcon  *string `json:"thread_icon_filename,omitempty"`
}

func (t *ThreadAPI) UpdateThread(params *UpdateThreadParams) (st *Response, err error) {
	resp, err := t.s.request(http.MethodPut, EndpointChatRoomsV1 + "update_thread", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
