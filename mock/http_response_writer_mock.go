package mock

import "net/http"

type responseWriterMock struct {
	status int
	writer *writer
	header http.Header
}

func NewHttpResponseWriterMock() *responseWriterMock {
	r := new(responseWriterMock)
	r.header = make(http.Header)
	r.writer = NewWriter()
	return r
}

func (r *responseWriterMock) Header() http.Header {
	return r.header
}

func (r *responseWriterMock) Write(b []byte) (int, error) {
	return r.writer.Write(b)
}

func (r *responseWriterMock) WriteHeader(status int) {
	r.status = status
}

func (r *responseWriterMock) Status() int {
	return r.status
}

func (r *responseWriterMock) Content() []byte {
	return r.writer.Content()
}
