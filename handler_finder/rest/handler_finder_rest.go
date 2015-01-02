package rest

import (
	"net/http"

	"github.com/bborbe/server/handler/fallback"
	"github.com/bborbe/server/handler_finder/method"
	"github.com/bborbe/server/handler_finder/part"
	"github.com/bborbe/server/handler_finder"
)

type restHandlerFinder struct {
	handlerFinder handler_finder.HandlerFinder
}

func New(prefix string, listHandler http.Handler, getHandler http.Handler, createHandler http.Handler, updateHandler http.Handler, patchHandler http.Handler, deleteHandler http.Handler) *restHandlerFinder {

	partHandlerFinder := part.New(prefix)
	partHandlerFinder.RegisterHandler("/", getHandler)

	methodHandlerFinder := method.New()
	methodHandlerFinder.RegisterHandler("GET", fallback.NewFallback(partHandlerFinder, listHandler))
	methodHandlerFinder.RegisterHandler("POST", createHandler)
	methodHandlerFinder.RegisterHandler("PUT", updateHandler)
	methodHandlerFinder.RegisterHandler("PATCH", patchHandler)
	methodHandlerFinder.RegisterHandler("DELETE", deleteHandler)


	h := new(restHandlerFinder)
	h.handlerFinder = methodHandlerFinder
	return h
}

func (h *restHandlerFinder) FindHandler(request *http.Request) http.Handler {
	return h.handlerFinder.FindHandler(request)
}
