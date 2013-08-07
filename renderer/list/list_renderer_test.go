package list

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/singletag"
	"testing"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewListRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
func TestImplementsListRenderer(t *testing.T) {
	v := NewListRenderer()
	var i (*ListRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderEmpty(t *testing.T) {
	var err error
	v := NewListRenderer()
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

func TestRenderOne(t *testing.T) {
	var err error
	v := NewListRenderer(singletag.NewSingletagRenderer("br"))
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<br/>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderTwo(t *testing.T) {
	var err error
	v := NewListRenderer(singletag.NewSingletagRenderer("br"), singletag.NewSingletagRenderer("hr"))
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("<br/><hr/>"))
	if err != nil {
		t.Fatal(err)
	}
}
