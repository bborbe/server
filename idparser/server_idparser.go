package idparser

import (
	"net/http"
	"strconv"
	"strings"
)

func ParseIdFormRequest(request *http.Request) (int, error) {
	return ParseIdFromUri(request.RequestURI)
}

func ParseIdFromUri(uri string) (int, error) {
	pos := strings.LastIndex(uri, "/")
	return strconv.Atoi(uri[pos+1:])
}
