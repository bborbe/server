package empty

import (
	"io"
)

type emptyRenderer struct {
}

func NewEmptyRenderer() *emptyRenderer {
	v := new(emptyRenderer)
	return v
}

func (v *emptyRenderer) Render(writer io.Writer) error {
	return nil
}
