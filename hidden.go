package yaygo

import (
	"encoding/json"
	"net/http"
)

type HiddenApi struct {
	s *Session
}

func newHiddenApi(s *Session) *HiddenApi {
	return &HiddenApi{
		s: s,
	}
}



type GetListParams struct {
	From   string `json:"from"`
	Number int    `json:"number"`
}

func (c *HiddenApi) GetList(params *CreatePrivateParams) (st * HiddenResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "GetList", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type HideUserParams struct {
	UserID int `json:"user_id"`
}

func (c *HiddenApi) HideUser(params *HideUserParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "HideUser", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type UnHideUsersParams struct {
	UserIDs []int `json:"user_ids[]"`
}

func (c *HiddenApi) UnHideUsers(params *CreatePrivateParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unHideUsers", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
