package tag

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewTagRenderer("mytag")
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
func TestImplementsTagRenderer(t *testing.T) {
	v := NewTagRenderer("mytag")
	var i (*TagRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewTagRenderer("mytag")
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<mytag>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("</mytag>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderAttributes(t *testing.T) {
	var err error
	v := NewTagRenderer("mytag")
	v.SetAttribute("a", "b")
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<mytag a=\"b\"></mytag>"))
	if err != nil {
		t.Fatal(err)
	}
}