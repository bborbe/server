package debug

import (
	"net/http"
	"time"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type logHandler struct {
	handler http.Handler
}

func New(handler http.Handler) *logHandler {
	m := new(logHandler)
	m.handler = handler
	return m
}

func (m *logHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	start := time.Now()
	defer logger.Debugf("%s %s takes %dms", request.Method, request.RequestURI, time.Now().Sub(start) / time.Millisecond)

	logger.Debugf("request %v: ", request)
	m.handler.ServeHTTP(responseWriter, request)
	logger.Debugf("response %v: ", responseWriter)
}
