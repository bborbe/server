package idparser

import (
	"net/http"
	"strconv"
	"regexp"
	"fmt"
)

func ParseIdFormRequest(request *http.Request) (int, error) {
	return ParseIdFromUri(request.RequestURI)
}

func ParseIdFromUri(uri string) (int, error) {
	re := regexp.MustCompile("(\\d+)[^/]*?$")
	matches := re.FindStringSubmatch(uri)
	if len(matches) > 1 {
		return strconv.Atoi(matches[1])
	}
	return 0, fmt.Errorf("parse id from uri %s failed", uri)
}
