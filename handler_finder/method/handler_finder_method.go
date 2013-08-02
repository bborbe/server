package method

import "net/http"

type handlerFinderMethod struct {
	handler map[string]http.Handler
}

func NewHandlerFinderMethod() *handlerFinderMethod {
	h := new(handlerFinderMethod)
	h.handler = make(map[string]http.Handler)
	return h
}

func (h *handlerFinderMethod) FindHandler(request *http.Request) http.Handler {
	return h.handler[request.Method]
}

func (h *handlerFinderMethod) RegisterHandler(method string, handler http.Handler) {
	h.handler[method] = handler
}
