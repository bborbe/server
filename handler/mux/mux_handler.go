package mux

import (
	"github.com/bborbe/server/handler_finder"
	"net/http"
)

type muxHandler struct {
	handlerFinder handler_finder.HandlerFinder
	errorHandler  http.Handler
}

func NewMuxHandler(handlerFinder handler_finder.HandlerFinder, errorHandler http.Handler) *muxHandler {
	m := new(muxHandler)
	m.handlerFinder = handlerFinder
	m.errorHandler = errorHandler
	return m
}

func (m *muxHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	handler := m.handlerFinder.FindHandler(request)
	if handler != nil {
		handler.ServeHTTP(responseWriter, request)
	} else {
		m.errorHandler.ServeHTTP(responseWriter, request)
	}
}
