package json

import (
	"encoding/json"
	"github.com/bborbe/server/handler/error"
	"net/http"
)

type jsonHandler struct {
	m interface{}
}

func NewJsonHandler(m interface{}) *jsonHandler {
	h := new(jsonHandler)
	h.m = m
	return h
}

func (m *jsonHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	b, err := json.Marshal(m.m)
	if err != nil {
		e := error.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(b)
}
