package idparser

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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
