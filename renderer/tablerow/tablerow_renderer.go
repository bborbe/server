package tablerow

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type TablerowRenderer interface {
	renderer.Renderer
	AddCell(renderer renderer.Renderer)
}

type tablerowRenderer struct {
	renderer tag.TagRenderer
	cells    list.ListRenderer
}

func NewTablerowRenderer() *tablerowRenderer {
	v := new(tablerowRenderer)
	tr := tag.NewTagRenderer("tr")
	cells := list.NewListRenderer()
	tr.SetContent(cells)
	v.cells = cells
	v.renderer = tr
	return v
}

func (v *tablerowRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}

func (v *tablerowRenderer) AddCell(renderer renderer.Renderer) {
	v.cells.Add(renderer)
}
