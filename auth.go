package yaygo

import (
	"encoding/json"
	"net/http"
)

type AuthAPI struct {
	s *Session
}

func newAuthAPI(s *Session) *AuthAPI {
	return &AuthAPI{
		s: s,
	}
}

type LoginParams struct {
	Email     string
	Password  string
	TwoFACode string
}

type LoginEmailUserRequest struct {
	APIKey    string `json:"api_key,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	TwoFACode string `json:"two_fa_code,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

func (a *AuthAPI) Login(params *LoginParams) (st *LoginUserResponse, err error) {
	resp, err := a.s.request(
		http.MethodPost,
		EndpointUsersLoginWithEmail(),
		nil,
		&LoginEmailUserRequest{
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
