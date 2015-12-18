package error

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
	server_mock "github.com/bborbe/server/mock"
)

func TestImplementsRequestHandler(t *testing.T) {
	r := NewError(http.StatusNotFound)
	var i (*http.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements http.Handler"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestContent(t *testing.T) {
	handler := NewError(http.StatusNotFound)
	responseWriter := server_mock.NewHttpResponseWriterMock()
	request, err := server_mock.NewHttpRequestMock("http://www.example.com/foobar")
	if err != nil {
		t.Error(err)
	}
	handler.ServeHTTP(responseWriter, request)
	{
		err := AssertThat(responseWriter.Status(), Is(http.StatusNotFound).Message("check status"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		err := AssertThat(string(responseWriter.Content()), Is("{\"status\":404,\"message\":\"Not Found\"}").Message("check content"))
		if err != nil {
			t.Fatal(err)
		}
	}
}
