package yaygo

import (
	"encoding/json"
	"net/http"
)

type UserApi struct {
	s *Session
}

func newUserApi(s *Session) *UserApi {
	return &UserApi{
		s: s,
	}
}

func (u *UserApi) GetUser(uID string) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointUsers(uID), nil, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &st)
	return
}
