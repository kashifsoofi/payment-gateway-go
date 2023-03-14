package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err        error  `json:"-"`           // low-level runtime error
	StatusCode int    `json:"status_code"` // http response status code
	Message    string `json:"message,omitempty"`

	ErrorCode int64  `json:"error_code,omitempty"` // application-specific error code
	ErrorText string `json:"error_text,omitempty"` // application-level error message
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf("err: %v, status_code: %d, error_code: %d, message: %s", e.Err, e.StatusCode, e.ErrorCode, e.Message)
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:        err,
		StatusCode: 400,
		Message:    "Invalid request",
		ErrorText:  err.Error(),
	}
}

var ErrNotFound = &ErrResponse{StatusCode: 404, Message: "Resource not found."}

func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:        err,
		StatusCode: 500,
		Message:    "Internal Server Error",
		ErrorText:  err.Error(),
	}
}
