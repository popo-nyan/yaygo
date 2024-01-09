package yaygo

var (
	EndpointYayAPI = "https://api.yay.space/"
	EndpointV1     = EndpointYayAPI + "v1/"
	EndpointV2     = EndpointYayAPI + "v2/"
	EndpointV3     = EndpointYayAPI + "v3/"

	EndpointUsersV1 = EndpointV1 + "users/"
	EndpointUsersV2 = EndpointV1 + "users/"
	EndpointUsersV3 = EndpointV1 + "users/"

	EndpointGroupsV1 = EndpointV1 + "groups/"
	EndpointGroupsV2 = EndpointV2 + "groups/"
	EndpointGroupsV3 = EndpointV3 + "groups/"

	EndpointReviewsV1 = EndpointUsersV1 + "reviews/"
	EndpointReviewsV2 = EndpointUsersV2 + "reviews/"

	EndpointCallsV1 = EndpointV1 + "calls/"
	EndpointCallsV2 = EndpointV2 + "calls/"

	EndpointPostsV1 = EndpointV1 + "posts/"
	EndpointPostsV2 = EndpointV2 + "posts/"
	EndpointPostsV3 = EndpointV3 + "posts/"

	EndpointChatRoomsV1 = EndpointV1 + "chat_rooms/"
	EndpointChatRoomsV2 = EndpointV2 + "chat_rooms/"
	EndpointChatRoomsV3 = EndpointV3 + "chat_rooms/"

	EndpointConversationsV2 = EndpointV2 + "conversations/"
	EndpointPinnedV1        = EndpointV1 + "pinned/"
	EndpointHiddenV1        = EndpointV1 + "hidden/"
	EndpoinWalletV1         = EndpointV1 + "wallet/"
	EndpintThreadsV1        = EndpointUsersV1 + "threads/"

	EndpointUsersRegister                                = func() string { return EndpointUsersV3 + "register" }
	EndpointUsersContactFriends                          = func() string { return EndpointUsersV1 + "contact_friends" }
	EndpointUsersDeleteFootprint                         = func(uID string, fuID string) string { return EndpointUsersV2 + uID + "/footprints/" + fuID }
	EndpointUsersDestroyUser                             = func() string { return EndpointUsersV2 + "destroy" }
	EndpointUsersFollowUser                              = func(uID string) string { return EndpointUsersV2 + uID + "/follow" }
	EndpointUsersFollow                                  = func() string { return EndpointUsersV2 + "follow" }
	EndpointUsersActiveFollowings                        = func() string { return EndpointUsersV1 + "active_followings" }
	EndpointUsersAdditionalSettings                      = func() string { return EndpointUsersV1 + "additonal_notification_setting" }
	EndpointUsersdAppReviewStatus                        = func(uuid string) string { return EndpointUsersV1 + uuid + "app_review_status" }
	EndpointUsersContactStatus                           = func() string { return EndpointUsersV1 + "contact_status" }
	EndpointUsersDefaultSettings                         = func() string { return EndpointUsersV1 + "default_settings" }
	EndpointFriends                                      = func() string { return EndpointV1 + "friends" }
	EndpointUsersFollowRequests                          = func() string { return EndpointUsersV2 + "follow_requests" }
	EndpointUsersFollowRequestsCount                     = func() string { return EndpointUsersV2 + "follow_requests_count" }
	EndpointUsersFollowingBornToday                      = func() string { return EndpointUsersV2 + "following_born_today" }
	EndpointUsersFootprints                              = func() string { return EndpointUsersV2 + "footprints" }
	EndpointUsersFresh                                   = func(uID string) string { return EndpointUsersV2 + "fresh/" + uID }
	EndpointUsersHimaUsers                               = func() string { return EndpointUsersV2 + "hima_users" }
	EndpointUsersFollowRecommended                       = func(uID string) string { return EndpointUsersV1 + uID + "/follow_recommended" }
	EndpointUsersResetCounters                           = func() string { return EndpointUsersV1 + "reset_counters" }
	EndpointUsersTimestamp                               = func() string { return EndpointUsersV2 + "timestamp" }
	EndpointUsers                                        = func(uID string) string { return EndpointUsersV2 + uID }
	EndpointUsersCustomDefinitions                       = func() string { return EndpointUsersV1 + "custom_definitions" }
	EndpointUsersFollowers                               = func(uID string) string { return EndpointUsersV2 + uID + "/followers" }
	EndpointUsersListFollowings                          = func(uID string) string { return EndpointUsersV2 + uID + "/list_followings" }
	EndpointUsersQrCodes                                 = func(qr string) string { return EndpointUsersV1 + "qr_codes/" + qr }
	EndpointUsersInterests                               = func() string { return EndpointUsersV1 + "interests" }
	EndpointUsersInfo                                    = func(uID string) string { return EndpointUsersV2 + "info/" + uID }
	EndpointUsersListId                                  = func() string { return EndpointUsersV1 + "list_id" }
	EndpointUsersRemoveUserAvatar                        = func() string { return EndpointUsersV2 + "remove_profile_photo" }
	EndpointUsersRemoveUserCover                         = func() string { return EndpointUsersV2 + "remove_cover_image" }
	EndpointUsersResetPassword                           = func() string { return EndpointUsersV1 + "reset_password" }
	EndpointUsersLobiFriends                             = func() string { return EndpointV1 + "lobi_friends" }
	EndpointUsersSearch                                  = func() string { return EndpointUsersV1 + "search" }
	EndpointUsersAdditonalMotificationSetting            = func() string { return EndpointUsersV1 + "additonal_notification_setting" }
	EndpointUsersEnditV2                                 = func() string { return EndpointUsersV2 + "edit" }
	EndpointUsersVisibleOnSnsFriendRecommendationSetting = func() string { return EndpointUsersV1 + "visible_on_sns_friend_recommendation_setting" }
	EndpointUsersFollowRequest                           = func(tID string) string { return EndpointUsersV2 + tID + "/follow_request" }
	EndpointUsersHima                                    = func() string { return EndpointUsersV1 + "hima" }
	EndpointUsersUnFollow                                = func(uID string) string { return EndpointUsersV2 + uID + "/unfollow" }
	EndpointUsersInviteContact                           = func() string { return EndpointUsersV1 + "invite_contact" }
	EndpointUsersLanguage                                = func() string { return EndpointUsersV1 + "language" }
	EndpointUsersEditV3                                  = func() string { return EndpointUsersV3 + "edit" }
	EndpointUsersBookmarks                               = func(uID string, bID string) string { return EndpointUsersV1 + uID + "/bookmarks/" + bID }
	EndpointUsersGetBookmark                             = func(uID string) string { return EndpointUsersV1 + uID + "/bookmarks" }

	EndpointPostsNewConferenceCall      = func() string { return EndpointPostsV2 + "new_conference_call" }
	EndpointPostsGroupPinnedPost        = func() string { return EndpointPostsV2 + "group_pinned_post" }
	EndpointPostsNew                    = func() string { return EndpointPostsV3 + "new" }
	EndpointPostsRepost                 = func() string { return EndpointPostsV3 + "repost" }
	EndpointPostsNewSharePost           = func() string { return EndpointPostsV2 + "new_share_post" }
	EndpointPostsDeleteAllPost          = func() string { return EndpointPostsV1 + "delete_all_post" }
	EndpointPostsCallTimeline           = func() string { return EndpointPostsV2 + "call_timeline" }
	EndpointGetConversation             = func(cID string) string { return EndpointConversationsV2 + cID }
	EndpointConversationRootPosts       = func() string { return EndpointConversationsV2 + "root_posts" }
	EndpointPostsCallFollowersTimeline  = func() string { return EndpointPostsV2 + "call_followers_timeline" }
	EndpointPostsFollowingTimeline      = func() string { return EndpointPostsV2 + "following_timeline" }
	EndpointPostsGroupTimeline          = func() string { return EndpointPostsV2 + "group_timeline" }
	EndpointPostsTag                    = func(tag string) string { return EndpointPostsV2 + "tags/" + tag }
	EndpointPostsMin                    = func() string { return EndpointPostsV2 + "mine" }
	EndpointPostsGetPost                = func(pID string) string { return EndpointPostsV2 + pID }
	EndpointPostsLikers                 = func(pID string) string { return EndpointPostsV1 + pID + "/likers" }
	EndpointPostsRePosts                = func(pID string) string { return EndpointPostsV2 + pID + "/reposts" }
	EndpointPostsMultiple               = func() string { return EndpointPostsV2 + "multiple" }
	EndpointPostsRecentEngagement       = func() string { return EndpointPostsV2 + "recent_engagement" }
	EndpointPostsRecommendedTag         = func() string { return EndpointPostsV1 + "recommended_tag" }
	EndpointPostsRecommendedTimeline    = func() string { return EndpointPostsV2 + "recommended_timeline" }
	EndpointPostsSearch                 = func() string { return EndpointPostsV2 + "search" }
	EndpointPostsNoreplyModeTimeline    = func(nMODE string) string { return EndpointPostsV2 + nMODE + "/timeline" }
	EndpointPostsUrlMetadata            = func() string { return EndpointPostsV2 + "url_metadata" }
	EndpointPostsUserTimeline           = func() string { return EndpointPostsV2 + "user_timeline" }
	EndpointPostsLike                   = func() string { return EndpointPostsV2 + "like" }
	EndpointPostsMassDestroy            = func() string { return EndpointPostsV2 + "mass_destroy" }
	EndpointPostsUnLike                 = func(pID string) string { return EndpointPostsV2 + pID + "/unlike" }
	EndpointPostsUpdate                 = func(pID string) string { return EndpointPostsV3 + pID }
	EndpointPostsRecommendationFeedback = func(pID string) string { return EndpointPostsV2 + pID + "/recommendation_feedback" }
	EndpointPostsValidate               = func() string { return EndpointPostsV1 + "validate" }
	EndpointPostsVideoView              = func(vID string) string { return EndpointPostsV1 + "/videos" + vID + "/view" }
	EndpointVoteSurvey                  = func(sID string) string { return EndpointV2 + "surveys" + sID + "/vote" }
)
