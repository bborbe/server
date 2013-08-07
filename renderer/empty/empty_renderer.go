package empty

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type EmptyRenderer interface {
	renderer.Renderer
}

type emptyRenderer struct {
}

func NewEmptyRenderer() *emptyRenderer {
	v := new(emptyRenderer)
	return v
}

func (v *emptyRenderer) Render(writer io.Writer) error {
	return nil
}
