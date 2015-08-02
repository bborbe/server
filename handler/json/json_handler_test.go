package json

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/mock"
)

func TestImplementsHandler(t *testing.T) {
	r := NewJsonHandler(nil)
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func TestRenderStruct(t *testing.T) {
	u := &user{
		FirstName: "Hello",
		LastName:  "World",
	}
	r := NewJsonHandler(u)
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal()
	}
	r.ServeHTTP(resp, req)
	if err = AssertThat(string(resp.Content()), Is(`{"firstname":"Hello","lastname":"World"}`)); err != nil {
		t.Fatal(err)
	}
}

func TestRenderListOfStruct(t *testing.T) {
	list := []user{user{
		FirstName: "Hello",
		LastName:  "World",
	}, user{
		FirstName: "Hello",
		LastName:  "World",
	}}
	r := NewJsonHandler(list)
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal()
	}
	r.ServeHTTP(resp, req)
	if err = AssertThat(string(resp.Content()), Is(`[{"firstname":"Hello","lastname":"World"},{"firstname":"Hello","lastname":"World"}]`)); err != nil {
		t.Fatal(err)
	}
}

func TestRenderEmptyList(t *testing.T) {
	list := make([]user, 0)
	r := NewJsonHandler(list)
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal()
	}
	r.ServeHTTP(resp, req)
	if err = AssertThat(string(resp.Content()), Is(`[]`)); err != nil {
		t.Fatal(err)
	}
}
func TestRenderNilList(t *testing.T) {
	var list []user
	r := NewJsonHandler(list)
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal()
	}
	r.ServeHTTP(resp, req)
	if err = AssertThat(string(resp.Content()), Is(`[]`)); err != nil {
		t.Fatal(err)
	}
}
