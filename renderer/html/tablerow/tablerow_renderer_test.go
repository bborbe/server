package tablerow

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
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
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Startswith("<tr>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Endswith("</tr>"))
	if err != nil {
		t.Fatal(err)
	}
}