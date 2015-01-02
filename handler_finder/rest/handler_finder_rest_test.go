package rest

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
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

func TestList(t *testing.T) {
	listHandler := static.NewHandlerStaticContent("list")
	getHandler := static.NewHandlerStaticContent("get")
	createHandler := static.NewHandlerStaticContent("create")
	updateHandler := static.NewHandlerStaticContent("update")
	patchHandler := static.NewHandlerStaticContent("patch")
	deleteHandler := static.NewHandlerStaticContent("delete")
	h := New("/test", listHandler, getHandler, createHandler, updateHandler, patchHandler, deleteHandler)
	r := &http.Request{}
	h.FindHandler(r)
}
