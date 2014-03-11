package log

import (
	"net/http"
	"time"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type logHandler struct {
	handler http.Handler
}

func NewLogHandler(handler http.Handler) *logHandler {
	m := new(logHandler)
	m.handler = handler
	return m
}

func (m *logHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	start := time.Now().Nanosecond()
	logger.Debugf("%s %s", request.Method, request.RequestURI)
	m.handler.ServeHTTP(responseWriter, request)
	end := time.Now().Nanosecond()
	logger.Debugf("%s %s takes %d ms", request.Method, request.RequestURI, (end-start)/1000/1000)
}
