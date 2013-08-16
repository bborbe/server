package ul

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/li"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type UlRenderer interface {
	renderer.Renderer
	Add(li li.LiRenderer)
}

type ulRenderer struct {
	renderer     renderer.Renderer
	listRenderer list.ListRenderer
}

func NewUlRenderer() *ulRenderer {
	r := new(ulRenderer)

	listRenderer := list.NewListRenderer()

	ulRenderer := tag.NewTagRenderer("ul")
	ulRenderer.SetContent(listRenderer)

	r.renderer = ulRenderer
	r.listRenderer = listRenderer
	return r
}

func (r *ulRenderer) Render(writer io.Writer) error {
	return r.renderer.Render(writer)
}

func (r *ulRenderer) Add(li li.LiRenderer) {
	r.listRenderer.Add(li)
}
