package failure

import (
	"net/http"
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
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
	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is(`{"status":500,"message":"Internal Server Error"}`))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFailureRendererMessage(t *testing.T) {
	r := NewFailureRendererMessage(http.StatusInternalServerError, "foo bar")
	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is(`{"status":500,"message":"foo bar"}`))
	if err != nil {
		t.Fatal(err)
	}
}
