package mock

import "net/http"

type httpRequestBuilderMock struct {
	url              string
	parameter        map[string][]string
	header           http.Header
	requestProvider RequestProvider
}

func NewHttpRequestBuilderMock(url string) *httpRequestBuilderMock {
	r := new(httpRequestBuilderMock)
	r.url = url
	r.parameter = make(map[string][]string)
	r.header = make(http.Header)
	return r
}

func (r *httpRequestBuilderMock) AddHeader(key string, values ...string) {
	r.header[key] = values
}

func (r *httpRequestBuilderMock) AddParameter(key string, values ...string) {
	r.parameter[key] = values
}

func (r *httpRequestBuilderMock) GetRequest() (*http.Request, error) {
	return r.requestProvider.GetRequest(), r.requestProvider.GetError()
}

func (r *httpRequestBuilderMock) SetRequestBuilder(requestProvider RequestProvider) {
	r.requestProvider = requestProvider
}

func (r *httpRequestBuilderMock) GetUrl() string {
	return r.url
}

func (r *httpRequestBuilderMock) GetParameter() map[string][]string {
	return r.parameter
}

func (r *httpRequestBuilderMock) GetHeader() http.Header {
	return r.header
}
