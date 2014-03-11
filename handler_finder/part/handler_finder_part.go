package part

import (
	"net/http"
	"strings"

	"github.com/bborbe/log"
)

type handlerFinderPart struct {
	prefix  string
	handler map[string]http.Handler
}

var logger = log.DefaultLogger

func NewHandlerFinderPart(prefix string) *handlerFinderPart {
	h := new(handlerFinderPart)
	h.handler = make(map[string]http.Handler)
	h.prefix = prefix
	return h
}

func (h *handlerFinderPart) RegisterHandler(part string, handler http.Handler) {
	h.handler[part] = handler
}

func (h *handlerFinderPart) FindHandler(request *http.Request) http.Handler {
	rest := request.RequestURI[len(h.prefix):]
	logger.Tracef("requestUri: %s prefix: %s => rest: %s", request.RequestURI, h.prefix, rest)
	if len(rest) == 0 {
		return h.handler[rest]
	}
	if rest[:1] == "/" {
		pos := findEndPos(rest[1:])
		var name string
		if pos != -1 {
			name = rest[:pos+1]
		} else {
			name = rest
		}
		handler := h.handler[name]
		if handler != nil {
			return handler
		}
		return h.handler["/"]
	}
	pos := findEndPos(rest)
	var name string
	if pos != -1 {
		name = rest[:pos]
	} else {
		name = rest
	}
	return h.handler[name]
}

func findEndPos(content string) int {
	return strings.IndexFunc(content, endRunes)
}

func endRunes(r rune) bool {
	return r == '/' || r == '?'
}
