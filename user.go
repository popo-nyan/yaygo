package yaygo

import (
	"encoding/json"
	"net/http"
)

type UserAPI struct {
	s *Session
}

func newUserAPI(s *Session) *UserAPI {
	return &UserAPI{
		s: s,
	}
}

func (u *UserAPI) GetTimestamp() (st *UserTimestampResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointUsersTimestamp(), nil, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (u *UserAPI) GetUser(uID int) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointUsers(uID), nil, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}




type ChangeInvitationCodeParams struct {
	Code string `json:"code,omitempty"`
}

func (u *UserAPI) ChangeInvitationCode(params *ChangeInvitationCodeParams) (st interface{}, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "change_invitation_code", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type CreateUserParams struct {
	AppVersion         string  `json:"app_version,omitempty"`
	Timestamp          int   `json:"timestamp,omitempty"`
	APIKey             string  `json:"api_key,omitempty"`
	SignedVersion      string  `json:"signed_version,omitempty"`
	SignedInfo         string  `json:"signed_info,omitempty"`
	UUID               string  `json:"uuid,omitempty"`
	Nickname           string  `json:"nickname,omitempty"`
	BirthDate          string  `json:"birth_date,omitempty"`
	Gender             int     `json:"gender,omitempty"`
	CountryCode        string  `json:"country_code,omitempty"`
	Biography          string `json:"biography,omitempty"`
	Prefecture         string `json:"prefecture,omitempty"`
	ProfileIconFileName string `json:"profile_icon_filename,omitempty"`
	CoverImageFileName string `json:"cover_image_filename,omitempty"`
	SnsInfoRequest     *SignUpSnsInfoRequest `json:"sns_info,omitempty"`
	Email              string `json:"email,omitempty"`
	Password           string `json:"password,omitempty"`
	EmailGrantToken    string `json:"email_grant_token,omitempty"`
	ExperimentNumber   int    `json:"en,omitempty"`
	VariantNumber      int    `json:"vn,omitempty"`
	ReferralCode       string  `json:"referral_code,omitempty"`
}

func (u *UserAPI) CreateUser(params *CreateUserParams) (st *CreateUserResponse, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV3 + "create_user", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) DeleteContactFriends() (st *Response, err error) {
	resp, err := u.s.request(http.MethodDelete, EndpointChatRoomsV1 + "delete_contact_friends", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DeleteFootprintParams struct {
	UserID     int `json:"user_id"`
	FootprintID int `json:"footprint_id"`
}

func (u *UserAPI) DeleteFootprint(params *DeleteFootprintParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodDelete, EndpointChatRoomsV2 + "delete_footprint", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type DestroyUserParams struct {
	UUID         string `json:"uuid,omitempty"`
	APIKey       string `json:"api_key,omitempty"`
	Timestamp    int  `json:"timestamp,omitempty"`
	SignedInfo   string `json:"signed_info,omitempty"`
}

func (u *UserAPI) DestroyUser(params *DestroyUserParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "destroy_user", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type Disable2FaParams struct {
	Code string `json:"code,omitempty"`
}

func (u *UserAPI) Disable2Fa(params *Disable2FaParams) (st interface{}, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "disable_2fa", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type Enable2FaParams struct {
	Code string `json:"code,omitempty"`
}

func (u *UserAPI) Enable2Fa(params *Enable2FaParams) (st *TwoStepAuthEnabledResponse, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "enable_2fa", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type FollowUserParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) FollowUser(params *FollowUserParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "follow_user", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type FollowUsersParams struct {
	UserIDs []int `json:"user_ids[],omitempty"`
}

func (u *UserAPI) FollowUsers(params *FollowUsersParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "follow_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) Get2FaRequestInfo(continuation *Continuation) (st *TwoStepAuthRequestInfoResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_2fa_request_info", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetActiveFollowingsParams struct {
	OnlyOnline      bool   `json:"only_online,omitempty"`
	FromLoggedInAt  int `json:"from_loggedin_at,omitempty"`
}

func (u *UserAPI) GetActiveFollowings(params *GetActiveFollowingsParams) (st *ActiveFollowingsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_active_followings", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetAdditionalSettings() (st *AdditionalSettingsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_additional_settings", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetAppReviewStatusParams struct {
	UUID string `json:"uuid,omitempty"`
}

func (u *UserAPI) GetAppReviewStatus(params *GetAppReviewStatusParams) (st *AppReviewStatusResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_app_review_status", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetContactStatusParams struct {
	MobileNumbers []string `json:"mobile_numbers[],omitempty"`
}

func (u *UserAPI) GetContactStatus(params *GetContactStatusParams) (st *ContactStatusResponse, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "get_contact_status", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetDefaultSettings() (st *DefaultSettingsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_default_settings", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFollowRecommendationsParams struct {
	FromTimestamp int  `json:"from_timestamp,omitempty"`
	Number        int    `json:"number,omitempty"`
	Sources       []string `json:"sources[],omitempty"`
}

func (u *UserAPI) GetFollowRecommendations(params *GetFollowRecommendationsParams) (st *FollowRecommendationsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_follow_recommendations", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetFollowRequest(params *GetFollowRequestParams) (st *UsersByTimestampResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_follow_request", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetFollowRequestCount() (st *FollowRequestCountResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_follow_request_count", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFollowingUsersBornParams struct {
	BirthdateSeconds int `json:"birthdate,omitempty"`
}

func (u *UserAPI) GetFollowingUsersBorn(params *GetFollowingUsersBornParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_following_users_born", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFootprintsParams struct {
	FromID   int `json:"from_id,omitempty"`
	Number   int   `json:"number,omitempty"`
	Mode     string `json:"mode,omitempty"`
}

func (u *UserAPI) GetFootprints(params *GetFootprintsParams) (st *FootprintsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_footprints", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFreshUserParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) GetFreshUser(params *GetFreshUserParams) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_fresh_user", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetHimaUsersParams struct {
	FromHimaID int `json:"from_hima_id,omitempty"`
	Number     int   `json:"number,omitempty"`
}

func (u *UserAPI) GetHimaUsers(params *GetHimaUsersParams) (st *HimaUsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_hima_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetInvitationCode(continuation *Continuation) (st *InvitationCodeResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_invitation_code", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetRecommendedUsersToFollowForProfileParams struct {
	UserID int  `json:"id,omitempty"`
	Number  int  `json:"number,omitempty"`
	Page    int  `json:"page,omitempty"`
}

func (u *UserAPI) GetRecommendedUsersToFollowForProfile(params *GetRecommendedUsersToFollowForProfileParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_recommended_users_to_follow_for_profile", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetRefreshCounterRequests() (st *RefreshCounterRequestsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_refresh_counter_requests", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetTimestamp() (st *UserTimestampResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_timestamp", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) GetUser(params *GetUserParams) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_user", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetUserCustomDefinitions() (st *UserCustomDefinitionsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_user_custom_definitions", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserEmailParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) GetUserEmail(params *GetUserEmailParams) (st *UserEmailResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_user_email", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserFollowersParams struct {
	UserID        int   `json:"id,omitempty"`
	FromFollowID  int  `json:"from_follow_id,omitempty"`
	FollowedByMe  bool   `json:"followed_by_me,omitempty"`
	Nickname      string `json:"user[nickname],omitempty"`
}

func (u *UserAPI) GetUserFollowers(params *GetUserFollowersParams) (st *FollowUsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_user_followers", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserFollowingsParams struct {
	UserID        int                   `json:"id,omitempty"`
	FromFollowID  int                  `json:"from_follow_id,omitempty"`
	FromTimestamp int                  `json:"from_timestamp,omitempty"`
	OrderBy       string                 `json:"order_by,omitempty"`
	Request       *SearchUsersRequest      `json:"request,omitempty"`
}

func (u *UserAPI) GetUserFollowings(params *GetUserFollowingsParams) (st *FollowUsersResponse, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "get_user_followings", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserFromQrParams struct {
	QR string `json:"qr,omitempty"`
}

func (u *UserAPI) GetUserFromQr(params *GetUserFromQrParams) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_user_from_qr", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) GetUserInterests(continuation *Continuation) (st *UserInterestsResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_user_interests", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserWithCallUserIdParams struct {
	CallID     int  `json:"callId,omitempty"`
	CallUserID string `json:"callUserId,omitempty"`
}

func (u *UserAPI) GetUserWithCallUserId(params *GetUserWithCallUserIdParams) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_user_with_call_user_id", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUserWithoutLeavingFootprintParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) GetUserWithoutLeavingFootprint(params *GetUserWithoutLeavingFootprintParams) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_user_without_leaving_footprint", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUsersParams struct {
	HeaderJwt string   `json:"X-Jwt,omitempty"`
	IDs       []int `json:"user_ids[],omitempty"`
}

func (u *UserAPI) GetUsers(params *GetUsersParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetUsersFromUuidParams struct {
	UUID string `json:"uuid,omitempty"`
}

func (u *UserAPI) GetUsersFromUuid(params *GetUsersFromUuidParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_users_from_uuid", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RecordAppReviewStatusParams struct {
	UUID string `json:"uuid,omitempty"`
}

func (u *UserAPI) RecordAppReviewStatus(params *RecordAppReviewStatusParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "record_app_review_status", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ReduceKentaPenaltyParams struct {
	UserID          int   `json:"id,omitempty"`
	AppVersion      string  `json:"app_version,omitempty"`
	Timestamp       int   `json:"timestamp,omitempty"`
	APIKey          string  `json:"api_key,omitempty"`
	SignedVersion   string  `json:"signed_version,omitempty"`
	SignedInfo      string  `json:"signed_info,omitempty"`
	UUID            string  `json:"uuid,omitempty"`
}

func (u *UserAPI) ReduceKentaPenalty(params *ReduceKentaPenaltyParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV3 + "reduce_kenta_penalty", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RefreshCounterParams struct {
	Counter string `json:"counter,omitempty"`
}

func (u *UserAPI) RefreshCounter(params *RefreshCounterParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "refresh_counter", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type RegisterInvitationCodeParams struct {
	Code string `json:"code,omitempty"`
}

func (u *UserAPI) RegisterInvitationCode(params *RegisterInvitationCodeParams) (st interface{}, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "register_invitation_code", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) RemoveUserAvatar() (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "remove_user_avatar", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) RemoveUserCover() (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "remove_user_cover", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ReportUserParams struct {
	UserID           int   `json:"user_id,omitempty"`
	CategoryID       int   `json:"category_id,omitempty"`
	Reason           string `json:"reason,omitempty"`
	ScreenshotFileName1 string `json:"screenshot_filename,omitempty"`
	ScreenshotFileName2 string `json:"screenshot_2_filename,omitempty"`
	ScreenshotFileName3 string `json:"screenshot_3_filename,omitempty"`
	ScreenshotFileName4 string `json:"screenshot_4_filename,omitempty"`
}

func (u *UserAPI) ReportUser(params *ReportUserParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV3 + "report_user", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type ResetPasswordParams struct {
	Email            string `json:"email,omitempty"`
	EmailGrantToken  string `json:"email_grant_token,omitempty"`
	Password         string `json:"password,omitempty"`
}

func (u *UserAPI) ResetPassword(params *ResetPasswordParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "reset_password", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SearchLobiUsersParams struct {
	Nickname string `json:"nickname,omitempty"`
	Number   int    `json:"number,omitempty"`
	From     string `json:"from,omitempty"`
}

func (u *UserAPI) SearchLobiUsers(params *SearchLobiUsersParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "search_lobi_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SearchUsersParams struct {
	Gender           int    `json:"gender,omitempty"`
	Nickname         string `json:"nickname,omitempty"`
	Title            string `json:"title,omitempty"`
	Biography        string `json:"biography,omitempty"`
	FromTimestamp    int  `json:"from_timestamp,omitempty"`
	SimilarAge       bool   `json:"similar_age,omitempty"`
	NoRecentPenalty   bool   `json:"not_recent_gomimushi,omitempty"`
	RecentlyCreated   bool   `json:"recently_created,omitempty"`
	SamePrefecture    bool   `json:"same_prefecture,omitempty"`
	ShouldSaveRecentSearch bool `json:"save_recent_search,omitempty"`
}

func (u *UserAPI) SearchUsers(params *SearchUsersParams) (st *UsersResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointChatRoomsV1 + "search_users", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SetAdditionalSettingEnabledParams struct {
	Mode   string `json:"mode,omitempty"`
	IsOn   int    `json:"on,omitempty"`
}

func (u *UserAPI) SetAdditionalSettingEnabled(params *SetAdditionalSettingEnabledParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "set_additional_setting_enabled", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SetFollowPermissionEnabledParams struct {
	Nickname      string  `json:"nickname,omitempty"`
	IsPrivate     bool    `json:"is_private,omitempty"`
	UUID          string  `json:"uuid,omitempty"`
	APIKey        string  `json:"api_key,omitempty"`
	Timestamp     int   `json:"timestamp,omitempty"`
	SignedInfo    string  `json:"signed_info,omitempty"`
}

func (u *UserAPI) SetFollowPermissionEnabled(params *SetFollowPermissionEnabledParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "set_follow_permission_enabled", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type SetSettingFollowRecommendationEnabledParams struct {
	On bool `json:"on"`
}

func (u *UserAPI) SetSettingFollowRecommendationEnabled(params *SetSettingFollowRecommendationEnabledParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "set_setting_follow_recommendation_enabled", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type TakeActionFollowRequestParams struct {
	TargetUserID int  `json:"target_id,omitempty"`
	Action       string `json:"action,omitempty"`
}

func (u *UserAPI) TakeActionFollowRequest(params *TakeActionFollowRequestParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "take_action_follow_request", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) TurnOnHima() (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "turn_on_hima", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UnfollowUserParams struct {
	UserID int `json:"id,omitempty"`
}

func (u *UserAPI) UnfollowUser(params *UnfollowUserParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV2 + "unfollow_user", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateInviteContactStatusParams struct {
	MobileNumber string `json:"mobile_number,omitempty"`
}

func (u *UserAPI) UpdateInviteContactStatus(params *UpdateInviteContactStatusParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "update_invite_contact_status", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateLanguageParams struct {
	UUID        string `json:"uuid,omitempty"`
	APIKey      string `json:"api_key,omitempty"`
	Timestamp   int  `json:"timestamp,omitempty"`
	SignedInfo  string `json:"signed_info,omitempty"`
	Language    string `json:"language,omitempty"`
}

func (u *UserAPI) UpdateLanguage(params *UpdateLanguageParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "update_language", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateUserParams struct {
	Nickname             string  `json:"nickname,omitempty"`
	Biography            string `json:"biography,omitempty"`
	Prefecture           string `json:"prefecture,omitempty"`
	Gender               int    `json:"gender,omitempty"`
	CountryCode          string `json:"country_code,omitempty"`
	ProfileIconFileName  string `json:"profile_icon_filename,omitempty"`
	CoverImageFileName   string `json:"cover_image_filename,omitempty"`
	Username             string `json:"username,omitempty"`
	UUID                 string  `json:"uuid,omitempty"`
	APIKey               string  `json:"api_key,omitempty"`
	Timestamp            int   `json:"timestamp,omitempty"`
	SignedInfo           string  `json:"signed_info,omitempty"`
}

func (u *UserAPI) UpdateUser(params *UpdateUserParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV3 + "update_user", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UpdateUserInterestsParams struct {
	CommonIdsRequest `json:"commonIdsRequest,omitempty"`
}

func (u *UserAPI) UpdateUserInterests(params *UpdateUserInterestsParams) (st interface{}, err error) {
	resp, err := u.s.request(http.MethodPut, EndpointChatRoomsV1 + "update_user_interests", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UploadContactsRequestParams struct {
	Request *UploadContactsRequest `json:"request,omitempty"`
}

func (u *UserAPI) UploadContactsFriends(params *UploadContactsRequestParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "upload_contacts", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type UploadTwitterFriendIdsParams struct {
	IDs []string `json:"twitter_friend_ids[],omitempty"`
}

func (u *UserAPI) UploadTwitterFriendIds(params *UploadTwitterFriendIdsParams) (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "upload_twitter_friend_ids", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (u *UserAPI) UserAlive() (st *Response, err error) {
	resp, err := u.s.request(http.MethodPost, EndpointChatRoomsV1 + "user_alive", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
