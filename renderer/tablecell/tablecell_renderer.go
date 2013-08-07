package tablecell

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type TablecellRenderer interface {
	renderer.Renderer
	SetContent(content renderer.Renderer) TablecellRenderer
	AddClass(class string) TablecellRenderer
}

type tablecellRenderer struct {
	renderer tag.TagRenderer
	classes  map[string]bool
}

func NewTablecellRenderer() *tablecellRenderer {
	v := new(tablecellRenderer)
	v.renderer = tag.NewTagRenderer("td")
	v.classes = make(map[string]bool)
	return v
}

func (v *tablecellRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}

func (v *tablecellRenderer) SetContent(content renderer.Renderer) TablecellRenderer {
	v.renderer.SetContent(content)
	return v
}

func (v *tablecellRenderer) AddClass(class string) TablecellRenderer {
	v.renderer.AddClass(class)
	return v
}
