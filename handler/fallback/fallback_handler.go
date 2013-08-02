package fallback

import (
	"github.com/bborbe/server/handler_finder"
	"net/http"
)

type fallback struct {
	handlerFinder handler_finder.HandlerFinder
	fallback      http.Handler
}

func NewFallback(handlerFinder handler_finder.HandlerFinder, fallbackHandler http.Handler) *fallback {
	m := new(fallback)
	m.handlerFinder = handlerFinder
	m.fallback = fallbackHandler
	return m
}

func (m *fallback) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	handler := m.handlerFinder.FindHandler(request)
	if handler == nil {
		handler = m.fallback
	}
	handler.ServeHTTP(responseWriter, request)
}
