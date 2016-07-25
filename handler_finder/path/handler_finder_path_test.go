package path

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http/mock"
	"github.com/bborbe/server/handler/mux"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
)

func TestPathImplementsHandlerFinder(t *testing.T) {
	h := New()
	var handlerFinder *handler_finder.HandlerFinder
	err := AssertThat(h, Implements(handlerFinder).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewMuxHandler(t *testing.T) {
	var handlerFinder handler_finder.HandlerFinder
	m := mux.NewMuxHandler(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", http.StatusNotFound))
	var expect *http.Handler
	err := AssertThat(m, Implements(expect).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNotHandlerFound(t *testing.T) {
	handlerFinder := New()
	m := mux.NewMuxHandler(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", http.StatusNotFound))
	responseWriter := mock.NewHttpResponseWriterMock()
	request, err := mock.NewHttpRequestMock("http://www.example.com")
	if err != nil {
		t.Error(err)
	}
	m.ServeHTTP(responseWriter, request)
	err = AssertThat(responseWriter.Status(), Is(http.StatusNotFound).Message("check status"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandlerFound(t *testing.T) {
	handlerFinder := New()
	handlerFinder.RegisterHandler("/", static.NewHandlerStaticContent("/"))
	handlerFinder.RegisterHandler("/test", static.NewHandlerStaticContent("/test"))
	m := mux.NewMuxHandler(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", http.StatusNotFound))
	{
		responseWriter := mock.NewHttpResponseWriterMock()
		request, err := mock.NewHttpRequestMock("http://www.example.com/")
		if err != nil {
			t.Error(err)
		}
		m.ServeHTTP(responseWriter, request)
		err = AssertThat(responseWriter.Status(), Is(http.StatusOK).Message("check status"))
		if err != nil {
			t.Fatal(err)
		}
		err = AssertThat(responseWriter.String(), Is("/").Message("compare / content"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		responseWriter := mock.NewHttpResponseWriterMock()
		request, err := mock.NewHttpRequestMock("http://www.example.com/test")
		if err != nil {
			t.Error(err)
		}
		m.ServeHTTP(responseWriter, request)
		err = AssertThat(responseWriter.Status(), Is(http.StatusOK).Message("check status"))
		if err != nil {
			t.Fatal(err)
		}
		err = AssertThat(responseWriter.String(), Is("/test").Message("compare /test content"))
		if err != nil {
			t.Fatal(err)
		}
	}
}
