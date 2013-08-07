package content

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type ContentRenderer interface {
	renderer.Renderer
}
type contentRenderer struct {
	content string
}

func NewContentRenderer(content string) *contentRenderer {
	v := new(contentRenderer)
	v.content = content
	return v
}

func (v *contentRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte(v.content))
	if err != nil {
		return err
	}
	return err
}
