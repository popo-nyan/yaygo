package yaygo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (s *Session) request(method, url string, params, body interface{}) (respBody []byte, err error) {
	var req *http.Request

	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		url = url + "?" + v.Encode()
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))

	resp, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
