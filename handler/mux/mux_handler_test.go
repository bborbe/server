package mux

import (
	. "github.com/bborbe/assert"
	"net/http"
	"testing"
)

func TestImplementsHandler(t *testing.T) {
	r := NewMuxHandler(nil, nil)
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
