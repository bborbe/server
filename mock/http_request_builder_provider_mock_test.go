package mock

import (
	"testing"
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/requestbuilder"
)

func TestImplementsNewHttpRequestBuilderProvider(t *testing.T) {
	p := NewHttpRequestBuilderProviderMock()
	var i *requestbuilder.HttpRequestBuilderProvider
	err := AssertThat(p, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewHttpRequestBuilder(t *testing.T) {
	var err error
	p := NewHttpRequestBuilderProviderMock()
	{
		rb := p.NewHttpRequestBuilder("http://example.com")
		err = AssertThat(rb, NilValue())
		if err != nil {
			t.Fatal(err)
		}
	}
	p.Register("http://example.com", NewHttpRequestBuilderMock("http://example.com"))
	{
		rb := p.NewHttpRequestBuilder("http://example.com")
		err = AssertThat(rb, NotNilValue())
		if err != nil {
			t.Fatal(err)
		}
		var i *requestbuilder.HttpRequestBuilder
		err = AssertThat(rb, Implements(i))
		if err != nil {
			t.Fatal(err)
		}
	}
}
