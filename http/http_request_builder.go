package http

type HttpRequestBuilder interface {
	AddParameter(key string, value ...string)
}

type httpRequestBuilder struct {
	url       string
	parameter map[string][]string
}

func NewHttpRequestBuilder(url string) *httpRequestBuilder {
	r := new(httpRequestBuilder)
	r.url = url
	r.parameter = make(map[string][]string)
	return r
}

func (r *httpRequestBuilder) AddParameter(key string, values ...string) {
	r.parameter[key] = values
}
