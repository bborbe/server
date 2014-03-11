package fallback

import (
	"net/http"

	"github.com/bborbe/log"
	"github.com/bborbe/server/handler_finder"
)

type fallback struct {
	handlerFinder handler_finder.HandlerFinder
	fallback      http.Handler
}

var logger = log.DefaultLogger

func NewFallback(handlerFinder handler_finder.HandlerFinder, fallbackHandler http.Handler) *fallback {
	m := new(fallback)
	m.handlerFinder = handlerFinder
	m.fallback = fallbackHandler
	return m
}

func (m *fallback) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	handler := m.handlerFinder.FindHandler(request)
	if handler == nil {
		logger.Debug("no handler found, use fallback")
		handler = m.fallback
	}
	handler.ServeHTTP(responseWriter, request)
}
