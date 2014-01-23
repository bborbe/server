package error

import (
	"github.com/bborbe/log"
	"github.com/bborbe/server/renderer/failure"
	"net/http"
)

var logger = log.DefaultLogger

type object struct {
	status  int
	message string
}

func NewError(status int) *object {
	return NewErrorMessage(status, http.StatusText(status))
}

func NewErrorMessage(status int, message string) *object {
	o := new(object)
	o.status = status
	o.message = message
	return o
}

func (o *object) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Debug("handle error")
	r := failure.NewFailureRendererMessage(o.status, o.message)
	logger.Debugf("set status: %d", o.status)
	responseWriter.WriteHeader(o.status)
	responseWriter.Header().Set("Content-Type", "application/json")
	err := r.Render(responseWriter)
	if err != nil {
		logger.Warnf("render failureRenderer failed! %v", err)
	}
}
