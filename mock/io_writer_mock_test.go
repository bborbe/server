package mock

import (
	"io"
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplementsIoWriter(t *testing.T) {
	writer := NewWriter()
	var expected *io.Writer
	err := AssertThat(writer, Implements(expected).Message("check type"))
	if err != nil {
		t.Error(err)
	}
}

func TestIoWriter(t *testing.T) {
	var err error
	writer := NewWriter()
	err = AssertThat(string(writer.Content()), Is(""))
	if err != nil {
		t.Error(err)
	}
	writer.Write([]byte("hello"))
	err = AssertThat(string(writer.Content()), Is("hello"))
	if err != nil {
		t.Error(err)
	}
	writer.Write([]byte(" world"))
	err = AssertThat(string(writer.Content()), Is("hello world"))
	if err != nil {
		t.Error(err)
	}
}
