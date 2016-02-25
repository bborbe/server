package failure

import (
	"encoding/json"
	"io"
	"net/http"
)

type message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type failureRenderer struct {
	status  int
	message string
}

func NewFailureRenderer(status int) *failureRenderer {
	return NewFailureRendererMessage(status, http.StatusText(status))
}

func NewFailureRendererMessage(status int, message string) *failureRenderer {
	r := new(failureRenderer)
	r.status = status
	r.message = message
	return r
}

func (r *failureRenderer) Render(writer io.Writer) error {
	m := new(message)
	m.Status = r.status
	m.Message = r.message
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	writer.Write(b)
	return nil
}
