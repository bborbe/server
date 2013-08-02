package method

import (
	. "github.com/bborbe/assert"
	"github.com/bborbe/server/handler_finder"
	"testing"
)

func TestImplementsHandlerFinder(t *testing.T) {
	h := NewHandlerFinderMethod()
	var handler *handler_finder.HandlerFinder
	err := AssertThat(h, Implements(handler).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
