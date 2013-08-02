package static

import (
	. "github.com/bborbe/assert"
	"net/http"
	"testing"
)

func TestImplementsRequestHandler(t *testing.T) {
	r := NewHandlerStaticContent("hello")
	var i *http.Handler
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}
