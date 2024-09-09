package yaygo

import (
	"encoding/json"
	"net/http"
)

type HiddenAPI struct {
	s *Session
}

func newHiddenAPI(s *Session) *HiddenAPI {
	return &HiddenAPI{
		s: s,
	}
}

type GetListParams struct {
	From   string `json:"from,omitempty"`
	Number int    `json:"number,omitempty"`
}

func (h *HiddenAPI) GetList(params *CreatePrivateParams) (st *HiddenResponse, err error) {
	resp, err := h.s.request(http.MethodGet, EndpointChatRoomsV1+"GetList", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type HideUserParams struct {
	UserID int `json:"user_id,omitempty"`
}

func (h *HiddenAPI) HideUser(params *HideUserParams) (st *Response, err error) {
	resp, err := h.s.request(http.MethodPost, EndpointChatRoomsV1+"HideUser", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type UnHideUsersParams struct {
	UserIDs []int `json:"user_ids[],omitempty"`
}

func (h *HiddenAPI) UnHideUsers(params *CreatePrivateParams) (st *Response, err error) {
	resp, err := h.s.request(http.MethodDelete, EndpointChatRoomsV1+"unHideUsers", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
