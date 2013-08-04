package list

import (
	"io"
	"github.com/bborbe/server/renderer"
)

type listRenderer struct {
	list []renderer.Renderer
}

func NewListRenderer(list ... renderer.Renderer) *listRenderer {
	v := new(listRenderer)
	v.list = list
	return v
}

func (v *listRenderer) Render(writer io.Writer) error {
	var err error
	for _, r := range v.list {
		err = r.Render(writer)
		if err != nil {
			return err
		}
	}
	return err
}
