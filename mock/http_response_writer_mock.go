package mock

import "net/http"

type responseWriterMock struct {
	Status  int
	Content []byte
	header  http.Header
}

func NewHttpResponseWriterMock() *responseWriterMock {
	r := new(responseWriterMock)
	r.header = make(http.Header)
	return r
}

func (r *responseWriterMock) Header() http.Header {
	return r.header
}

func (r *responseWriterMock) Write(b []byte) (int, error) {
	r.Content = append(b)
	return len(b), nil
}

func (r *responseWriterMock) WriteHeader(status int) {
	r.Status = status
}
