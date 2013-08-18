package mock

import "net/http"

type ResponseProvider interface {
	GetResponse() *http.Response
	GetError() error
}

type responseProvider struct {
	content string
	err     error
}

func NewResponseProvider(content string, err error) *responseProvider {
	p := new(responseProvider)
	p.content = content
	p.err = err
	return p
}

func (p *responseProvider) GetResponse() *http.Response {
	response := new(http.Response)
	response.Body = NewReadCloserString(p.content)
	return response
}

func (p *responseProvider) GetError() error {
	return p.err
}
