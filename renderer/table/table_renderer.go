package table

import (
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type tableRenderer struct {
	renderer tag.TagRenderer
}

func NewTableRenderer() *tableRenderer {
	v := new(tableRenderer)
	v.renderer = tag.NewTagRenderer("table")
	return v
}

func (v *tableRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}
