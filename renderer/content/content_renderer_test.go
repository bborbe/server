package content

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"testing"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewContentRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsContentRenderer(t *testing.T) {
	v := NewContentRenderer()
	var i (*ContentRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewContentRenderer()
	v.SetContent("mycontent")
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("mycontent"))
	if err != nil {
		t.Fatal(err)
	}
}
