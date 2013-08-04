package link

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type linkRenderer struct {
	renderer tag.TagRenderer
}

func NewLinkRenderer() *linkRenderer {
	v := new(linkRenderer)
	v.renderer = tag.NewTagRenderer("a")
	return v
}

func (v *linkRenderer) SetHref(href string) {
	v.renderer.SetAttribute("href", href)
}

func (v *linkRenderer) SetContent(content renderer.Renderer) {
	v.renderer.SetContent(content)
}

func (v *linkRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}
