package yaygo

import (
	"encoding/json"
	"net/http"
)

type MiscApi struct {
	s *Session
}

func newMiscApi(s *Session) *MiscApi {
	return &MiscApi{
		s: s,
	}
}




type AcceptPolicyAgreementParams struct {
	Type string `json:"type"`
}

func (m *MiscAPI) AcceptPolicyAgreement(params *AcceptPolicyAgreementParams) (st *Response, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1 + "accept_policy_agreement/" + params.Type, nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GenerateSnsThumbnailParams struct {
	ResourceType string `json:"resource_type"`
	ResourceID   int64  `json:"resource_id"`
}

func (m *MiscAPI) GenerateSnsThumbnail(params *GenerateSnsThumbnailParams) (st *Response, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "generate_sns_thumbnail", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetEmailVerificationPresignedUrlParams struct {
	UUID   string `json:"device_uuid"`
	Email  string `json:"email"`
	Locale string `json:"locale"`
	Intent *string `json:"intent,omitempty"`
}

func (m *MiscAPI) GetEmailVerificationPresignedUrl(params *GetEmailVerificationPresignedUrlParams) (st *EmailVerificationPresignedUrlResponse, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1 + "get_email_verification_presigned_url", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFileUploadPresignedUrlsParams struct {
	FileNames []string `json:"file_names"`
}

func (m *MiscAPI) GetFileUploadPresignedUrls(params *GetFileUploadPresignedUrlsParams) (st *PresignedUrlsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_file_upload_presigned_urls", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetIdCheckerPresignedUrlParams struct {
	Model  string            `json:"model"`
	Action string            `json:"action"`
	Params map[string]string `json:"params"`
}

func (m *MiscAPI) GetIdCheckerPresignedUrl(params *GetIdCheckerPresignedUrlParams) (st *IdCheckerPresignedUrlResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_id_checker_presigned_url/" + params.Model + "/" + params.Action, params.Params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetOldFileUploadPresignedUrlParams struct {
	VideoFileName string `json:"video_file_name"`
}

func (m *MiscAPI) GetOldFileUploadPresignedUrl(params *GetOldFileUploadPresignedUrlParams) (st *PresignedUrlResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_old_file_upload_presigned_url", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (m *MiscAPI) GetPolicyAgreements() (st *PolicyAgreementsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_policy_agreements", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPromotionsParams struct {
	Page  int `json:"page"`
	Limit *int `json:"number,omitempty"`
}

func (m *MiscAPI) GetPromotions(params *GetPromotionsParams) (st *PromotionsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_promotions", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetVipGameRewardUrlParams struct {
	DeviceType string `json:"device_type"`
}

func (m *MiscAPI) GetVipGameRewardUrl(params *GetVipGameRewardUrlParams) (st *VipGameRewardUrlResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_vip_game_reward_url", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

func (m *MiscAPI) GetWebSocketToken() (st *WebSocketTokenResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_websocket_token", nil, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type VerifyDeviceParams struct {
	AppVersion        string `json:"app_version"`
	Platform          string `json:"platform"`
	UUID              string `json:"device_uuid"`
	VerificationString string `json:"verification_string"`
}

func (m *MiscAPI) VerifyDevice(params *VerifyDeviceParams) (st *VerifyDeviceResponse, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1 + "verify_device", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}