package tablecell

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
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
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Startswith("<td>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Endswith("</td>"))
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
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<td>hello world</td>"))
	if err != nil {
		t.Fatal(err)
	}
}
