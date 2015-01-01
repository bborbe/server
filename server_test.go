package server

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsServer(t *testing.T) {
	addr := ":12345"
	handler := new(helloWorld)
	srv := NewServer(addr, handler)
	var server *Server
	err := AssertThat(srv, Implements(server).Message("should be Server"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStop(t *testing.T) {
	var err error
	port := 12345
	addr := "0.0.0.0:" + strconv.Itoa(port)
	handler := new(helloWorld)
	srv := NewServer(addr, handler)
	err = AssertThat(srv, NotNilValue().Message("create server failed"))
	if err != nil {
		t.Fatal(err)
	}

	err = AssertThat(srv.Start(), NilValue().Message("not started"))
	if err != nil {
		t.Fatal(err)
	}

	err = AssertThat(srv.Start(), NotNilValue().Message("already started"))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://" + addr)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = AssertThat(string(b), Is("hello, world!\n").Message("content false"))
	if err != nil {
		t.Fatal(err)
	}

	err = AssertThat(srv.Stop(), NilValue().Message("not stopped"))
	if err != nil {
		t.Fatal(err)
	}

	err = AssertThat(srv.Stop(), NotNilValue().Message("already stopped"))
	if err != nil {
		t.Fatal(err)
	}
}

type helloWorld struct {
}

func (h *helloWorld) GetPath() string {
	return "/"
}

func (h *helloWorld) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello, world!\n"))
}
