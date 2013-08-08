package content

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type ContentRenderer interface {
	renderer.Renderer
	SetContent(content string)
}
type contentRenderer struct {
	content string
}

func NewContentRenderer() *contentRenderer {
	v := new(contentRenderer)
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

func (v *contentRenderer) SetContent(content string) {
	v.content = content
}
