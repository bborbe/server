package html

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewHtmlRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
func TestImplementsHtmlRenderer(t *testing.T) {
	v := NewHtmlRenderer()
	var i (*HtmlRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewHtmlRenderer()
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Startswith("<!doctype html>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<html>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("</html>"))
	if err != nil {
		t.Fatal(err)
	}
}
