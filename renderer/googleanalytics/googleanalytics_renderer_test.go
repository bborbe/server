package googleanalytics

import (
	"testing"

	. "github.com/bborbe/assert"
	io "github.com/bborbe/io/mock"
	"github.com/bborbe/server/renderer"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewGoogleanalyticsRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsGoogleanalyticsRenderer(t *testing.T) {
	v := NewGoogleanalyticsRenderer()
	var i (*GoogleanalyticsRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewGoogleanalyticsRenderer()
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
}
