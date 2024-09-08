package yaygo

import (
	"encoding/json"
	"net/http"
)

type NotificationAPI struct {
	s *Session
}

func newNotificationAPI(s *Session) *NotificationAPI {
	return &NotificationAPI{
		s: s,
	}
}




type GetUserActivitiesParams struct {
	Important     bool `json:"important"`
	FromTimestamp int  `json:"from_timestamp,omitempty"`
	Number        int  `json:"number"`
}

func (h *HiddenAPI) GetUserActivities(params *GetUserActivitiesParams) (st *ActivitiesResponse, err error) {
	resp, err := h.s.request(http.MethodGet, EndpointChatRoomsV1+"get_user_activities", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetUserMergedActivitiesParams struct {
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (h *HiddenAPI) GetUserMergedActivities(params *GetUserMergedActivitiesParams) (st *ActivitiesResponse, err error) {
	resp, err := h.s.request(http.MethodGet, EndpointChatRoomsV2+"get_user_merged_activities", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReceivedNotificationParams struct {
	PID      string `json:"pid"`
	Type     string `json:"type"`
	OpenedAt int    `json:"opened_at"` // long型はint64にマップ
}

func (h *HiddenAPI) ReceivedNotification(params *ReceivedNotificationParams) (st *Response, err error) {
	resp, err := h.s.request(http.MethodPost, EndpointChatRoomsV1+"received_notification", nil, params, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
