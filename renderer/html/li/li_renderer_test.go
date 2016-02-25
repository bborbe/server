package li

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/content"
)

func TestImplementsRenderer(t *testing.T) {
	r := NewLiRenderer()
	var i *renderer.Renderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsUlRenderer(t *testing.T) {
	r := NewLiRenderer()
	var i *LiRenderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	r := NewLiRenderer()
	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<li></li>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithContent(t *testing.T) {
	r := NewLiRenderer()
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("hello world")
	r.SetContent(contentRenderer)

	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<li>hello world</li>"))
	if err != nil {
		t.Fatal(err)
	}
}
