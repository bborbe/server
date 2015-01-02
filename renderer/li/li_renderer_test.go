package li

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
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
	writer := io_mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<li></li>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithContent(t *testing.T) {
	r := NewLiRenderer()
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("hello world")
	r.SetContent(contentRenderer)

	writer := io_mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<li>hello world</li>"))
	if err != nil {
		t.Fatal(err)
	}
}
