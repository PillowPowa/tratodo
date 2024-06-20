package api

import (
	"fmt"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

func NewApiError(statusCode int, message string) Error {
	statusText := http.StatusText(statusCode)

	if statusText == "" {
		return NewApiError(http.StatusInternalServerError, message)
	}

	return Error{
		StatusCode: statusCode,
		StatusText: statusText,
		Message:    message,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s]: %s", e.StatusText, e.Message)
}
