package static

import (
	"net/http"
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplementsHandler(t *testing.T) {
	r := NewHandlerStaticContent("hello")
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
