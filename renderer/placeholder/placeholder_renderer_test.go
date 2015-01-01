package placeholder

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/io"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewPlaceholderRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsPlaceholderRenderer(t *testing.T) {
	v := NewPlaceholderRenderer()
	var i (*PlaceholderRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithoutContent(t *testing.T) {
	var err error
	v := NewPlaceholderRenderer()
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Eq(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithContent(t *testing.T) {
	var err error
	v := NewPlaceholderRenderer()
	v.SetRenderer(tag.NewTagRenderer("h1"))
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<h1></h1>"))
	if err != nil {
		t.Fatal(err)
	}
}
