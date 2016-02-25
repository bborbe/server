package openingtag

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
)

func TestNewOpenRenderer(t *testing.T) {
	var err error
	r := NewOpenRenderer("div")
	writer := bytes.NewBufferString("")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<div>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCloseRenderer(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := bytes.NewBufferString("")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<div/>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddClass(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := bytes.NewBufferString("")
	r.AddClass("content")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<div class=\"content\"/>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddClassSorted(t *testing.T) {
	var err error
	r := NewCloseRenderer("div")
	writer := bytes.NewBufferString("")
	r.AddClass("c")
	r.AddClass("a")
	r.AddClass("b")
	err = r.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Contains("<div class=\"a b c\"/>"))
	if err != nil {
		t.Fatal(err)
	}
}
