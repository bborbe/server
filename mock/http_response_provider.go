package mock

import (
	"net/http"

	io_mock "github.com/bborbe/io/mock"
)

type ResponseProvider interface {
	GetResponse() *http.Response
	GetError() error
}

type responseProvider struct {
	content string
	err     error
	status  int
}

func NewResponseProvider(status int, content string, err error) *responseProvider {
	p := new(responseProvider)
	p.status = status
	p.content = content
	p.err = err
	return p
}

func (p *responseProvider) GetResponse() *http.Response {
	response := new(http.Response)
	response.StatusCode = p.status
	response.Body = io_mock.NewReadCloserString(p.content)
	return response
}

func (p *responseProvider) GetError() error {
	return p.err
}
