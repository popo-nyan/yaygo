package yaygo

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	AuthGrantTypePassword     = "password"
	AuthGrantTypeRefreshToken = "refresh_token"
)

type AuthAPI struct {
	s *Session
}

func newAuthAPI(s *Session) *AuthAPI {
	return &AuthAPI{
		s: s,
	}
}

type ChageEmailParams struct {
	Email           string
	Password        string
	EmailGrantToken string
}

func (a *AuthAPI) ChangeEmail(params *ChageEmailParams) (st *LoginUpdateResponse, err error) {
	resp, err := a.s.request(
		http.MethodPut,
		EndpointAuthChangeEmail(),
		nil,
		&struct {
			APIKey          string `json:"api_key,omitempty"`
			Email           string `json:"email,omitempty"`
			Password        string `json:"password,omitempty"`
			EmailGrantToken string `json:"email_grant_token,omitempty"`
		}{
			APIKey:          API_KEY,
			Email:           params.Email,
			Password:        params.Password,
			EmailGrantToken: params.EmailGrantToken,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type ChangePasswordParams struct {
	CurrentPassword string
	Password        string
}

func (a *AuthAPI) ChangePassword(params *ChangePasswordParams) (st *LoginUpdateResponse, err error) {
	resp, err := a.s.request(
		http.MethodPut,
		EndpointAuthChangePassword(),
		nil,
		&struct {
			APIKey          string `json:"api_key,omitempty"`
			CurrentPassword string `json:"current_password,omitempty"`
			Password        string `json:"password,omitempty"`
		}{
			APIKey:          API_KEY,
			CurrentPassword: params.CurrentPassword,
			Password:        params.Password,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetTokenParams struct {
	GrantType    string `json:"grant_type,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (a *AuthAPI) GetToken(params *GetTokenParams) (st *TokenResponse, err error) {
	// formurlencoded
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthGetToken(),
		nil,
		&GetTokenParams{
			GrantType:    params.GrantType,
			Email:        params.Email,
			Password:     params.Password,
			RefreshToken: params.RefreshToken,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type LoginParams struct {
	Email     string
	Password  string
	TwoFACode string
}

func (a *AuthAPI) Login(params *LoginParams) (st *LoginUserResponse, err error) {
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthLoginWithEmail(),
		nil,
		&struct {
			APIKey    string `json:"api_key,omitempty"`
			Email     string `json:"email,omitempty"`
			Password  string `json:"password,omitempty"`
			TwoFACode string `json:"two_fa_code,omitempty"`
			UUID      string `json:"uuid,omitempty"`
		}{
			APIKey:    API_KEY,
			Email:     params.Email,
			Password:  params.Password,
			TwoFACode: params.TwoFACode,
			UUID:      a.s.state.DeviceUUID,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (a *AuthAPI) LogoutDevice() (st *Response, err error) {
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthLogoutDevice(),
		nil,
		&struct {
			UUID string `json:"uuid,omitempty"`
		}{
			UUID: a.s.state.DeviceUUID,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (a *AuthAPI) MigrateToken(token string) (st *TokenResponse, err error) {
	// formurlencoded
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthMigrateToken(),
		nil,
		&struct {
			Token string `json:"token,omitempty"`
		}{
			Token: token,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (a *AuthAPI) ResendComfirmEmail() (st *Response, err error) {
	resp, err := a.s.request(http.MethodPost, EndpointAuthResendConfirmEmail(), nil, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (a *AuthAPI) RestoreUser(uID int) (st *LoginUserResponse, err error) {
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthRestoreUser(),
		nil,
		&struct {
			UserID     int    `json:"user_id,omitempty"`
			APIKey     string `json:"api_key,omitempty"`
			UUID       string `json:"uuid,omitempty"`
			Timestamp  int    `json:"timestamp,omitempty"`
			SignedInfo string `json:"signed_info,omitempty"`
		}{
			UserID:    uID,
			APIKey:    API_KEY,
			UUID:      a.s.state.DeviceUUID,
			Timestamp: int(time.Now().Unix()),
			// TODO: SignedInfo: a.s.state.SignedInfo,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type SaveAccoutWithEmailParams struct {
	Email           string
	Password        string
	CurrentPassword string
	EmailGrantToken string
}

func (a *AuthAPI) SaveAccountWithEmail(params *SaveAccoutWithEmailParams) (st *LoginUpdateResponse, err error) {
	resp, err := a.s.request(
		http.MethodPost,
		EndpointAuthSaveAccountWithEmail(),
		nil,
		&struct {
			APIKey          string `json:"api_key,omitempty"`
			Email           string `json:"email,omitempty"`
			Password        string `json:"password,omitempty"`
			CurrentPassword string `json:"current_password,omitempty"`
			EmailGrantToken string `json:"email_grant_token,omitempty"`
		}{
			APIKey:          API_KEY,
			Email:           params.Email,
			Password:        params.Password,
			CurrentPassword: params.CurrentPassword,
			EmailGrantToken: params.EmailGrantToken,
		},
		false,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
