package yaygo

import (
	"encoding/json"
	"net/http"
)

type GroupApi struct {
	s *Session
}

func newGroupApi(s *Session) *GroupApi {
	return &GroupApi{
		s: s,
	}
}


type AcceptModeratorOfferParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) AcceptModeratorOffer(params *AcceptModeratorOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPut, EndpointChatRoomsV1 + "accept_moderator_offer", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type AcceptOwnershipOfferParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) AcceptOwnershipOffer(params *AcceptOwnershipOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPut, EndpointChatRoomsV1 + "accept_ownership_offer", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type AcceptUserRequestParams struct {
	ID      int `json:"id,omitempty"`
	UserID  int `json:"user_id,omitempty"`
}

func (g *GroupApi) AcceptUserRequest(params *AcceptUserRequestParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "accept_user_request", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type AddRelatedGroupsParams struct {
	ID         int   `json:"id,omitempty"`
	RelatedIDs []int `json:"related_group_id[],omitempty"`
}

func (g *GroupApi) AddRelatedGroups(params *AddRelatedGroupsParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPut, EndpointChatRoomsV1 + "add_related_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type BanUserParams struct {
	ID      int `json:"id,omitempty"`
	UserID  int `json:"userId,omitempty"`
}

func (g *GroupApi) BanUser(params *BanUserParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "ban_user", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CheckUnreadStatusParams struct {
	FromTime int `json:"from_time,omitempty"`
}

func (g *GroupApi) CheckUnreadStatus(params *CheckUnreadStatusParams) (st *UnreadStatusResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "check_unread_status", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateGroupParams struct {
	Topic                            string `json:"topic,omitempty"`
	Description                      string `json:"description,omitempty"`
	IsInviteOnly                     bool   `json:"secret,omitempty"`
	HideReportedPosts                bool   `json:"hide_reported_posts,omitempty"`
	HideConferenceCall               bool   `json:"hide_conference_call,omitempty"`
	IsPrivate                        bool   `json:"is_private,omitempty"`
	IsOnlyForAgeVerifiedUsers        bool   `json:"only_verified_age,omitempty"`
	IsOnlyForPhoneVerifiedUsers      bool   `json:"only_mobile_verified,omitempty"`
	CallTimelineDisplay              bool   `json:"call_timeline_display,omitempty"`
	AllowOwnershipTransfer           bool   `json:"allow_ownership_transfer,omitempty"`
	AllowThreadCreationBy            string `json:"allow_thread_creation_by,omitempty"`
	Gender                           int    `json:"gender,omitempty"`
	GenerationGroupsLimit            int    `json:"generation_groups_limit,omitempty"`
	CatId                            int    `json:"group_category_id,omitempty"`
	CoverImageFileName               string `json:"cover_image_filename,omitempty"`
	IconImageFileName                string `json:"group_icon_filename,omitempty"`
	UUID                             string `json:"uuid,omitempty"`
	APIKey                           string `json:"api_key,omitempty"`
	Timestamp                        int    `json:"timestamp,omitempty"`
	SignedInfo                       string `json:"signed_info,omitempty"`
	SubCatId                         string `json:"sub_category_id,omitempty"`
	HideFromGameEight                bool   `json:"hide_from_game_eight,omitempty"`
	AllowMembersToPostMedia          bool   `json:"allow_members_to_post_image_and_video,omitempty"`
	AllowMembersToPostUrl            bool   `json:"allow_members_to_post_url,omitempty"`
	Guidelines                       string `json:"guidelines,omitempty"`
}

func (g *GroupApi) Create(params *CreateGroupParams) (st *CreateGroupResponse, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV3 + "create_group", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreatePinGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) CreatePinGroup(params *CreatePinGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "create_pin_group", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeclineModeratorOfferParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) DeclineModeratorOffer(params *DeclineModeratorOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "decline_moderator_offer", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeclineOwnershipOfferParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) DeclineOwnershipOffer(params *DeclineOwnershipOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "decline_ownership_offer", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeclineUserRequestParams struct {
	ID      int `json:"id,omitempty"`
	UserID  int `json:"userId,omitempty"`
}

func (g *GroupApi) DeclineUserRequest(params *DeclineUserRequestParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "decline_user_request", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeleteCoverParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) DeleteCover(params *DeleteCoverParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV3 + "delete_cover", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeleteIconParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) DeleteIcon(params *DeleteIconParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV3 + "delete_icon", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeletePinGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) DeletePinGroup(params *DeletePinGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "delete_pin_group", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetBannedMembersParams struct {
	ID   int `json:"id,omitempty"`
	Page int `json:"page,omitempty"`
}

func (g *GroupApi) GetBannedMembers(params *GetBannedMembersParams) (st *UsersResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_banned_members", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetCategoriesParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"number,omitempty"`
}

func (g *GroupApi) GetCategories(params *GetCategoriesParams) (st *GroupCategoriesResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_categories", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetCommunityCampaigParams struct {
	// Define fields as needed
}

func (g *GroupApi) GetCommunityCampaign(params *GetCommunityCampaigParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_community_campaign", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetCreateQuotaParams struct {
	// Define fields as needed
}

func (g *GroupApi) GetCreateQuota(params *GetCreateQuotaParams) (st *CreateQuotaResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_create_quota", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) GetGroup(params *GetGroupParams) (st *GroupResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_group", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupNotificationSettingsParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) GetGroupNotificationSettings(params *GetGroupNotificationSettingsParams) (st *GroupNotificationSettingsResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_group_notification_settings", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetGroupsParams struct {
	CategoryID    int    `json:"group_category_id,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	FromTimestamp int    `json:"from_timestamp,omitempty"`
	SubCategoryID int    `json:"sub_category_id,omitempty"`
}

func (g *GroupApi) GetGroups(params *GetGroupsParams) (st *GroupsResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetInCircleUserLeaderboardParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) GetInCircleUserLeaderboard(params *GetInCircleUserLeaderboardParams) (st *GroupInCircleUserLeaderboardResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_in_circle_user_leaderboard", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetInvitableUsersParams struct {
	ID             int    `json:"group_id,omitempty"`
	FromTimestamp  int    `json:"from_timestamp,omitempty"`
	Nickname       string `json:"user[nickname],omitempty"`
}

func (g *GroupApi) GetInvitableUsers(params *GetInvitableUsersParams) (st *UsersByTimestampResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_invitable_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetJoinedStatusesParams struct {
	IDs []int `json:"ids[],omitempty"`
}

func (g *GroupApi) GetJoinedStatuses(params *GetJoinedStatusesParams) (st map[string]string, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_joined_statuses", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMemberParams struct {
	ID      int `json:"id,omitempty"`
	UserID  int `json:"userId,omitempty"`
}

func (g *GroupApi) GetMember(params *GetMemberParams) (st *GroupUserResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_member", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMembersParams struct {
	ID              int    `json:"id,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Keyword         string `json:"keyword,omitempty"`
	FromID          int    `json:"from_id,omitempty"`
	PageSize        int    `json:"number,omitempty"`
	FromTimestamp   int    `json:"from_timestamp,omitempty"`
	OrderBy         string `json:"order_by,omitempty"`
	FollowedByMe    bool   `json:"followed_by_me,omitempty"`
}

func (g *GroupApi) GetMembers(params *GetMembersParams) (st *GroupUsersResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_members", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMembersObjectParams struct {
	ID      int    `json:"id,omitempty"`
	Number  string `json:"number,omitempty"`
	FromID  int    `json:"from_id,omitempty"`
}

func (g *GroupApi) GetMembersObject(params *GetMembersObjectParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_members_object", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMutedUsersParams struct {
	ID      int     `json:"id,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Cursor  string  `json:"cursor,omitempty"`
	Size    int     `json:"size,omitempty"`
}

func (g *GroupApi) GetMutedUsers(params *GetMutedUsersParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_muted_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetMyGroupsParams struct {
	FromTimestamp int `json:"from_timestamp,omitempty"`
}

func (g *GroupApi) GetMyGroups(params *GetMyGroupsParams) (st *GroupsResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_my_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetOverallGroupLeaderboardParams struct {
	// Define fields as needed
}

func (g *GroupApi) GetOverallGroupLeaderboard(params *GetOverallGroupLeaderboardParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_overall_group_leaderboard", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRelatableGroupsParams struct {
	ID      int    `json:"id,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	From    string `json:"from,omitempty"`
}

func (g *GroupApi) GetRelatableGroups(params *GetRelatableGroupsParams) (st *GroupsRelatedResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_relatable_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRelatedGroupsParams struct {
	ID      int    `json:"id,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	From    string `json:"from,omitempty"`
}

func (g *GroupApi) GetRelatedGroups(params *GetRelatedGroupsParams) (st *GroupsRelatedResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_related_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserGroupsParams struct {
	Page    int `json:"page,omitempty"`
	UserID  int `json:"user_id,omitempty"`
}

func (g *GroupApi) GetUserGroups(params *GetUserGroupsParams) (st *GroupsResponse, err error) {
	resp, err := g.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_user_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type InviteUsersParams struct {
	ID      int   `json:"id,omitempty"`
	UserIDs []int `json:"user_ids,omitempty"`
}

func (g *GroupApi) InviteUsers(params *InviteUsersParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "invite_users", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type JoinGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) Join(params *JoinGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "join_group", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type LeaveGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) Leave(params *LeaveGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "leave_group", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type MuteUserParams struct {
	ID      int `json:"id,omitempty"`
	UserID  int `json:"user_id,omitempty"`
}

func (g *GroupApi) MuteUser(params *MuteUserParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "mute_user", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveModeratorParams struct {
	ID      int `json:"group_id,omitempty"`
	UserID  int `json:"user_id,omitempty"`
}

func (g *GroupApi) RemoveModerator(params *RemoveModeratorParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "remove_moderator", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RemoveRelatedGroupsParams struct {
	ID         int   `json:"id,omitempty"`
	RelatedIDs []int `json:"related_group_id[],omitempty"`
}

func (g *GroupApi) RemoveRelatedGroups(params *RemoveRelatedGroupsParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "remove_related_groups", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ReportGroupParams struct {
	ID                  int    `json:"group_id,omitempty"`
	CategoryID          int    `json:"category_id,omitempty"`
	Reason              string `json:"reason,omitempty"`
	OpponentID          int    `json:"opponent_id,omitempty"`
	ScreenshotFileName1 string `json:"screenshot_filename,omitempty"`
	ScreenshotFileName2 string `json:"screenshot_2_filename,omitempty"`
	ScreenshotFileName3 string `json:"screenshot_3_filename,omitempty"`
	ScreenshotFileName4 string `json:"screenshot_4_filename,omitempty"`
}

func (g *GroupApi) Report(params *ReportGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV3 + "report_group", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SendModeratorOffersParams struct {
	ID          int    `json:"group_id,omitempty"`
	UserIDs     []int  `json:"user_ids,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	APIKey      string `json:"api_key,omitempty"`
	Timestamp   int    `json:"timestamp,omitempty"`
	SignedInfo  string `json:"signed_info,omitempty"`
}

func (g *GroupApi) SendModeratorOffers(params *SendModeratorOffersParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV3 + "send_moderator_offers", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SendOwnershipOfferParams struct {
	ID         int    `json:"id,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	APIKey     string `json:"api_key,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
	SignedInfo string `json:"signed_info,omitempty"`
}

func (g *GroupApi) SendOwnershipOffer(params *SendOwnershipOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV3 + "send_ownership_offer", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SetGroupNotificationSettingsParams struct {
	ID                     int `json:"id,omitempty"`
	IsNewPostSettingOn     int `json:"notification_group_post,omitempty"`
	IsNewMemberSettingOn   int `json:"notification_group_join,omitempty"`
	IsJoinRequestSettingOn int `json:"notification_group_request,omitempty"`
	IsMentionAllSettingOn  int `json:"notification_group_message_tag_all,omitempty"`
}

func (g *GroupApi) SetGroupNotificationSettings(params *SetGroupNotificationSettingsParams) (st *AdditionalSettingsResponse, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV2 + "set_group_notification_settings", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SetTitleParams struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func (g *GroupApi) SetTitle(params *SetTitleParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "set_title", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type TakeoverOwnershipParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) TakeoverOwnership(params *TakeoverOwnershipParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "takeover_ownership", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UnbanUserParams struct {
	ID     int `json:"id,omitempty"`
	UserID int `json:"userId,omitempty"`
}

func (g *GroupApi) UnbanUser(params *UnbanUserParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "unban_user", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UnmuteUserParams struct {
	ID     int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

func (g *GroupApi) UnmuteUser(params *UnmuteUserParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unmute_user", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateGroupParams struct {
	ID                          int    `json:"id,omitempty"`
	Topic                       string `json:"topic,omitempty"`
	Description                 string `json:"description,omitempty"`
	HideReportedPosts           bool   `json:"hide_reported_posts,omitempty"`
	HideConferenceCall          bool   `json:"hide_conference_call,omitempty"`
	IsPrivate                   bool   `json:"is_private,omitempty"`
	IsInviteOnly                bool   `json:"secret,omitempty"`
	IsOnlyForAgeVerifiedUsers   bool   `json:"only_verified_age,omitempty"`
	IsOnlyForPhoneVerifiedUsers bool   `json:"only_mobile_verified,omitempty"`
	CallTimelineDisplay         bool   `json:"call_timeline_display,omitempty"`
	AllowOwnershipTransfer      bool   `json:"allow_ownership_transfer,omitempty"`
	AllowThreadCreationBy       string `json:"allow_thread_creation_by,omitempty"`
	Gender                      int    `json:"gender,omitempty"`
	GenerationGroupsLimit       int    `json:"generation_groups_limit,omitempty"`
	CatId                       int    `json:"group_category_id,omitempty"`
	SubCatId                    string `json:"sub_category_id,omitempty"`
	CoverImageFileName          string `json:"cover_image_filename,omitempty"`
	IconImageFileName           string `json:"group_icon_filename,omitempty"`
	UUID                        string `json:"uuid,omitempty"`
	APIKey                      string `json:"api_key,omitempty"`
	Timestamp                   int    `json:"timestamp,omitempty"`
	SignedInfo                  string `json:"signed_info,omitempty"`
	HideFromGameEight           bool   `json:"hide_from_game_eight,omitempty"`
	AllowMembersToPostMedia     bool   `json:"allow_members_to_post_image_and_video,omitempty"`
	AllowMembersToPostUrl       bool   `json:"allow_members_to_post_url,omitempty"`
	Guidelines                  string `json:"guidelines,omitempty"`
}

func (g *GroupApi) Update(params *UpdateGroupParams) (st *GroupResponse, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV3 + "update_group", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type VisitGroupParams struct {
	ID int `json:"id,omitempty"`
}

func (g *GroupApi) Visit(params *VisitGroupParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPost, EndpointChatRoomsV1 + "visit_group", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type WithdrawModeratorOfferParams struct {
	ID     int `json:"group_id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

func (g *GroupApi) WithdrawModeratorOffer(params *WithdrawModeratorOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPut, EndpointChatRoomsV1 + "withdraw_moderator_offer", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type WithdrawOwnershipOfferParams struct {
	ID     int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

func (g *GroupApi) WithdrawOwnershipOffer(params *WithdrawOwnershipOfferParams) (st *Response, err error) {
	resp, err := g.s.request(http.MethodPut, EndpointChatRoomsV1 + "withdraw_ownership_offer", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
