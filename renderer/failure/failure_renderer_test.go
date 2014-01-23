package failure

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"net/http"
	"testing"
)

func TestImplementsRenderer(t *testing.T) {
	r := NewFailureRenderer(http.StatusInternalServerError)
	var expect *renderer.Renderer
	err := AssertThat(r, Implements(expect))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFailureRenderer(t *testing.T) {
	r := NewFailureRenderer(http.StatusInternalServerError)
	writer := mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is(`{"status":500,"message":"Internal Server Error"}`))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFailureRendererMessage(t *testing.T) {
	r := NewFailureRendererMessage(http.StatusInternalServerError, "foo bar")
	writer := mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is(`{"status":500,"message":"foo bar"}`))
	if err != nil {
		t.Fatal(err)
	}
}