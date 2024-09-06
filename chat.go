package yaygo

import (
	"encoding/json"
	"net/http"
)

type ChatApi struct {
	s *Session
}

func newChatApi(s *Session) *ChatApi {
	return &ChatApi{
		s: s,
	}
}



type AcceptRequestParams struct {
	Chat_room_ids []int `json:"chat_room_ids[],omitempty"`
}

func (c *ChatApi) AcceptRequest(params *AcceptRequestParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "acceptRequest", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CheckUnreadStatusParams struct {
	From_time int `json:"from_time,omitempty"`
}

func (c *ChatApi) CheckUnreadStatus(params *CheckUnreadStatusParams) (st * UnreadStatusResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "checkUnreadStatus", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CreateGroupParams struct {
	Name                string `json:"name,omitempty"`
	With_user_ids       []int  `json:"with_user_ids[],omitempty"`
	Icon_filename       string `json:"icon_filename,omitempty"`
	Background_filename string `json:"background_filename,omitempty"`
}

func (c *ChatApi) CreateGroup(params *CreateGroupParams) (st * CreateChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "createGroup", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CreatePrivateParams struct {
	With_user_id int  `json:"with_user_id,omitempty"`
	Matching_id  int  `json:"matching_id,omitempty"`
	Hima_chat    bool `json:"hima_chat,omitempty"`
}

func (c *ChatApi) CreatePrivate(params *CreatePrivateParams) (st * CreateChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "createPrivate", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type DeleteBackgroundParams struct {
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) DeleteBackground(params *DeleteBackgroundParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "deleteBackground", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type EditParams struct {
	Id                  int    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Icon_filename       string `json:"icon_filename,omitempty"`
	Background_filename string `json:"background_filename,omitempty"`
}

func (c *ChatApi) Edit(params *EditParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "edit", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

