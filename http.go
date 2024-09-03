package yaygo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/go-querystring/query"
)

var (
	ErrUnknown                    = errors.New("ErrUnknown")
	ErrInvalidParameter           = errors.New("ErrInvalidParameter")
	ErrRegisteredUser             = errors.New("ErrRegisteredUser")
	ErrAccessTokenExpired         = errors.New("ErrAccessTokenExpired")
	ErrScreenNameAlreadyBeenTaken = errors.New("ErrScreenNameAlreadyBeenTaken")
	ErrUserNotFound               = errors.New("ErrUserNotFound")

	ErrStatusBadRequest          = errors.New("ErrStatusBadRequest")
	ErrStatusUnauthorized        = errors.New("ErrStatusUnauthorized")
	ErrStatusForbidden           = errors.New("ErrStatusForbidden")
	ErrStatusNotFound            = errors.New("ErrStatusNotFound")
	ErrStatusInternalServerError = errors.New("ErrStatusInternalServerError")
)

func mapError(errResp *ErrorResponse) error {
	switch errResp.ErrorCode {
	case 0:
		return fmt.Errorf("%w: %s", ErrUnknown, errResp.Message)
	case -1:
		return fmt.Errorf("%w: %s", ErrInvalidParameter, errResp.Message)
	case -2:
		return fmt.Errorf("%w: %s", ErrRegisteredUser, errResp.Message)
	case -3:
		return fmt.Errorf("%w: %s", ErrAccessTokenExpired, errResp.Message)
	case -4:
		return fmt.Errorf("%w: %s", ErrScreenNameAlreadyBeenTaken, errResp.Message)
	case -5:
		return fmt.Errorf("%w: %s", ErrUserNotFound, errResp.Message)
	default:
		return fmt.Errorf("unknown error code %d: %s", errResp.ErrorCode, errResp.Message)
	}
}

func (s *Session) baseRequest(method, url string, params, body interface{}) ([]byte, error) {
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

	slog.Debug(fmt.Sprintf("[REQ] %s :: %s", method, url))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	slog.Debug(fmt.Sprintf("[RES] %s :: %s", resp.Status, byteResp))

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

func (s *Session) request(method, url string, params, body interface{}) ([]byte, error) {
	return s.baseRequest(method, url, params, body)
}
