package tablerow

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"testing"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewTablerowRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsTablerowRenderer(t *testing.T) {
	v := NewTablerowRenderer()
	var i (*TablerowRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewTablerowRenderer()
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Startswith("<tr>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Endswith("</tr>"))
	if err != nil {
		t.Fatal(err)
	}
}
