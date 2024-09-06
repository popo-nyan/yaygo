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
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "accept_chat_request", nil, params, false)
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
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "unread_status", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}





