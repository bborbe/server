package link

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/io"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/content"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewLinkRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
func TestImplementsLinkRenderer(t *testing.T) {
	v := NewLinkRenderer()
	var i (*LinkRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderEmpty(t *testing.T) {
	var err error
	v := NewLinkRenderer()
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<a></a>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewLinkRenderer()
	v.SetHref("/links")
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("foo bar")
	v.SetContent(contentRenderer)
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<a href=\"/links\">foo bar</a>"))
	if err != nil {
		t.Fatal(err)
	}
}
