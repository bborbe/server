package input

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
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
	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<input/>"))
	if err != nil {
		t.Fatal(err)
	}
}
