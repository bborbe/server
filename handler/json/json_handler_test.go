package json

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsHandler(t *testing.T) {
	r := NewJsonHandler(nil)
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
