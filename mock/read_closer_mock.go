package mock

import (
	"io"
	"strings"
)

type readCloser struct {
	reader io.Reader
}

func NewReadCloserString(content string) *readCloser {
	r := new(readCloser)
	r.reader = strings.NewReader(content)
	return r
}

func (r *readCloser) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r *readCloser) Close() error {
	return nil
}
