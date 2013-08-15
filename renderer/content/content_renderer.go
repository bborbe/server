package content

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type ContentRenderer interface {
	renderer.Renderer
	SetContent(content []byte)
	SetContentString(content string)
}

type contentRenderer struct {
	content []byte
}

func NewContentRenderer() *contentRenderer {
	v := new(contentRenderer)
	return v
}

func (v *contentRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write(v.content)
	if err != nil {
		return err
	}
	return err
}

func (v *contentRenderer) SetContentString(content string) {
	v.SetContent([]byte(content))
}

func (v *contentRenderer) SetContent(content []byte) {
	v.content = content
}
