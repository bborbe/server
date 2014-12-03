package mock

import (
	"github.com/bborbe/log"
	"github.com/bborbe/http/requestbuilder"
)

type httpRequestBuilderProviderMock struct {
	builder map[string]requestbuilder.HttpRequestBuilder
}

var logger = log.DefaultLogger

func NewHttpRequestBuilderProviderMock() *httpRequestBuilderProviderMock {
	p := new(httpRequestBuilderProviderMock)
	p.builder = make(map[string]requestbuilder.HttpRequestBuilder)
	return p
}

func (p *httpRequestBuilderProviderMock) NewHttpRequestBuilder(url string) requestbuilder.HttpRequestBuilder {
	logger.Debugf("httpRequestBuilderProviderMock.NewHttpRequestBuilder url: %s", url)
	return p.builder[url]
}

func (p *httpRequestBuilderProviderMock) Register(url string, requestbuilder requestbuilder.HttpRequestBuilder) {
	logger.Debugf("httpRequestBuilderProviderMock.Register url: %s rb: %v", url, requestbuilder)
	p.builder[url] = requestbuilder
}
