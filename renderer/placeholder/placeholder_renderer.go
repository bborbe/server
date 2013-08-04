package placeholder

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type PlaceholderRenderer interface {
	renderer.Renderer
	SetRenderer(renderer renderer.Renderer)
}

type placeholderRenderer struct {
	renderer renderer.Renderer
}

func NewPlaceholderRenderer() *placeholderRenderer {
	v := new(placeholderRenderer)
	return v
}

func (v *placeholderRenderer) SetRenderer(renderer renderer.Renderer) {
	v.renderer = renderer
}

func (v *placeholderRenderer) Render(writer io.Writer) error {
	if v.renderer != nil {
		return v.renderer.Render(writer)
	}
	return nil
}
