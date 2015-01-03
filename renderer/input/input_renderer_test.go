package input

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
	"github.com/bborbe/server/renderer"
)

func TestImplementsRenderer(t *testing.T) {
	r := NewInputRenderer()
	var i *renderer.Renderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsInputRenderer(t *testing.T) {
	r := NewInputRenderer()
	var i *InputRenderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	r := NewInputRenderer()
	writer := io_mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<input/>"))
	if err != nil {
		t.Fatal(err)
	}
}
