package head

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/html/tag"
)

type HeadRenderer interface {
	renderer.Renderer
}

type headRenderer struct {
	head renderer.Renderer
}

func NewHeadRenderer() *headRenderer {
	v := new(headRenderer)
	v.head = tag.NewTagRenderer("head")
	return v
}

func (v *headRenderer) Render(writer io.Writer) error {
	return v.head.Render(writer)
}
