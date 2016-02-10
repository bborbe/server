package list

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/singletag"
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
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Eq(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderOne(t *testing.T) {
	var err error
	v := NewListRenderer(singletag.NewSingletagRenderer("br"))
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<br/>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderTwo(t *testing.T) {
	var err error
	v := NewListRenderer(singletag.NewSingletagRenderer("br"), singletag.NewSingletagRenderer("hr"))
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<br/><hr/>"))
	if err != nil {
		t.Fatal(err)
	}
}
