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
