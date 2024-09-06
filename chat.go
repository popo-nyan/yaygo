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


type GetChatableUsersParams struct {
	Request        SearchUsersRequest `json:"request,omitempty"`
	From_follow_id int                `json:"from_follow_id,omitempty"`
	From_timestamp int                `json:"from_timestamp,omitempty"`
	Order_by       string             `json:"order_by,omitempty"`
}




type GetGifsDataParams struct {}







type GetHiddenChatRoomsParams struct {
	Number int         `json:"number,omitempty"`
	From_timestamp int `json:"from_timestamp,omitempty"`
}






type GetMainRoomsParams struct {
	From_timestamp int `json:"from_timestamp,omitempty"`
}




type GetMessagesParams struct {
	Id              int `json:"id,omitempty"`
	Number          int `json:"number,omitempty"`
	From_message_id int `json:"from_message_id,omitempty"`
	To_message_id   int `json:"to_message_id,omitempty"`
}




type GetNotificationSettingsParams struct {
	Id int `json:"id,omitempty"`
}




type GetRequestRoomsParams struct {
	From_timestamp int `json:"from_timestamp,omitempty"`
}




type GetRoomParams struct {
	Id int `json:"id,omitempty"`
}




type GetStickerPacksParams struct {}




type GetTotalRequestsParams struct {}




type HideChatParams struct {
	Id int `json:"id,omitempty"`
}




type InviteParams struct {
	Id            int   `json:"id,omitempty"`
	With_user_ids []int `json:"with_user_ids[],omitempty"`
}



type KickUsersParam struct {
	Id            int   `json:"id,omitempty"`
	With_user_ids []int `json:"with_user_ids[],omitempty"`
}


type PinParams struct {
	Id int `json:"id,omitempty"`
}




type ReadAttachmentParams struct {
	Id      int                   `json:"id,omitempty"`
	Request ReadAttachmentRequest `json:"request,omitempty"`
}




type ReadMessageParams struct {
	Id         int `json:"id,omitempty"`
	Message_id int `json:"message_id,omitempty"`
}




type ReadVideoMessageParams struct {
	Id            int `json:"id,omitempty"`
	Video_msg_ids []int `json:"video_msg_ids[],omitempty"`
	Message_id    int `json:"message_id,omitempty"`
}




type RefreshRoomsParams struct {
	From_timestamp int `json:"from_timestamp,omitempty"`
}




type RemoveParams struct {
	Chat_room_ids []int `json:"chat_room_ids[],omitempty"`
	Id            int `json:"id,omitempty"`
}




type ReportParams struct {
	Chat_room_id          []int   `json:"chat_room_id,omitempty"`
	Category_id           int     `json:"category_id,omitempty"`
	Reason                string  `json:"reason,omitempty"`
	Opponent_id           int     `json:"opponent_id,omitempty"`
	Screenshot_filename   string  `json:"screenshot_filename,omitempty"`
	Screenshot_2_filename string  `json:"screenshot_2_filename,omitempty"`
	Screenshot_3_filename string  `json:"screenshot_3_filename,omitempty"`
	Screenshot_4_filename string  `json:"screenshot_4_filename,omitempty"`
}






type SendMediaScreenshotNotificationParams struct {
	Id int `json:"id,omitempty"`
}







type SendMessageParams struct {
	Id                   int    `json:"id,omitempty"`
	Message_type         string `json:"message_type,omitempty"`
	Call_type            string `json:"call_type,omitempty"`
	Text                 string `json:"text,omitempty"`
	Font_size            int    `json:"font_size,omitempty"`
	Gif_image_id         int    `json:"gif_image_id,omitempty"`
	Attachment_file_name string `json:"attachment_file_name,omitempty"`
	Sticker_pack_id      int    `json:"sticker_pack_id,omitempty"`
	Sticker_id           int    `json:"sticker_id,omitempty"`
	Video_file_name      string `json:"video_file_name,omitempty"`
	Parent_id            int    `json:"parent_id,omitempty"`
}







type SetNotificationSettingsParams struct {
	Id                int `json:"id,omitempty"`
	Notification_chat int `json:"notification_chat,omitempty"`
}







type UnHideChatParams struct {
	Chat_room_ids int `json:"chat_room_ids,omitempty"`
}






type UnpinParams struct {
	Id int `json:"id,omitempty"`








