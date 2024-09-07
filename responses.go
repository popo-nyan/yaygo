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



type CallActionSignatureResponse struct {
	Response
	SignaturePayload SignaturePayload `json:"signature_payload"`
}

type PostResponse struct {
	Response
	Post Post `json:"post"`
}

type BgmsResponse struct {
	Response
	Bgm []Bgm `json:"bgm"`
}

type ConferenceCallResponse struct {
	Response
	ConferenceCall         ConferenceCall `json:"conference_call"`
	ConferenceCallUserUUID string         `json:"conference_call_user_uuid"`
}

type UsersByTimestampResponse struct {
	Response
	LastTimestamp int    `json:"last_timestamp"`
	Users         []User `json:"users"`
}

type CallStatusResponse struct {
	Response
	PhoneStatus bool   `json:"phone_status"`
	VideoStatus bool   `json:"video_status"`
	RoomUrl     string `json:"room_url"`
}

type GamesResponse struct {
	Response
	Games  []Game `json:"games"`
	FromID int    `json:"from_id"`
}

type GenresResponse struct {
	Response
	Genres        []Genre `json:"genres"`
	NextPageValue int     `json:"next_page_value"`
}

type PostsResponse struct {
	Response
	NextPageValue   string `json:"mext_page_value"`
	Posts           []Post `json:"posts"`
	PinnedPosts     []Post `json:"pinned_posts"`
	HasMoreHotPosts bool   `json:"has_more_hot_posts"`
}

type RtmTokenResponse struct {
	Response
	Token string `json:"token"`
}




type HiddenResponse struct {
	Response
	HiddenUsers   []User `json:"hidden_users"`
	NextPageValue string `json:"next_page_value"`
	TotalCount    int    `json:"total_count"`
	Limit         int    `json:"limit"`
}




type ReviewsResponse struct {
	Response
	Reviews       []Review `json:"review"`
	PinnedReviews []Review `json:"pinned_reviews"`
}



type CreateGroupResponse struct {
	Response

}



type UsersResponse struct {
	Response
	
}


type CreateQuotaResponse struct {
	Response
	
}

type GroupCategoriesResponse struct {
	Response
	
}

type GroupResponse struct {
	Response
	
}

type GroupNotificationSettingsResponse struct {
	Response
	
}

type GroupsResponse struct {
	Response
	
}

type GroupUsersResponse struct {
	Response
	
}

type GroupUserResponse struct {
	Response
	
}

type GroupsRelatedResponse struct {
	Response
	
}

type GroupInCircleUserLeaderboardResponse struct {
	Response
	
}
