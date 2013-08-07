package openingtag

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
	"testing"
)

func TestNewOpenRenderer(t *testing.T) {
	var err error
	r := NewOpenRenderer("div")
	writer := mock.NewWriter()
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<div>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCloseRenderer(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := mock.NewWriter()
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<div/>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddClass(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := mock.NewWriter()
	r.AddClass("content")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<div class=\"content\"/>"))
	if err != nil {
		t.Fatal(err)
	}
}
