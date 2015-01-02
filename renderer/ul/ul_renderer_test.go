package ul

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/li"
)

func TestImplementsRenderer(t *testing.T) {
	r := NewUlRenderer()
	var i *renderer.Renderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsUlRenderer(t *testing.T) {
	r := NewUlRenderer()
	var i *UlRenderer
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	r := NewUlRenderer()
	writer := io_mock.NewWriter()
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<ul></ul>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithLi(t *testing.T) {
	r := NewUlRenderer()
	writer := io_mock.NewWriter()
	li := li.NewLiRenderer()
	r.Add(li)
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<ul><li></li></ul>"))
	if err != nil {
		t.Fatal(err)
	}
}
