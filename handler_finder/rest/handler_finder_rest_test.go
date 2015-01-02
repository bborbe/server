package rest

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
	"github.com/bborbe/server/mock"
)

func TestImplementsHandlerFinder(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	h := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	var handler *handler_finder.HandlerFinder
	err := AssertThat(h, Implements(handler).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"GET", RequestURI: "/test/1"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("get"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestList(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"GET", RequestURI: "/test"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("list"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"POST", RequestURI: "/test"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("create"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"PUT", RequestURI: "/test"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("update"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPatch(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"PATCH", RequestURI: "/test"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("patch"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	hf := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{Method:"DELETE", RequestURI: "/test"}
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(string(resp.Content()), Is("delete"))
	if err != nil {
		t.Fatal(err)
	}
}
