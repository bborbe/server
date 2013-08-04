package body

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type bodyRenderer struct {
	body renderer.Renderer
}

func NewBodyRenderer() *bodyRenderer {
	v := new(bodyRenderer)
	v.body = tag.NewTagRenderer("body")
	return v
}

func (v *bodyRenderer) Render(writer io.Writer) error {
	return v.body.Render(writer)
}
