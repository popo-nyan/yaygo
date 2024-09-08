package yaygo

import (
	"encoding/json"
	"net/http"
)

type MiscAPI struct {
	s *Session
}

func newMiscAPI(s *Session) *MiscAPI {
	return &MiscAPI{
		s: s,
	}
}

type AcceptPolicyAgreementParams struct {
	Type string `json:"type,omitempty"`
}

func (m *MiscAPI) AcceptPolicyAgreement(params *AcceptPolicyAgreementParams) (st *Response, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1+"accept_policy_agreement", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GenerateSnsThumbnailParams struct {
	ResourceType string `json:"resource_type,omitempty"`
	ResourceID   int    `json:"resource_id,omitempty"`
}

func (m *MiscAPI) GenerateSnsThumbnail(params *GenerateSnsThumbnailParams) (st *Response, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"generate_sns_thumbnail", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetEmailVerificationPresignedURLParams struct {
	DeviceUUID string `json:"device_uuid,omitempty"`
	Email      string `json:"email,omitempty"`
	Locale     string `json:"locale,omitempty"`
	Intent     string `json:"intent,omitempty"`
}

func (m *MiscAPI) GetEmailVerificationPresignedURL(params *GetEmailVerificationPresignedURLParams) (st *EmailVerificationPresignedURLResponse, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1+"get_email_verification_presigned_url", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetFileUploadPresignedURLsParams struct {
	FileNames []string `json:"file_names[],omitempty"`
}

func (m *MiscAPI) GetFileUploadPresignedURLs(params *GetFileUploadPresignedURLsParams) (st *PresignedURLsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_file_upload_presigned_urls", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetIDCheckerPresignedURLParams struct {
	Model  string            `json:"model,omitempty"`
	Action string            `json:"action,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}

func (m *MiscAPI) GetIDCheckerPresignedURL(params *GetIDCheckerPresignedURLParams) (st *IDCheckerPresignedURLResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_id_checker_presigned_url", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetOldFileUploadPresignedURLParams struct {
	VideoFileName string `json:"video_file_name,omitempty"`
}

func (m *MiscAPI) GetOldFileUploadPresignedURL(params *GetOldFileUploadPresignedURLParams) (st *PresignedURLResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_old_file_upload_presigned_url", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPolicyAgreementsParams struct{}

func (m *MiscAPI) GetPolicyAgreements(params *GetPolicyAgreementsParams) (st *PolicyAgreementsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_policy_agreements", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetPromotionsParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"number,omitempty"`
}

func (m *MiscAPI) GetPromotions(params *GetPromotionsParams) (st *PromotionsResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_promotions", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetVipGameRewardURLParams struct {
	DeviceType string `json:"device_type,omitempty"`
}

func (m *MiscAPI) GetVipGameRewardURL(params *GetVipGameRewardURLParams) (st *VipGameRewardURLResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_vip_game_reward_url", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type GetWebSocketTokenParams struct{}

func (m *MiscAPI) GetWebSocketToken(params *GetWebSocketTokenParams) (st *WebSocketTokenResponse, err error) {
	resp, err := m.s.request(http.MethodGet, EndpointChatRoomsV1+"get_websocket_token", params, nil, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}

type VerifyDeviceParams struct {
	AppVersion         string `json:"app_version,omitempty"`
	Platform           string `json:"platform,omitempty"`
	DeviceUUID         string `json:"device_uuid,omitempty"`
	VerificationString string `json:"verification_string,omitempty"`
}

func (m *MiscAPI) VerifyDevice(params *VerifyDeviceParams) (st *VerifyDeviceResponse, err error) {
	resp, err := m.s.request(http.MethodPost, EndpointChatRoomsV1+"verify_device", nil, params, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
