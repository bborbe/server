package ul

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/html/li"
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
	writer := bytes.NewBufferString("")
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<ul></ul>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderWithLi(t *testing.T) {
	r := NewUlRenderer()
	writer := bytes.NewBufferString("")
	li := li.NewLiRenderer()
	r.Add(li)
	err := r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<ul><li></li></ul>"))
	if err != nil {
		t.Fatal(err)
	}
}
