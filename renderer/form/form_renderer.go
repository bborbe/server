package form

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
)

type FormRenderer interface {
	renderer.Renderer
	renderer.Attribute
	SetTarget(target string) FormRenderer
	SetContent(content renderer.Renderer) FormRenderer
}

type formRenderer struct {
	renderer tag.TagRenderer
}

func NewFormRenderer() *formRenderer {
	v := new(formRenderer)
	v.renderer = tag.NewTagRenderer("form")
	return v
}

func (v *formRenderer) SetTarget(target string) FormRenderer {
	v.renderer.SetAttribute("target", target)
	return v
}

func (v *formRenderer) SetContent(content renderer.Renderer) FormRenderer {
	v.renderer.SetContent(content)
	return v
}

func (v *formRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}

func (v *formRenderer) SetAttribute(key, value string) {
	v.renderer.SetAttribute(key, value)
}

func (v *formRenderer) RemoveAttribute(key string) {
	v.renderer.RemoveAttribute(key)
}
