package base

import (
	"github.com/bborbe/log"
	"github.com/bborbe/server/renderer"
	"net/http"
)

type ViewRendererProvider interface {
	GetViewRenderer(request *http.Request) (renderer.Renderer, error)
}

type FailureRendererProvider interface {
	GetFailureRenderer(err error) renderer.Renderer
}

var logger = log.DefaultLogger

type baseHandler struct {
	viewRendererProvider    ViewRendererProvider
	failureRendererProvider FailureRendererProvider
}

func NewBaseHandler(viewRendererProvider ViewRendererProvider, failureRendererProvider FailureRendererProvider) *baseHandler {
	h := new(baseHandler)
	h.viewRendererProvider = viewRendererProvider
	h.failureRendererProvider = failureRendererProvider
	return h
}

func (m *baseHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	v, err := m.viewRendererProvider.GetViewRenderer(request)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		logger.Errorf("getView failed: %v", err)
		v = m.failureRendererProvider.GetFailureRenderer(err)
	} else {
		responseWriter.WriteHeader(http.StatusOK)
	}
	err = v.Render(responseWriter)
	if err != nil {
		logger.Errorf("View.Render failed: %v", err)
	}
}
