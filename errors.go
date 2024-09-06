package yaygo

import (
	"errors"
	"fmt"
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
