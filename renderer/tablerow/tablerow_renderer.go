package tablerow

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tablecell"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type TablerowRenderer interface {
	renderer.Renderer
	AddCell(cell tablecell.TablecellRenderer) TablerowRenderer
}

type tablerowRenderer struct {
	renderer tag.TagRenderer
}

func NewTablerowRenderer() *tablerowRenderer {
	v := new(tablerowRenderer)
	v.renderer = tag.NewTagRenderer("tr")
	return v
}

func (v *tablerowRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}

func (v *tablerowRenderer) AddCell(cell tablecell.TablecellRenderer) TablerowRenderer {
	return v
}
