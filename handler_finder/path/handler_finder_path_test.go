package path

import (
	"github.com/bborbe/server/handler/mux"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
	"github.com/bborbe/server/mock"
	. "github.com/bborbe/assert"
	"net/http"
	"testing"
)

func TestPathImplementsHandlerFinder(t *testing.T) {
	h := NewHandlerFinderPath()
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
	handlerFinder := NewHandlerFinderPath()
	m := mux.NewMuxHandler(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", http.StatusNotFound))
	responseWriter := mock.NewHttpResponseWriterMock()
	request, err := mock.NewHttpRequestMock("http://www.example.com")
	if err != nil {
		t.Error(err)
	}
	m.ServeHTTP(responseWriter, request)
	err = AssertThat(responseWriter.Status, Is(http.StatusNotFound).Message("check status"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandlerFound(t *testing.T) {
	handlerFinder := NewHandlerFinderPath()
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
		err = AssertThat(responseWriter.Status, Is(http.StatusOK).Message("check status"))
		if err != nil {
			t.Fatal(err)
		}
		err = AssertThat(string(responseWriter.Content), Is("/").Message("compare / content"))
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
		err = AssertThat(responseWriter.Status, Is(http.StatusOK).Message("check status"))
		if err != nil {
			t.Fatal(err)
		}
		err = AssertThat(string(responseWriter.Content), Is("/test").Message("compare /test content"))
		if err != nil {
			t.Fatal(err)
		}
	}
}
