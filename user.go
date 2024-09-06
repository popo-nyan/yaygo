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

func (u *UserApi) GetTimestamp() (st *UserTimestampResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointUsersTimestamp(), nil, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

func (u *UserApi) GetUser(uID int) (st *UserResponse, err error) {
	resp, err := u.s.request(http.MethodGet, EndpointUsers(uID), nil, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
