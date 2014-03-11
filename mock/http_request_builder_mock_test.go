package mock

import (
	"testing"
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/requestbuilder"
)

func TestImplementsHttpRequestBuilder(t *testing.T) {
	r := NewHttpRequestBuilderMock("http://www.example.com")
	var i *requestbuilder.HttpRequestBuilder
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
