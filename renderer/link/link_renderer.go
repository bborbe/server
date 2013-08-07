package link

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type LinkRenderer interface {
	renderer.Renderer
	SetHref(href string) LinkRenderer
	SetContent(content renderer.Renderer) LinkRenderer
}

type linkRenderer struct {
	renderer tag.TagRenderer
}

func NewLinkRenderer() *linkRenderer {
	v := new(linkRenderer)
	v.renderer = tag.NewTagRenderer("a")
	return v
}

func (v *linkRenderer) SetHref(href string) LinkRenderer {
	v.renderer.SetAttribute("href", href)
	return v
}

func (v *linkRenderer) SetContent(content renderer.Renderer) LinkRenderer {
	v.renderer.SetContent(content)
	return v
}

func (v *linkRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}
