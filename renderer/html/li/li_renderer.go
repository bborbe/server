package li

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/html/tag"
)

type LiRenderer interface {
	renderer.Renderer
	SetContent(renderer renderer.Renderer) LiRenderer
}

type liRenderer struct {
	renderer tag.TagRenderer
}

func NewLiRenderer() *liRenderer {
	r := new(liRenderer)
	r.renderer = tag.NewTagRenderer("li")
	return r
}

func (r *liRenderer) Render(writer io.Writer) error {
	return r.renderer.Render(writer)
}

func (r *liRenderer) SetContent(renderer renderer.Renderer) LiRenderer {
	r.renderer.SetContent(renderer)
	return r
}
