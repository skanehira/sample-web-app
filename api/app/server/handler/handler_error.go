package handler

import "errors"

var (
	ErrMissingUserID = errors.New("missing user id")
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}
