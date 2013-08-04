package tag

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"github.com/bborbe/server/renderer"
	"testing"
)

func TestImplementsRequestHandler(t *testing.T) {
	v := NewTagRenderer("mytag")
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i).Message("check implements view.Renderer"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewTagRenderer("mytag")
	writer := mock.NewWriter()
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<mytag>"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("</mytag>"))
	if err != nil {
		t.Fatal(err)
	}
}
