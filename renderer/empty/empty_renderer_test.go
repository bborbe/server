package empty

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"testing"
)

func TestImplementsRequestHandler(t *testing.T) {
	v := NewEmptyRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i).Message("check implements view.Renderer"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewEmptyRenderer()
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Eq(0))
	if err != nil {
		t.Fatal(err)
	}
}
