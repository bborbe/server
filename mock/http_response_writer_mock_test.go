package mock

import (
	. "github.com/bborbe/assert"
	"net/http"
	"testing"
)

func TestNewHttpResponseWriterMock(t *testing.T) {
	o := NewHttpResponseWriterMock()
	var expected *http.ResponseWriter
	err := AssertThat(o, Implements(expected).Message("check type"))
	if err != nil {
		t.Error(err)
	}
}
