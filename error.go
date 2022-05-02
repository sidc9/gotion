package gotion

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *ErrResponse) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func (e *ErrResponse) Unwrap() error {
	return e.Err
}

func NewErrResponse(code, message string, err error) *ErrResponse {
	return &ErrResponse{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
