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
	ChatRoomIds []int `json:"chat_room_ids[],omitempty"`
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
	FromTime int `json:"from_time,omitempty"`
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
	Name               string `json:"name,omitempty"`
	WithUserIds        []int  `json:"with_user_ids[],omitempty"`
	IconFilename       string `json:"icon_filename,omitempty"`
	BackgroundFilename string `json:"background_filename,omitempty"`
}

func (c *ChatApi) CreateGroup(params *CreateGroupParams) (st * CreateChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "createGroup", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type CreatePrivateParams struct {
	WithUserId int  `json:"with_user_id,omitempty"`
	MatchingId int  `json:"matching_id,omitempty"`
	HimaChat   bool `json:"hima_chat,omitempty"`
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
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV2 + "deleteBackground", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type EditParams struct {
	Id                 int    `json:"id,omitempty"`
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
	FromFollowId  int                `json:"from_follow_id,omitempty"`
	FromTimestamp int                `json:"from_timestamp,omitempty"`
	OrderBy       string             `json:"order_by,omitempty"`
}

func (c *ChatApi) GetChatableUsers(params *GetChatableUsersParams) (st * FollowUsersResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "getChatableUsers", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetGifsDataParams struct {}

func (c *ChatApi) GetGifsData(params *GetGifsDataParams) (st * GifsDataResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "getGifsData", params, nil, false)
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
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "getHiddenChatRooms", params, nil, false)
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
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "getHiddenChatRooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetMessagesParams struct {
	Id            int `json:"id,omitempty"`
	Number        int `json:"number,omitempty"`
	FromMessageId int `json:"from_message_id,omitempty"`
	ToMessageId   int `json:"to_message_id,omitempty"`
}

func (c *ChatApi) GetMessages(params *GetMessagesParams) (st * MessagesResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "getMessages", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetNotificationSettingsParams struct {
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) GetNotificationSettings(params *GetNotificationSettingsParams) (st * AdditionalSettingsResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "getNotificationSettings", params, nil, false)
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
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "getRequestRooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetRoomParams struct {
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) GetRoom(params *GetRoomParams) (st * ChatRoomResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "getRoom", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type GetStickerPacksParams struct {}

func (c *ChatApi) GetStickerPacks(params *GetStickerPacksParams) (st * StickerPacksResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "getStickerPacks", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetTotalRequestsParams struct {}

func (c *ChatApi) GetTotalRequests(params *GetTotalRequestsParams) (st * TotalChatRequestResponse, err error) {
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV1 + "getTotalRequests", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type HideChatParams struct {
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) HideChat(params *HideChatParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "hideChat", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type InviteParams struct {
	Id          int   `json:"id,omitempty"`
	WithUserIds []int `json:"with_user_ids[],omitempty"`
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
	Id          int   `json:"id,omitempty"`
	WithUserIds []int `json:"with_user_ids[],omitempty"`
}

func (c *ChatApi) KickUsers(params *KickUsersParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "kickUsers", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type PinParams struct {
	Id int `json:"id,omitempty"`
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
	Id      int                   `json:"id,omitempty"`
	Request ReadAttachmentRequest `json:"request,omitempty"`
}

func (c *ChatApi) ReadAttachment(params *ReadAttachmentParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "readAttachment", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReadMessageParams struct {
	Id        int `json:"id,omitempty"`
	MessageId int `json:"message_id,omitempty"`
}

func (c *ChatApi) ReadMessage(params *ReadMessageParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "readMessage", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type ReadVideoMessageParams struct {
	Id          int `json:"id,omitempty"`
	VideoMsgIds []int `json:"video_msg_ids[],omitempty"`
	MessageId   int `json:"message_id,omitempty"`
}

func (c *ChatApi) ReadVideoMessage(params *ReadVideoMessageParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "readVideoMessage", nil, params, false)
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
	resp, err := c.s.request(http.MethodGet, EndpointChatRoomsV2 + "refreshRooms", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type RemoveParams struct {
	ChatRoomIds []int `json:"chat_room_ids[],omitempty"`
	Id          int `json:"id,omitempty"`
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
	ChatRoomId          []int   `json:"chat_room_id,omitempty"`
	CategoryId          int     `json:"category_id,omitempty"`
	Reason              string  `json:"reason,omitempty"`
	OpponentId          int     `json:"opponent_id,omitempty"`
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
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) SendMediaScreenshotNotification(params *SendMediaScreenshotNotificationParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV1 + "sendMediaScreenshotNotification", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type SendMessageParams struct {
	Id                 int    `json:"id,omitempty"`
	MessageType        string `json:"message_type,omitempty"`
	CallType           string `json:"call_type,omitempty"`
	Text               string `json:"text,omitempty"`
	FontSize           int    `json:"font_size,omitempty"`
	GifImageId         int    `json:"gif_image_id,omitempty"`
	AttachmentFileName string `json:"attachment_file_name,omitempty"`
	StickerPackId      int    `json:"sticker_pack_id,omitempty"`
	StickerId          int    `json:"sticker_id,omitempty"`
	VideoFileName      string `json:"video_file_name,omitempty"`
	ParentId           int    `json:"parent_id,omitempty"`
}

func (c *ChatApi) SendMessage(params *SendMessageParams) (st * MessageResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV3 + "sendMessage", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type SetNotificationSettingsParams struct {
	Id               int `json:"id,omitempty"`
	NotificationChat int `json:"notification_chat,omitempty"`
}

func (c *ChatApi) SetNotificationSettings(params *SetNotificationSettingsParams) (st * NotificationSettingResponse, err error) {
	resp, err := c.s.request(http.MethodPost, EndpointChatRoomsV2 + "setNotificationSettings", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type UnHideChatParams struct {
	ChatRoomIds int `json:"chat_room_ids,omitempty"`
}

func (c *ChatApi) UnHideChat(params *UnHideChatParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unHideChat", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}


type UnpinParams struct {
	Id int `json:"id,omitempty"`
}

func (c *ChatApi) Unpin(params UnpinParams) (st * Response, err error) {
	resp, err := c.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unpin", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
