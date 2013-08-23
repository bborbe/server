package view

import (
	"net/http"
)

type View interface {
	Render(responseWriter http.ResponseWriter) error
}
