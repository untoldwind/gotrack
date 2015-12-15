package http

import (
	"fmt"
	"github.com/untoldwind/gotrack/server/logging"
	"github.com/untoldwind/routing"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error %d: %s", err.Code, err.Message)
}

func MethodNotAllowed() *HTTPError {
	return &HTTPError{Code: 405, Message: "Method not allowed"}
}

func NotFound() *HTTPError {
	return &HTTPError{Code: 404, Message: "Not found"}
}

func BadRequest() *HTTPError {
	return &HTTPError{Code: 400, Message: "Bad request"}
}

func Unauthorized() *HTTPError {
	return &HTTPError{Code: 401, Message: "Unauthorized"}
}

func SendError(logger logging.Logger, err *HTTPError) routing.Matcher {
	return func(remainingPath string, resp http.ResponseWriter, req *http.Request) bool {
		encodeError(logger, resp, req, err)
		return true
	}
}
