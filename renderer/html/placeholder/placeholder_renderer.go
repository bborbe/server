package placeholder

import (
	"io"

	"github.com/bborbe/server/renderer"
)

type PlaceholderRenderer interface {
	renderer.Renderer
	SetRenderer(renderer renderer.Renderer) PlaceholderRenderer
}

type placeholderRenderer struct {
	renderer renderer.Renderer
}

func NewPlaceholderRenderer() *placeholderRenderer {
	v := new(placeholderRenderer)
	return v
}

func (v *placeholderRenderer) SetRenderer(renderer renderer.Renderer) PlaceholderRenderer {
	v.renderer = renderer
	return v
}

func (v *placeholderRenderer) Render(writer io.Writer) error {
	if v.renderer != nil {
		return v.renderer.Render(writer)
	}
	return nil
}