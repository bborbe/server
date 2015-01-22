package idparser

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestParseIdFromUriInvalid(t *testing.T) {
	var err error
	_, err = ParseIdFromUri("")
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseIdFromUriValid(t *testing.T) {
	id, err := ParseIdFromUri("/foo/bar/1")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(id, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseIdFromUriValidWithQuestionMark(t *testing.T) {
	id, err := ParseIdFromUri("/foo/bar/1?asdf=asdf")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(id, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseIdFromUriValidWithHash(t *testing.T) {
	id, err := ParseIdFromUri("/foo/bar/1?#top")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(id, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseIdFromUriValidWithNumbers(t *testing.T) {
	id, err := ParseIdFromUri("/4/bar/1?#2=3")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(id, Is(1))
	if err != nil {
		t.Fatal(err)
	}
}