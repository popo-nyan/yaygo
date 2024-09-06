package yaygo

type Response struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Response
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	BanUntil  int    `json:"ban_until"`
	RetryIn   int    `json:"retry_in"`
}

type LoginUserResponse struct {
	Response
	AccessToken  string  `json:"access_token"`
	ExpiresIn    int     `json:"expires_in"`
	IsNew        bool    `json:"is_new"`
	RefreshToken string  `json:"refresh_token"`
	SNSInfo      SNSInfo `json:"sns_info"`
	UserID       int     `json:"user_id"`
	Username     string  `json:"username"`
}

type UserTimestampResponse struct {
	Response
	Country   string `json:"country"`
	IPAddress string `json:"ip_address"`
	Time      int    `json:"time"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UnreadStatusResponse struct {
	Response
	IsUnread bool `json:"is_unread"`
}

type CreateChatRoomResponse struct {
	Response
	RoomId int `json:"room_id"`
}

type FollowUsersResponse struct {
	Response
	LastFollowId int  `json:"last_follow_id"`
	Users []User      `json:"users"`
}

type GifsDataResponse struct {
	Response

}

type ChatRoomsResponse struct {
	Response
	PinnedChatRooms []ChatRoom `json:"pinned_chat_rooms"`
	ChatRooms       []ChatRoom `json:"chat_rooms"`
	NextPageValue   int        `json:"next_page_value"`
}

type MessagesResponse struct {
	Response
	Messages []Message `json:"messages"`
}

type AdditionalSettingsResponse struct {
	Response
	Settings Settings `json:"settings"`
}

type ChatRoomResponse struct {
	Response
	ChatRoom ChatRoom `json:"chat_room"`
}

type StickerPacksResponse struct {
	Response
	StickerPacks []StickerPack `json:"sticker_packs"`
}

type TotalChatRequestResponse struct {
	Response
	Total int `json:"total"`
}

type NotificationSettingResponse struct {
	Response
	Setting UserSetting `json:"setting"`
}
