package tablecell

import (
	"testing"

	. "github.com/bborbe/assert"
	io "github.com/bborbe/io/mock"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/content"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewTablecellRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsTablecellRenderer(t *testing.T) {
	v := NewTablecellRenderer()
	var i (*TablecellRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewTablecellRenderer()
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Startswith("<td>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Endswith("</td>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetContent(t *testing.T) {
	var err error
	v := NewTablecellRenderer()
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("hello world")
	v.SetContent(contentRenderer)
	writer := io.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<td>hello world</td>"))
	if err != nil {
		t.Fatal(err)
	}
}
