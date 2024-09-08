package yaygo

import (
	"encoding/json"
	"net/http"
)

type CallAPI struct {
	s *Session
}

func newCallAPI(s *Session) *CallAPI {
	return &CallAPI{
		s: s,
	}
}

type BumpCallParams struct {
	CallId           int `json:"call_id,omitempty"`
	ParticipantLimit int `json:"participant_limit,omitempty"`
}

func (c *CallAPI) BumpCall(params *BumpCallParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"bump_call", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GenerateActionSignatureParams struct {
	ConferenceID   int    `json:"conference_id,omitempty"`
	TargetUserUUID string `json:"target_user_uuid,omitempty"`
	Action         string `json:"action,omitempty"`
}

func (c *CallAPI) GenerateActionSignature(params *GenerateActionSignatureParams) (st *CallActionSignatureResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"generate_action_signature", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetActiveCallParams struct {
	UserID int `json:"user_id,omitempty"`
}

func (c *CallAPI) GetActiveCall(params *GetActiveCallParams) (st *PostResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_active_call", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetBgmsParams struct{}

func (c *CallAPI) GetBgms(params *GetBgmsParams) (st *BgmsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_bgms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetCallParams struct {
	CallID int `json:"call_id,omitempty"`
}

func (c *CallAPI) GetCall(params *GetCallParams) (st *ConferenceCallResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2+"get_call", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetCallInvitableUsersParams struct {
	CallID        int    `json:"call_id,omitempty"`
	FromTimestamp int    `json:"from_timestamp,omitempty"`
	UserNickname  string `json:"user_nickname,omitempty"`
}

func (c *CallAPI) GetCallInvitableUsers(params *GetCallInvitableUsersParams) (st *UsersByTimestampResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_call_invitable_users", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetCallStatusParams struct {
	OpponentID int `json:"opponent_id,omitempty"`
}

func (c *CallAPI) GetCallStatus(params *GetCallStatusParams) (st *CallStatusResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_call_status", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetGamesParams struct {
	Number int   `json:"number,omitempty"`
	FromID int   `json:"from_id,omitempty"`
	IDs    []int `json:"ids,omitempty"`
}

func (c *CallAPI) GetGames(params *GetGamesParams) (st *GamesResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_games", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetGenresParams struct {
	Number int `json:"number,omitempty"`
	From   int `json:"from,omitempty"`
}

func (c *CallAPI) GetGenres(params *GetGenresParams) (st *GenresResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_genres", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetGroupCallsParams struct {
	Number          int    `json:"number,omitempty"`
	GroupCategoryID int    `json:"group_category_id,omitempty"`
	FromTimestamp   int    `json:"from_timestamp,omitempty"`
	Scope           string `json:"scope,omitempty"`
}

func (c *CallAPI) GetGroupCalls(params *GetGroupCallsParams) (st *PostsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1+"get_group_calls", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetRtmTokenParams struct {
	CallID int `json:"call_id,omitempty"`
}

func (c *CallAPI) GetRtmToken(params *GetRtmTokenParams) (st *RtmTokenResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2+"get_rtm_token", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type InviteToCallBulkParams struct {
	CallID  int `json:"call_id,omitempty"`
	GroupID int `json:"group_id,omitempty"`
}

func (c *CallAPI) InviteToCallBulk(params *InviteToCallBulkParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"invite_to_call_bulk", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type InviteUsersToCallParams struct {
	CallID  int   `json:"call_id,omitempty"`
	UserIDs []int `json:"user_ids,omitempty"`
}

func (c *CallAPI) InviteUsersToCall(params *InviteUsersToCallParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"invite_users_to_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type InviteUsersToChatCallParams struct {
	ChatRoomID int    `json:"chat_room_id,omitempty"`
	RoomID     int    `json:"room_id,omitempty"`
	RoomURL    string `json:"room_url,omitempty"`
}

func (c *CallAPI) InviteUsersToChatCall(params *InviteUsersToChatCallParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2+"invite_users_to_chat_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type KickFromCallParams struct {
	CallID int    `json:"call_id,omitempty"`
	UUID   string `json:"uuid,omitempty"`
	Ban    bool   `json:"ban,omitempty"`
}

func (c *CallAPI) KickFromCall(params *KickFromCallParams) (st *CallActionSignatureResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3+"kick_from_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type NotifyAnonymousUserLeaveAgoraChannelParams struct {
	ConferenceID int    `json:"conference_id,omitempty"`
	AgoraUID     string `json:"agora_uid,omitempty"`
}

func (c *CallAPI) NotifyAnonymousUserLeaveAgoraChannel(params *NotifyAnonymousUserLeaveAgoraChannelParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"notify_anonymous_user_leave_agora_channel", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type NotifyUserLeaveAgoraChannelParams struct {
	ConferenceID int `json:"conference_id,omitempty"`
	UserID       int `json:"user_id,omitempty"`
}

func (c *CallAPI) NotifyUserLeaveAgoraChannel(params *NotifyUserLeaveAgoraChannelParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"notify_user_leave_agora_channel", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type SendCallScreenshotParams struct {
	ScreenshotFilename string `json:"screenshot_filename,omitempty"`
	ConferenceID       int    `json:"conference_id,omitempty"`
}

func (c *CallAPI) SendCallScreenshot(params *SendCallScreenshotParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPut, EndpointChatRoomsV1+"send_call_screenshot", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type SetCallParams struct {
	CallID     int    `json:"call_id,omitempty"`
	JoinableBy string `json:"joinable_by,omitempty"`
	GameTitle  string `json:"game_title,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
}

func (c *CallAPI) SetCall(params *SetCallParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPut, EndpointChatRoomsV1+"set_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type SetUserRoleParams struct {
	CallID int    `json:"call_id,omitempty"`
	UserID int    `json:"user_id,omitempty"`
	Role   string `json:"role,omitempty"`
}

func (c *CallAPI) SetUserRole(params *SetUserRoleParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPut, EndpointChatRoomsV1+"set_user_role", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type StartCallParams struct {
	ConferenceID int    `json:"conference_id,omitempty"`
	CallSID      string `json:"call_sid,omitempty"`
}

func (c *CallAPI) StartCall(params *StartCallParams) (st *ConferenceCallResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2+"start_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type StopCallParams struct {
	ConferenceID     int    `json:"conference_id,omitempty"`
	CallSID          string `json:"call_sid,omitempty"`
	Duration         int    `json:"duration,omitempty"`
	TotalUsersInCall int    `json:"total_users_in_call,omitempty"`
}

func (c *CallAPI) StopCall(params *StopCallParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1+"stop_call", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type ValidateActionSignatureParams struct {
	CallID       int    `json:"call_id,omitempty"`
	SenderUUID   string `json:"sender_uuid,omitempty"`
	ReceiverUUID string `json:"receiver_uuid,omitempty"`
	Action       string `json:"action,omitempty"`
	Timestamp    int    `json:"timestamp,omitempty"`
	Signature    string `json:"signature,omitempty"`
}

func (c *CallAPI) ValidateActionSignature(params *ValidateActionSignatureParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2+"validate_action_signature", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
