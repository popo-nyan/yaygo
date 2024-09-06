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
	ChatRoomIDs []int `json:"chat_room_ids[],omitempty"`
}

func (c *ChatApi) AcceptRequest(params *AcceptRequestParams) (st *Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "accept_request", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CheckUnreadStatusParams struct {
	FromTime int `json:"from_time,omitempty"`
}

func (c *ChatApi) CheckUnreadStatus(params *CheckUnreadStatusParams) (st * UnreadStatusResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "check_unread_status", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CreateGroupParams struct {
	Name               string `json:"name,omitempty"`
	WithUserIDs        []int  `json:"with_user_ids[],omitempty"`
	IconFilename       string `json:"icon_filename,omitempty"`
	BackgroundFilename string `json:"background_filename,omitempty"`
}

func (c *ChatApi) CreateGroup(params *CreateGroupParams) (st * CreateChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "create_group", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CreatePrivateParams struct {
	WithUserID int  `json:"with_user_id,omitempty"`
	MatchingID int  `json:"matching_id,omitempty"`
	HimaChat   bool `json:"hima_chat,omitempty"`
}

func (c *ChatApi) CreatePrivate(params *CreatePrivateParams) (st * CreateChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "create_private", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type DeleteBackgroundParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) DeleteBackground(params *DeleteBackgroundParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV2 + "delete_background", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type EditParams struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	IconFilename       string `json:"icon_filename,omitempty"`
	BackgroundFilename string `json:"background_filename,omitempty"`
}

func (c *ChatApi) Edit(params *EditParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "edit", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetChatableUsersParams struct {
	Request       SearchUsersRequest `json:"request,omitempty"`
	FromFollowID  int                `json:"from_follow_id,omitempty"`
	FromTimestamp int                `json:"from_timestamp,omitempty"`
	OrderBy       string             `json:"order_by,omitempty"`
}

func (c *ChatApi) GetChatableUsers(params *GetChatableUsersParams) (st * FollowUsersResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "get_chatable_users", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetGifsDataParams struct {}

func (c *ChatApi) GetGifsData(params *GetGifsDataParams) (st * GifsDataResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_gifs_data", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetHiddenChatRoomsParams struct {
	Number        int `json:"number,omitempty"`
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (c *ChatApi) GetHiddenChatRooms(params *GetHiddenChatRoomsParams) (st * ChatRoomsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_hidden_chat_rooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetMainRoomsParams struct {
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (c *ChatApi) GetMainRooms(params *GetMainRoomsParams) (st * ChatRoomsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_hidden_chat_rooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetMessagesParams struct {
	ID            int `json:"id,omitempty"`
	Number        int `json:"number,omitempty"`
	FromMessageID int `json:"from_message_id,omitempty"`
	ToMessageID   int `json:"to_message_id,omitempty"`
}

func (c *ChatApi) GetMessages(params *GetMessagesParams) (st * MessagesResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_messages", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetNotificationSettingsParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) GetNotificationSettings(params *GetNotificationSettingsParams) (st * AdditionalSettingsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_notification_settings", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetRequestRoomsParams struct {
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (c *ChatApi) GetRequestRooms(params *GetRequestRoomsParams) (st * ChatRoomsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_request_rooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetRoomParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) GetRoom(params *GetRoomParams) (st * ChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_room", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetStickerPacksParams struct {}

func (c *ChatApi) GetStickerPacks(params *GetStickerPacksParams) (st * StickerPacksResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_sticker_packs", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetTotalRequestsParams struct {}

func (c *ChatApi) GetTotalRequests(params *GetTotalRequestsParams) (st * TotalChatRequestResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_total_requests", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type HideChatParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) HideChat(params *HideChatParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "hide_chat", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type InviteParams struct {
	ID          int   `json:"id,omitempty"`
	WithUserIDs []int `json:"with_user_ids[],omitempty"`
}

func (c *ChatApi) Invite(params *InviteParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "invite", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type KickUsersParam struct {
	ID          int   `json:"id,omitempty"`
	WithUserIDs []int `json:"with_user_ids[],omitempty"`
}

func (c *ChatApi) KickUsers(params *KickUsersParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "kick_users", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type PinParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) Pin(params *PinParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "pin", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReadAttachmentParams struct {
	ID      int                   `json:"id,omitempty"`
	Request ReadAttachmentRequest `json:"request,omitempty"`
}

func (c *ChatApi) ReadAttachment(params *ReadAttachmentParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "read_attachment", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReadMessageParams struct {
	ID        int `json:"id,omitempty"`
	MessageID int `json:"message_id,omitempty"`
}

func (c *ChatApi) ReadMessage(params *ReadMessageParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "read_message", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReadVideoMessageParams struct {
	ID          int `json:"id,omitempty"`
	VideoMsgIDs []int `json:"video_msg_ids[],omitempty"`
	MessageID   int `json:"message_id,omitempty"`
}

func (c *ChatApi) ReadVideoMessage(params *ReadVideoMessageParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "read_video_message", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type RefreshRoomsParams struct {
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (c *ChatApi) RefreshRooms(params *RefreshRoomsParams) (st * ChatRoomsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "refresh_rooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type RemoveParams struct {
	ChatRoomIDs []int `json:"chat_room_ids[],omitempty"`
	ID          int `json:"id,omitempty"`
}

func (c *ChatApi) Remove(params *RemoveParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "remove", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReportParams struct {
	ChatRoomID          []int   `json:"chat_room_id,omitempty"`
	CategoryID          int     `json:"category_id,omitempty"`
	Reason              string  `json:"reason,omitempty"`
	OpponentID          int     `json:"opponent_id,omitempty"`
	ScreenshotFilename  string  `json:"screenshot_filename,omitempty"`
	Screenshot2Filename string  `json:"screenshot_2_filename,omitempty"`
	Screenshot3Filename string  `json:"screenshot_3_filename,omitempty"`
	Screenshot4Filename string  `json:"screenshot_4_filename,omitempty"`
}

func (c *ChatApi) Report(params *ReportParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "report", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type SendMediaScreenshotNotificationParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) SendMediaScreenshotNotification(params *SendMediaScreenshotNotificationParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "send_media_screenshot_notification", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type SendMessageParams struct {
	ID                 int    `json:"id,omitempty"`
	MessageType        string `json:"message_type,omitempty"`
	CallType           string `json:"call_type,omitempty"`
	Text               string `json:"text,omitempty"`
	FontSize           int    `json:"font_size,omitempty"`
	GifImageID         int    `json:"gif_image_id,omitempty"`
	AttachmentFileName string `json:"attachment_file_name,omitempty"`
	StickerPackID      int    `json:"sticker_pack_id,omitempty"`
	StickerID          int    `json:"sticker_id,omitempty"`
	VideoFileName      string `json:"video_file_name,omitempty"`
	ParentID           int    `json:"parent_id,omitempty"`
}

func (c *ChatApi) SendMessage(params *SendMessageParams) (st * MessageResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "send_message", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type SetNotificationSettingsParams struct {
	ID               int `json:"id,omitempty"`
	NotificationChat int `json:"notification_chat,omitempty"`
}

func (c *ChatApi) SetNotificationSettings(params *SetNotificationSettingsParams) (st * NotificationSettingResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "set_notification_settings", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type UnHideChatParams struct {
	ChatRoomIDs int `json:"chat_room_ids,omitempty"`
}

func (c *ChatApi) UnHideChat(params *UnHideChatParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "un_hide_chat", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type UnpinParams struct {
	ID int `json:"id,omitempty"`
}

func (c *ChatApi) Unpin(params UnpinParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unpin", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
