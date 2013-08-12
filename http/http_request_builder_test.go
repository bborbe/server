package http

import (
	. "github.com/bborbe/assert"
	"testing"
)

func TestImplementsHttpRequestBuilder(t *testing.T) {
	r := NewHttpRequestBuilder("http://www.example.com")
	var i *HttpRequestBuilder
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
