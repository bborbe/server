package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestCreateServer(t *testing.T) {
	server, err := createServer()
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(server, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}
