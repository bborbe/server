package json

import (
	"encoding/json"
	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/error"
	"net/http"
)

var logger = log.DefaultLogger

type jsonHandler struct {
	m interface{}
}

func NewJsonHandler(m interface{}) *jsonHandler {
	h := new(jsonHandler)
	h.m = m
	return h
}

func (m *jsonHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Debug("write json")
	b, err := json.Marshal(m.m)
	if err != nil {
		logger.Debugf("Marshal json failed: %v", err)
		e := error.NewErrorMessage(http.StatusInternalServerError, err.Error())
		e.ServeHTTP(responseWriter, request)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(b)
}
