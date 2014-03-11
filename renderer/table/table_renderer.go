package table

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/tablerow"
	"github.com/bborbe/server/renderer/tag"
)

type TableRenderer interface {
	renderer.Renderer
	renderer.Attribute
	AddRow(row tablerow.TablerowRenderer)
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

func (v *tableRenderer) AddRow(tablerow tablerow.TablerowRenderer) {
	v.rows.Add(tablerow)
}

func (v *tableRenderer) SetAttribute(key, value string) {
	v.renderer.SetAttribute(key, value)
}

func (v *tableRenderer) RemoveAttribute(key string) {
	v.renderer.RemoveAttribute(key)
}
