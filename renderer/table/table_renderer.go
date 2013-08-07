package table

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/tablerow"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type TableRenderer interface {
	renderer.Renderer
	AddRow(row tablerow.TablerowRenderer) TableRenderer
}

type tableRenderer struct {
	renderer tag.TagRenderer
	rows     list.ListRenderer
}

func NewTableRenderer() *tableRenderer {
	v := new(tableRenderer)
	table := tag.NewTagRenderer("table")
	rows := list.NewListRenderer()
	table.SetContent(rows)
	v.rows = rows
	v.renderer = table
	return v
}

func (v *tableRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}

func (v *tableRenderer) AddRow(tablerow tablerow.TablerowRenderer) TableRenderer {
	v.rows.Add(tablerow)
	return v
}
