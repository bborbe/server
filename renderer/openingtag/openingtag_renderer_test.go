package openingtag

import (
	"testing"
	. "github.com/bborbe/assert"
	"github.com/bborbe/io"
)

func TestNewOpenRenderer(t *testing.T) {
	var err error
	r := NewOpenRenderer("div")
	writer := io.NewWriter()
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
	writer := io.NewWriter()
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
	writer := io.NewWriter()
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

func TestAddClassSorted(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := io.NewWriter()
	r.AddClass("c")
	r.AddClass("a")
	r.AddClass("b")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.Content()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Contains("<div class=\"a b c\"/>"))
	if err != nil {
		t.Fatal(err)
	}
}
