package rest

import (
	"net/http"

	"github.com/bborbe/server/handler_finder/method"
	"github.com/bborbe/server/handler_finder/part"
)

type RestHandlerFinder interface {
	RegisterListHandler(handler http.Handler)
	RegisterGetHandler(handler http.Handler)
	RegisterCreateHandler(handler http.Handler)
	RegisterDeleteHandler(handler http.Handler)
	RegisterUpdateHandler(handler http.Handler)
	RegisterPatchHandler(handler http.Handler)
}

type restHandlerFinder struct {
	prefix              string
	methodHandlerFinder method.MethodHandlerFinder
	getHandlerFinder    part.PartHandlerFinder
	postHandlerFinder   part.PartHandlerFinder
	putHandlerFinder    part.PartHandlerFinder
	patchHandlerFinder  part.PartHandlerFinder
	deleteHandlerFinder part.PartHandlerFinder
}

func New(prefix string) *restHandlerFinder {
	methodHandlerFinder := method.New()

	getHandlerFinder := part.New(prefix)
	methodHandlerFinder.RegisterHandlerFinder("GET", getHandlerFinder)
	methodHandlerFinder.RegisterHandlerFinder("", getHandlerFinder)

	postHandlerFinder := part.New(prefix)
	methodHandlerFinder.RegisterHandlerFinder("POST", postHandlerFinder)

	putHandlerFinder := part.New(prefix)
	methodHandlerFinder.RegisterHandlerFinder("PUT", putHandlerFinder)

	patchHandlerFinder := part.New(prefix)
	methodHandlerFinder.RegisterHandlerFinder("PATCH", patchHandlerFinder)

	deleteHandlerFinder := part.New(prefix)
	methodHandlerFinder.RegisterHandlerFinder("DELETE", deleteHandlerFinder)

	h := new(restHandlerFinder)
	h.prefix = prefix
	h.methodHandlerFinder = methodHandlerFinder
	h.getHandlerFinder = getHandlerFinder
	h.postHandlerFinder = postHandlerFinder
	h.putHandlerFinder = putHandlerFinder
	h.patchHandlerFinder = patchHandlerFinder
	h.deleteHandlerFinder = deleteHandlerFinder

	return h
}

func (h *restHandlerFinder) FindHandler(request *http.Request) http.Handler {
	return h.methodHandlerFinder.FindHandler(request)
}

func (h *restHandlerFinder) RegisterListHandler(handler http.Handler) {
	h.getHandlerFinder.RegisterHandler("", handler)
}

func (h *restHandlerFinder) RegisterGetHandler(handler http.Handler) {
	h.getHandlerFinder.RegisterHandler("/", handler)
}

func (h *restHandlerFinder) RegisterCreateHandler(handler http.Handler) {
	h.postHandlerFinder.RegisterHandler("", handler)
	h.postHandlerFinder.RegisterHandler("/", handler)
}

func (h *restHandlerFinder) RegisterDeleteHandler(handler http.Handler) {
	h.deleteHandlerFinder.RegisterHandler("", handler)
	h.deleteHandlerFinder.RegisterHandler("/", handler)
}

func (h *restHandlerFinder) RegisterUpdateHandler(handler http.Handler) {
	h.putHandlerFinder.RegisterHandler("", handler)
	h.putHandlerFinder.RegisterHandler("/", handler)
}

func (h *restHandlerFinder) RegisterPatchHandler(handler http.Handler) {
	h.patchHandlerFinder.RegisterHandler("", handler)
	h.patchHandlerFinder.RegisterHandler("/", handler)
}
