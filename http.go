package yaygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

func (s *Session) buildRequest(method, url string, params, body interface{}, jwt bool) (*http.Request, error) {
	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		url = url + "?" + v.Encode()
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	if s.state.IP == "" && !strings.Contains(url, EndpointUsersTimestamp()) {
		resp, err := s.User.GetTimestamp()
		if err != nil {
			return nil, err
		}
		s.state.IP = resp.IPAddress
	}

	ua := "android 11 (3.5x 1440x2960 Galaxy S9)"
	// We are not using req.Header.Set() to avoid header key conversion.
	req.Header = http.Header{
		"Host":               []string{EndpointHost},
		"User-Agent":         []string{ua},
		"X-Timestamp":        []string{strconv.FormatInt(time.Now().Unix(), 10)},
		"X-App-Version":      []string{API_VERSION_NAME},
		"X-Device-Info":      []string{fmt.Sprintf("yay %s %s", VERSION_NAME, ua)},
		"X-Device-UUID":      []string{s.state.DeviceUUID},
		"X-Client-IP":        []string{s.state.IP},
		"X-Connection-Type":  []string{"wifi"},
		"X-Connection-Speed": []string{"0kbps"},
		"Accept-Language":    []string{"ja"},
		"Content-Type":       []string{"application/json;charset=UTF-8"},
	}

	if jwt {
		req.Header["X-Jwt"] = []string{"TODO: generateJWT()"}
	}

	if s.state.AccessToken != "" {
		req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", s.state.AccessToken)}
	}

	return req, nil
}

func (s *Session) baseRequest(method, url string, params, body interface{}, jwt bool) ([]byte, error) {
	req, err := s.buildRequest(method, url, params, body, jwt)
	if err != nil {
		return nil, err
	}

	slog.Debug(fmt.Sprintf("[REQ] %s :: %s", method, req.URL))
	for k, v := range req.Header {
		slog.Debug(fmt.Sprintf("[REQ] HEADERS %s: %s", k, v))
	}
	slog.Debug(fmt.Sprintf("[REQ] %s", req.Body))
	slog.Debug("[FET] ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range resp.Header {
		slog.Debug(fmt.Sprintf("[RES] HEADERS %s: %s", k, v))
	}
	slog.Debug(fmt.Sprintf("[RES] %s :: %s\n\n\n", resp.Status, byteResp))

	var errResp ErrorResponse
	if err := json.Unmarshal(byteResp, &errResp); err == nil && errResp.Result == "error" {
		return nil, mapError(&errResp)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%w: %s", ErrStatusBadRequest, byteResp)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%w: %s", ErrStatusUnauthorized, byteResp)
	case http.StatusForbidden:
		return nil, fmt.Errorf("%w: %s", ErrStatusForbidden, byteResp)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%w: %s", ErrStatusNotFound, byteResp)
	default:
		if resp.StatusCode >= 500 {
			return nil, fmt.Errorf("%w: %s", ErrStatusInternalServerError, byteResp)
		} else {
			return nil, fmt.Errorf("http error %d: %s", resp.StatusCode, byteResp)
		}
	}

	return byteResp, nil
}

func (s *Session) request(method, url string, params, body interface{}, jwt bool) ([]byte, error) {
	return s.baseRequest(method, url, params, body, jwt)
}
