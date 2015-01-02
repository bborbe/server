package empty

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
	"github.com/bborbe/server/renderer"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewEmptyRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
func TestImplementsEmptyRenderer(t *testing.T) {
	v := NewEmptyRenderer()
	var i (*EmptyRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewEmptyRenderer()
	writer := io_mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Eq(0))
	if err != nil {
		t.Fatal(err)
	}
}
