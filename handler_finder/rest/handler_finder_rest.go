package rest

import (
	"net/http"

	"github.com/bborbe/server/handler/fallback"
	"github.com/bborbe/server/handler_finder/method"
	"github.com/bborbe/server/handler_finder/part"
)

type restHandlerFinder struct {
	handler http.Handler
}

func New(prefix string, listHandler http.Handler, getHandler http.Handler, createHandler http.Handler, updateHandler http.Handler, patchHandler http.Handler, deleteHandler http.Handler) *restHandlerFinder {
	h := new(restHandlerFinder)
	methodHandlerFinder := method.New()
	methodHandlerFinder.RegisterHandler("POST", createHandler)
	methodHandlerFinder.RegisterHandler("PUT", updateHandler)
	methodHandlerFinder.RegisterHandler("PATCH", patchHandler)
	methodHandlerFinder.RegisterHandler("DELETE", deleteHandler)
	partHandlerFinder := part.New(prefix)
	partHandlerFinder.RegisterHandler("/", getHandler)
	h.handler = fallback.NewFallback(methodHandlerFinder, fallback.NewFallback(partHandlerFinder, listHandler))
	return h
}

func (h *restHandlerFinder) FindHandler(request *http.Request) http.Handler {
	return h.handler
}
