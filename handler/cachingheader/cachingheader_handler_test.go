package cachingheader

import (
	. "github.com/bborbe/assert"
	"net/http"
	"testing"
)

func TestImplementsHandler(t *testing.T) {
	r := NewCachingHeaderHandler(nil)
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
