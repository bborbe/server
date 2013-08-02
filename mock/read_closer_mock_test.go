package mock

import (
	. "github.com/bborbe/assert"
	"io"
	"io/ioutil"
	"testing"
)

func TestNewReadCloserString(t *testing.T) {
	o := NewReadCloserString("foo")
	var expected *io.ReadCloser
	err := AssertThat(o, Implements(expected).Message("check type"))
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	input := "foo"
	o := NewReadCloserString(input)
	output, err := ioutil.ReadAll(o)
	if err != nil {
		t.Error(err)
	}
	err = AssertThat(string(output), Is(input).Message("check input"))
	if err != nil {
		t.Error(err)
	}
}

func TestClose(t *testing.T) {
	o := NewReadCloserString("foo")
	err := AssertThat(o.Close(), NilValue().Message("check close"))
	if err != nil {
		t.Error(err)
	}
}
