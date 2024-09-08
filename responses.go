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
	RoomID int `json:"room_id"`
}

type FollowUsersResponse struct {
	Response
	LastFollowID int    `json:"last_follow_id"`
	Users        []User `json:"users"`
}

type GifsDataResponse struct {
	Response
}

type ChatRoomsResponse struct {
	Response
	// PinnedChatRooms []ChatRoom `json:"pinned_chat_rooms"`
	// ChatRooms       []ChatRoom `json:"chat_rooms"`
	NextPageValue int `json:"next_page_value"`
}

type MessagesResponse struct {
	Response
	// Messages []Message `json:"messages"`
}

type AdditionalSettingsResponse struct {
	Response
	// Settings Settings `json:"settings"`
}

type ChatRoomResponse struct {
	Response
	// ChatRoom ChatRoom `json:"chat_room"`
}

type StickerPacksResponse struct {
	Response
	// StickerPacks []StickerPack `json:"sticker_packs"`
}

type TotalChatRequestResponse struct {
	Response
	Total int `json:"total"`
}

type NotificationSettingResponse struct {
	Response
	// Setting UserSetting `json:"setting"`
}

type CallActionSignatureResponse struct {
	Response
	// SignaturePayload SignaturePayload `json:"signature_payload"`
}

type PostResponse struct {
	Response
	// Post Post `json:"post"`
}

type BGMsResponse struct {
	Response
	// BGM []BGM `json:"bgm"`
}

type ConferenceCallResponse struct {
	Response
	// ConferenceCall         ConferenceCall `json:"conference_call"`
	ConferenceCallUserUUID string `json:"conference_call_user_uuid"`
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
	RoomURL     string `json:"room_url"`
}

type GamesResponse struct {
	Response
	// Games  []Game `json:"games"`
	FromID int `json:"from_id"`
}

type GenresResponse struct {
	Response
	// Genres        []Genre `json:"genres"`
	NextPageValue int `json:"next_page_value"`
}

type PostsResponse struct {
	Response
	NextPageValue string `json:"mext_page_value"`
	// Posts           []Post `json:"posts"`
	// PinnedPosts     []Post `json:"pinned_posts"`
	HasMoreHotPosts bool `json:"has_more_hot_posts"`
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
	// Reviews       []Review `json:"review"`
	// PinnedReviews []Review `json:"pinned_reviews"`
}

type CreateGroupResponse struct {
	Response
	GroupID int `json:"group_id"`
}

type UsersResponse struct {
	Response
	Users         []User `json:"users"`
	NextPageValue string `json:"next_page_value"`
}

type CreateQuotaResponse struct {
	Response
	// Create CreateGroupQuota `json:"create"`
}

type GroupCategoriesResponse struct {
	Response
	// GroupCategories []GroupCategory `json:"group_categories"`
}

type GroupResponse struct {
	Response
	// Group Group `json:"group"`
}

type GroupNotificationSettingsResponse struct {
	Response
	// Setting Setting `json:"setting"`
}

type GroupsResponse struct {
	Response
	// PinnedGroups []Group `json:"pinned_groups"`
	// Groups       []Group `json:"groups"`
}

type GroupUsersResponse struct {
	Response
	// GroupUsers []GroupUser `json:"group_users"`
}

type GroupUserResponse struct {
	Response
	// GroupUser GroupUser `json:"group_user"`
}

type GroupsRelatedResponse struct {
	Response
	// Groups        []Group `json:"groups"`
	NextPageValue string `json:"next_page_value"`
}

type GroupInCircleUserLeaderboardResponse struct {
	Response
	// UserLeaderboard []UserRank `json:"user_leaderboard"`
}

type ThreadInfo struct {
	Response
	ID    int    `json:"id"`
	Title string `json:"title"`
	Owner User   `json:"owner"`
	// LastPost    Post   `json:"last_post"`
	UnreadCount int    `json:"unread_count"`
	PostsCount  int    `json:"posts_count"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	ThreadIcon  string `json:"thread_icon"`
	IsJoined    bool   `json:"is_joined"`
	NewUpdates  bool   `json:"new_updates"`
}

type GroupThreadListResponse struct {
	Response
	Threads       []ThreadInfo `json:"threads"`
	NextPageValue string       `json:"next_page_value"`
}

type EmailVerificationPresignedURLResponse struct {
	Response
	URL string `json:"url"`
}

type PresignedURLsResponse struct {
	Response
	// PresignedURLs PresignedURL `json:"presigned_urls"`
}

type IDCheckerPresignedURLResponse struct {
	Response
	PresignedURL string `json:"presigned_url"`
}

type PresignedURLResponse struct {
	Response
	PresignedURL string `json:"presigned_url"`
}

type PolicyAgreementsResponse struct {
	Response
	LatestPrivacyPolicyAgreed bool `json:"latest_privacy_policy_agreed"`
	LatestTermsOfUseAgreed    bool `json:"latest_terms_of_use_agreed"`
}

type PromotionsResponse struct {
	Response
	// Promotions []Promotion `json:"promotions"`
}

type VipGameRewardURLResponse struct {
	Response
	URL string `json:"url"`
}

type WebSocketTokenResponse struct {
	Response
	Token string `json:"token"`
}

type VerifyDeviceResponse struct {
	Response
	Verified   bool   `json:"verified"`
	VerifiedAt string `json:"verified_at"`
}

type MessageResponse struct {
	Response
	// ConferenceCall ConferenceCall `json:"conference_call"`
	ID int `json:"id"`
}


type CreateUserResponse struct {
	Response
	Id           int    `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type ActiveFollowingsResponse struct {
	Response
	LastLoggedInAt int    `json:"last_logged_in_at"`
	Users          []User `json:"users"`
}

type AppReviewStatusResponse struct {
	Response
	IsAppReviewed bool `json:"is_app_reviewed"`
}

type ContactStatusResponse struct {
	Response
	// 未指定のフィールドにはコメントがあるので、そのままにしています。
}

type DefaultSettingsResponse struct {
	Response
	TimelineSettings TimelineSettings `json:"timeline_settings"`
}

type FollowRecommendationsResponse struct {
	Response
	Total int    `json:"total"`
	Users []User `json:"users"`
	Next  int    `json:"next"`
}

type FollowRequestCountResponse struct {
	Response
	UsersCount int `json:"users_count"`
}

type FootprintsResponse struct {
	Response
	Footprints []Footprint `json:"footprints"`
}

type HimaUsersResponse struct {
	Response
	HimaUsers []UserWrapper `json:"hima_users"`
}

type UserCustomDefinitionsResponse struct {
	Response
	Age            int    `json:"age"`
	FollowersCount int    `json:"followers_count"`
	FollowingsCount int   `json:"followings_count"`
	CreatedAt      int    `json:"created_at"`
	LastLoggedInAt int    `json:"last_logged_in_at"`
	Status         string `json:"status"`
	ReportedCount  int    `json:"reported_count"`
}

type UserEmailResponse struct {
	Response
	Email string `json:"email"`
}
