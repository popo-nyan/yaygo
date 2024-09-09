package yaygo

import (
	"encoding/json"
	"net/http"
)

type PostAPI struct {
	s *Session
}

func newPostAPI(s *Session) *PostAPI {
	return &PostAPI{
		s: s,
	}
}

type AddBookmarkParams struct {
	UserID int `json:"user_id,omitempty"`
	ID     int `json:"id,omitempty"`
}

func (p *PostAPI) AddBookmark(params *AddBookmarkParams) (st *BookmarkPostResponse, err error) {
	resp, err := p.s.request(http.MethodPut, EndpointChatRoomsV1+"add_bookmark", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type AddGroupHighlightPostParams struct {
	GroupID int `json:"group_id,omitempty"`
	PostID  int `json:"post_id,omitempty"`
}

func (p *PostAPI) AddGroupHighlightPost(params *AddGroupHighlightPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPut, EndpointChatRoomsV1+"add_group_highlight_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateGroupCallPostParams struct {
	Text       string `json:"text,omitempty"`
	FontSize   int    `json:"font_size,omitempty"`
	Color      int    `json:"color,omitempty"`
	GroupID    int    `json:"group_id,omitempty"`
	CallType   string `json:"call_type,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	APIKey     string `json:"api_key,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
	SignedInfo string `json:"signed_info,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	GameTitle  string `json:"game_title,omitempty"`
	JoinableBy string `json:"joinable_by,omitempty"`
	// MessageTags    *RequestBody `json:"message_tags,omitempty"`
	ImageFileName  string `json:"attachment_filename,omitempty"`
	ImageFileName2 string `json:"attachment_2_filename,omitempty"`
	ImageFileName3 string `json:"attachment_3_filename,omitempty"`
	ImageFileName4 string `json:"attachment_4_filename,omitempty"`
	ImageFileName5 string `json:"attachment_5_filename,omitempty"`
	ImageFileName6 string `json:"attachment_6_filename,omitempty"`
	ImageFileName7 string `json:"attachment_7_filename,omitempty"`
	ImageFileName8 string `json:"attachment_8_filename,omitempty"`
	ImageFileName9 string `json:"attachment_9_filename,omitempty"`
}

func (p *PostAPI) CreateGroupCallPost(params *CreateGroupCallPostParams) (st *CreatePostResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"create_group_call_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateGroupPinPostParams struct {
	PostID  int `json:"post_id,omitempty"`
	GroupID int `json:"group_id,omitempty"`
}

func (p *PostAPI) CreateGroupPinPost(params *CreateGroupPinPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPut, EndpointChatRoomsV2+"create_group_pin_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreatePinPostParams struct {
	ID int `json:"id,omitempty"`
}

func (p *PostAPI) CreatePinPost(params *CreatePinPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"create_pin_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreatePostParams struct {
	HeaderJwt  string   `json:"header_jwt,omitempty"`
	Text       string   `json:"text,omitempty"`
	FontSize   int      `json:"font_size,omitempty"`
	Color      int      `json:"color,omitempty"`
	InReplyTo  int      `json:"in_reply_to,omitempty"`
	GroupID    int      `json:"group_id,omitempty"`
	PostType   string   `json:"post_type,omitempty"`
	MentionIDs []int    `json:"mention_ids[],omitempty"`
	Choices    []string `json:"choices[],omitempty"`
	// SharedUrl      *RequestBody `json:"shared_url,omitempty"`
	// MessageTags    *RequestBody `json:"message_tags,omitempty"`
	ImageFileName  string `json:"attachment_filename,omitempty"`
	ImageFileName2 string `json:"attachment_2_filename,omitempty"`
	ImageFileName3 string `json:"attachment_3_filename,omitempty"`
	ImageFileName4 string `json:"attachment_4_filename,omitempty"`
	ImageFileName5 string `json:"attachment_5_filename,omitempty"`
	ImageFileName6 string `json:"attachment_6_filename,omitempty"`
	ImageFileName7 string `json:"attachment_7_filename,omitempty"`
	ImageFileName8 string `json:"attachment_8_filename,omitempty"`
	ImageFileName9 string `json:"attachment_9_filename,omitempty"`
	VideoFileName  string `json:"video_file_name,omitempty"`
}

func (p *PostAPI) CreatePost(params *CreatePostParams) (st *Post, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV3+"create_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateRepostParams struct {
	HeaderJwt  string   `json:"header_jwt,omitempty"`
	PostID     int      `json:"post_id,omitempty"`
	Text       string   `json:"text,omitempty"`
	FontSize   int      `json:"font_size,omitempty"`
	Color      int      `json:"color,omitempty"`
	InReplyTo  int      `json:"in_reply_to,omitempty"`
	GroupID    int      `json:"group_id,omitempty"`
	PostType   string   `json:"post_type,omitempty"`
	MentionIDs []int    `json:"mention_ids[],omitempty"`
	Choices    []string `json:"choices[],omitempty"`
	// SharedUrl      *RequestBody `json:"shared_url,omitempty"`
	// MessageTags    *RequestBody `json:"message_tags,omitempty"`
	ImageFileName  string `json:"attachment_filename,omitempty"`
	ImageFileName2 string `json:"attachment_2_filename,omitempty"`
	ImageFileName3 string `json:"attachment_3_filename,omitempty"`
	ImageFileName4 string `json:"attachment_4_filename,omitempty"`
	ImageFileName5 string `json:"attachment_5_filename,omitempty"`
	ImageFileName6 string `json:"attachment_6_filename,omitempty"`
	ImageFileName7 string `json:"attachment_7_filename,omitempty"`
	ImageFileName8 string `json:"attachment_8_filename,omitempty"`
	ImageFileName9 string `json:"attachment_9_filename,omitempty"`
	VideoFileName  string `json:"video_file_name,omitempty"`
}

func (p *PostAPI) CreateRepost(params *CreateRepostParams) (st *CreatePostResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV3+"create_repost", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateSharePostParams struct {
	ShareableType string `json:"shareable_type,omitempty"`
	ShareableID   int    `json:"shareable_id,omitempty"`
	Text          string `json:"text,omitempty"`
	FontSize      int    `json:"font_size,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	APIKey        string `json:"api_key,omitempty"`
	Timestamp     int    `json:"timestamp,omitempty"`
	SignedInfo    string `json:"signed_info,omitempty"`
	GroupID       int    `json:"group_id,omitempty"`
	Color         int    `json:"color,omitempty"`
}

func (p *PostAPI) CreateSharePost(params *CreateSharePostParams) (st *Post, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"create_share_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateThreadPostParams struct {
	HeaderJwt  string   `json:"header_jwt,omitempty"`
	ThreadID   int      `json:"id,omitempty"`
	Text       string   `json:"text,omitempty"`
	FontSize   int      `json:"font_size,omitempty"`
	Color      int      `json:"color,omitempty"`
	InReplyTo  int      `json:"in_reply_to,omitempty"`
	GroupID    int      `json:"group_id,omitempty"`
	PostType   string   `json:"post_type,omitempty"`
	MentionIDs []int    `json:"mention_ids[],omitempty"`
	Choices    []string `json:"choices[],omitempty"`
	// SharedUrl      *RequestBody `json:"shared_url,omitempty"`
	// MessageTags    *RequestBody `json:"message_tags,omitempty"`
	ImageFileName  string `json:"attachment_filename,omitempty"`
	ImageFileName2 string `json:"attachment_2_filename,omitempty"`
	ImageFileName3 string `json:"attachment_3_filename,omitempty"`
	ImageFileName4 string `json:"attachment_4_filename,omitempty"`
	ImageFileName5 string `json:"attachment_5_filename,omitempty"`
	ImageFileName6 string `json:"attachment_6_filename,omitempty"`
	ImageFileName7 string `json:"attachment_7_filename,omitempty"`
	ImageFileName8 string `json:"attachment_8_filename,omitempty"`
	ImageFileName9 string `json:"attachment_9_filename,omitempty"`
	VideoFileName  string `json:"video_file_name,omitempty"`
}

func (p *PostAPI) CreateThreadPost(params *CreateThreadPostParams) (st *Post, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"create_thread_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeleteAllPostParams struct{}

func (p *PostAPI) DeleteAllPost(params *DeleteAllPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"delete_all_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeleteGroupPinPostParams struct {
	GroupID int `json:"group_id,omitempty"`
}

func (p *PostAPI) DeleteGroupPinPost(params *DeleteGroupPinPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodDelete, EndpointChatRoomsV2+"delete_group_pin_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeletePinPostParams struct {
	ID int `json:"id,omitempty"`
}

func (p *PostAPI) DeletePinPost(params *DeletePinPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodDelete, EndpointChatRoomsV1+"delete_pin_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetBookmarkParams struct {
	UserID int    `json:"user_id,omitempty"`
	From   string `json:"from,omitempty"`
}

func (p *PostAPI) GetBookmark(params *GetBookmarkParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV1+"get_book_mark", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetCallTimelineParams struct {
	GroupID                int    `json:"group_id,omitempty"`
	FromTimestamp          int    `json:"from_timestamp,omitempty"`
	Number                 int    `json:"number,omitempty"`
	CategoryID             int    `json:"category_id,omitempty"`
	CallType               string `json:"call_type,omitempty"`
	IncludeCircleCall      bool   `json:"include_circle_call,omitempty"`
	CrossGeneration        bool   `json:"cross_generation,omitempty"`
	ExcludeRecentGomimushi bool   `json:"exclude_recent_gomimushi,omitempty"`
}

func (p *PostAPI) GetCallTimeline(params *GetCallTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_call_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetConversationParams struct {
	ID       int  `json:"id,omitempty"`
	GroupID  int  `json:"group_id,omitempty"`
	ThreadID int  `json:"thread_id,omitempty"`
	Number   int  `json:"number,omitempty"`
	FromID   int  `json:"from_post_id,omitempty"`
	Reverse  bool `json:"reverse,omitempty"`
}

func (p *PostAPI) GetConversation(params *GetConversationParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_conversation", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetConversationRootPostsParams struct {
	ConversationIDs []int `json:"ids[],omitempty"`
}

func (p *PostAPI) GetConversationRootPosts(params *GetConversationRootPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_conversation_root_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFollowingCallTimelineParams struct {
	FromTimestamp          int    `json:"from_timestamp,omitempty"`
	Number                 int    `json:"number,omitempty"`
	CategoryID             int    `json:"category_id,omitempty"`
	CallType               string `json:"call_type,omitempty"`
	IncludeCircleCall      bool   `json:"include_circle_call,omitempty"`
	ExcludeRecentGomimushi bool   `json:"exclude_recent_gomimushi,omitempty"`
}

func (p *PostAPI) GetFollowingCallTimeline(params *GetFollowingCallTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_following_call_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFollowingTimelineParams struct {
	From                            string `json:"from,omitempty"`
	FromID                          int    `json:"from_post_id,omitempty"`
	OnlyRoot                        bool   `json:"only_root,omitempty"`
	OrderBy                         string `json:"order_by,omitempty"`
	Number                          int    `json:"number,omitempty"`
	MaxHotPosts                     int    `json:"mxn,omitempty"`
	IsReduceSelfieForHotPostEnabled bool   `json:"reduce_selfie,omitempty"`
	IsCustomGenerationRangeEnabled  bool   `json:"custom_generation_range,omitempty"`
}

func (p *PostAPI) GetFollowingTimeline(params *GetFollowingTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_following_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupHighlightPostsParams struct {
	GroupID int `json:"group_id,omitempty"`
	From    int `json:"from,omitempty"`
	Number  int `json:"number,omitempty"`
}

func (p *PostAPI) GetGroupHighlightPosts(params *GetGroupHighlightPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV1+"get_group_highlight", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupSearchPostsParams struct {
	ID              int    `json:"id,omitempty"`
	Keyword         string `json:"keyword,omitempty"`
	FromID          int    `json:"from_post_id,omitempty"`
	Limit           int    `json:"number,omitempty"`
	OnlyThreadPosts bool   `json:"only_thread_posts,omitempty"`
}

func (p *PostAPI) GetGroupSearchPosts(params *GetGroupSearchPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"getG_group_search", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupTimelineParams struct {
	GroupID  int    `json:"group_id,omitempty"`
	FromID   int    `json:"from_post_id,omitempty"`
	Reverse  bool   `json:"reverse,omitempty"`
	PostType string `json:"post_type,omitempty"`
	Number   int    `json:"number,omitempty"`
	NoReply  bool   `json:"only_root,omitempty"`
}

func (p *PostAPI) GetGroupTimeline(params *GetGroupTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_group_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetHashtagTimelineParams struct {
	Tag    string `json:"tag,omitempty"`
	FromID int    `json:"from_post_id,omitempty"`
	Limit  int    `json:"number,omitempty"`
}

func (p *PostAPI) GetHashtagTimeline(params *GetHashtagTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_hashtag_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMyPostsParams struct {
	FromID           int  `json:"from_post_id,omitempty"`
	Number           int  `json:"number,omitempty"`
	IncludeGroupPost bool `json:"include_group_post,omitempty"`
}

func (p *PostAPI) GetMyPosts(params *GetMyPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_my_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPostParams struct {
	ID           int    `json:"id,omitempty"`
	CacheControl string `json:"cache_control,omitempty"`
}

func (p *PostAPI) GetPost(params *GetPostParams) (st *PostResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPostLikersParams struct {
	ID     int `json:"id,omitempty"`
	FromID int `json:"from_id,omitempty"`
}

func (p *PostAPI) GetPostLikers(params *GetPostLikersParams) (st *PostLikersResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV1+"get_post_likers", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPostRepostsParams struct {
	ID     int `json:"id,omitempty"`
	FromID int `json:"from_post_id,omitempty"`
}

func (p *PostAPI) GetPostReposts(params *GetPostRepostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_post_reposts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPostsParams struct {
	PostIDs []int `json:"post_ids[],omitempty"`
}

func (p *PostAPI) GetPosts(params *GetPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRecentEngagementsPostsParams struct {
	Number int `json:"number,omitempty"`
}

func (p *PostAPI) GetRecentEngagementsPosts(params *GetRecentEngagementsPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_recent_engagements_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRecommendedPostTagsParams struct {
	Tag                    string `json:"tag,omitempty"`
	ShouldSaveRecentSearch bool   `json:"save_recent_search,omitempty"`
}

func (p *PostAPI) GetRecommendedPostTags(params *GetRecommendedPostTagsParams) (st *PostTagsResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"get_recommended_post_tags", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRecommendedPostsParams struct {
	ExperimentNumber       int  `json:"experiment_num,omitempty"`
	VariantNumber          int  `json:"variant_num,omitempty"`
	Number                 int  `json:"number,omitempty"`
	ShouldFilterByInterest bool `json:"shared_interest_categories,omitempty"`
}

func (p *PostAPI) GetRecommendedPosts(params *GetRecommendedPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_recommended_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetSearchPostsParams struct {
	Keyword        string `json:"keyword,omitempty"`
	PostOwnerScope string `json:"post_owner_scope,omitempty"`
	OnlyMedia      bool   `json:"only_media,omitempty"`
	FromID         int    `json:"from_post_id,omitempty"`
	Limit          int    `json:"number,omitempty"`
}

func (p *PostAPI) GetSearchPosts(params *GetSearchPostsParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_search_posts", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetTimelineParams struct {
	NoReply                          string `json:"noreply_mode,omitempty"`
	OrderBy                          string `json:"order_by,omitempty"`
	IsExperimentOlderAgeRulesEnabled bool   `json:"experiment_older_age_rules,omitempty"`
	IsInterestsFilterEnabled         bool   `json:"shared_interest_categories,omitempty"`
	From                             string `json:"from,omitempty"`
	FromID                           int    `json:"from_post_id,omitempty"`
	Number                           int    `json:"number,omitempty"`
	MaxHotPosts                      int    `json:"mxn,omitempty"`
	ExperimentNumber                 int    `json:"en,omitempty"`
	VariantNumber                    int    `json:"vn,omitempty"`
	IsReduceSelfieForHotPostEnabled  bool   `json:"reduce_selfie,omitempty"`
	IsCustomGenerationRangeEnabled   bool   `json:"custom_generation_range,omitempty"`
}

func (p *PostAPI) GetTimeline(params *GetTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUrlMetadataParams struct {
	URL string `json:"url,omitempty"`
}

func (p *PostAPI) GetUrlMetadata(params *GetUrlMetadataParams) (st *SharedUrl, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_url_metadata", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserTimelineParams struct {
	UserID   int    `json:"user_id,omitempty"`
	FromID   int    `json:"from_post_id,omitempty"`
	PostType string `json:"post_type,omitempty"`
	Number   int    `json:"number,omitempty"`
}

func (p *PostAPI) GetUserTimeline(params *GetUserTimelineParams) (st *PostsResponse, err error) {
	resp, err := p.s.request(http.MethodGet, EndpointChatRoomsV2+"get_user_timeline", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type LikePostsParams struct {
	PostIDs []int `json:"post_ids[],omitempty"`
}

func (p *PostAPI) LikePosts(params *LikePostsParams) (st *LikePostsResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"like_posts", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveBookmarkParams struct {
	UserID int `json:"user_id,omitempty"`
	ID     int `json:"id,omitempty"`
}

func (p *PostAPI) RemoveBookmark(params *RemoveBookmarkParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodDelete, EndpointChatRoomsV1+"remove_bookmark", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveGroupHighlightPostParams struct {
	GroupID int `json:"group_id,omitempty"`
}

func (p *PostAPI) RemoveGroupHighlightPost(params *RemoveGroupHighlightPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodDelete, EndpointChatRoomsV1+"remove_group_highlight_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemovePostsParams struct {
	PostsIDs []int `json:"posts_ids[],omitempty"`
}

func (p *PostAPI) RemovePosts(params *RemovePostsParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"remove_posts", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ReportPostParams struct {
	PostID              int    `json:"post_id,omitempty"`
	CategoryID          int    `json:"category_id,omitempty"`
	Reason              string `json:"reason,omitempty"`
	OpponentID          int    `json:"opponent_id,omitempty"`
	ScreenshotFileName  string `json:"screenshot_filename,omitempty"`
	ScreenshotFileName2 string `json:"screenshot_2_filename,omitempty"`
	ScreenshotFileName3 string `json:"screenshot_3_filename,omitempty"`
	ScreenshotFileName4 string `json:"screenshot_4_filename,omitempty"`
}

func (p *PostAPI) ReportPost(params *ReportPostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV3+"report_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UnlikePostParams struct {
	ID int `json:"id,omitempty"`
}

func (p *PostAPI) UnlikePost(params *UnlikePostParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodDelete, EndpointChatRoomsV1+"unlike_post", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdatePostParams struct {
	ID       int    `json:"id,omitempty"`
	Text     string `json:"text,omitempty"`
	FontSize int    `json:"font_size,omitempty"`
	Color    int    `json:"color,omitempty"`
	// MessageTags *RequestBody `json:"message_tags,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	APIKey     string `json:"api_key,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
	SignedInfo string `json:"signed_info,omitempty"`
}

func (p *PostAPI) UpdatePost(params *UpdatePostParams) (st *Post, err error) {
	resp, err := p.s.request(http.MethodPut, EndpointChatRoomsV3+"update_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateRecommendationFeedbackParams struct {
	ID               int    `json:"id,omitempty"`
	ExperimentNumber int    `json:"experiment_num,omitempty"`
	VariantNumber    int    `json:"variant_num,omitempty"`
	FeedbackResult   string `json:"feedback_result,omitempty"`
}

func (p *PostAPI) UpdateRecommendationFeedback(params *UpdateRecommendationFeedbackParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"update_recommendation_feedback", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ValidatePostParams struct {
	Text     string `json:"text,omitempty"`
	GroupID  int    `json:"group_id,omitempty"`
	ThreadID int    `json:"thread_id,omitempty"`
}

func (p *PostAPI) ValidatePost(params *ValidatePostParams) (st *ValidationPostResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"validate_post", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ViewVideoParams struct {
	VideoID int `json:"videoId,omitempty"`
}

func (p *PostAPI) ViewVideo(params *ViewVideoParams) (st *Response, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV1+"view_video", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type VoteSurveyParams struct {
	ID       int `json:"id,omitempty"`
	ChoiceID int `json:"choice_id,omitempty"`
}

func (p *PostAPI) VoteSurvey(params *VoteSurveyParams) (st *VoteSurveyResponse, err error) {
	resp, err := p.s.request(http.MethodPost, EndpointChatRoomsV2+"vote_survey", nil, params, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
