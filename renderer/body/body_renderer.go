package body

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type BodyRenderer interface {
	renderer.Renderer
	SetContent(renderer renderer.Renderer) BodyRenderer
}

type bodyRenderer struct {
	renderer tag.TagRenderer
}

func NewBodyRenderer() *bodyRenderer {
	v := new(bodyRenderer)
	v.renderer = tag.NewTagRenderer("body")
	return v
}

func (v *bodyRenderer) SetContent(renderer renderer.Renderer) BodyRenderer {
	v.renderer.SetContent(renderer)
	return v
}

func (v *bodyRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}
