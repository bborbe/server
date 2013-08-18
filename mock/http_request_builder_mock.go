package mock

import "net/http"

type httpRequestBuilderMock struct {
	url              string
	parameter        map[string][]string
	header           http.Header
	responseProvider ResponseProvider
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

func (r *httpRequestBuilderMock) GetResponse() (*http.Response, error) {
	return r.responseProvider.GetResponse(), r.responseProvider.GetError()
}

func (r *httpRequestBuilderMock) SetResponseBuilder(responseProvider ResponseProvider) {
	r.responseProvider = responseProvider
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
