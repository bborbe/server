package list

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type ListRenderer interface {
	renderer.Renderer
	Add(renderer renderer.Renderer) ListRenderer
}

type listRenderer struct {
	list []renderer.Renderer
}

func NewListRenderer(list ...renderer.Renderer) *listRenderer {
	v := new(listRenderer)
	v.list = list
	return v
}

func (v *listRenderer) Add(renderer renderer.Renderer) ListRenderer {
	v.list = append(v.list, renderer)
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
