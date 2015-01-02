package mock

import (
	"net/http"
	"net/url"

	io_mock "github.com/bborbe/io/mock"
)

func NewHttpRequestMock(urlString string) (*http.Request, error) {
	r := new(http.Request)
	r.Header = make(http.Header)
	r.Body = io_mock.NewReadCloserString("")
	var err error
	r.URL, err = url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	r.RequestURI = r.URL.RequestURI()
	return r, nil
}
