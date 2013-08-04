package content

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"testing"
)

func TestImplementsRequestHandler(t *testing.T) {
	v := NewContentRenderer("mycontent")
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i).Message("check implements view.Renderer"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewContentRenderer("mycontent")
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
