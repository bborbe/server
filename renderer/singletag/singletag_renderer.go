package singletag

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/openingtag"
	"io"
)

type SingletagRenderer interface {
	renderer.Renderer
	SetAttribute(key, value string) SingletagRenderer
	RemoveAttribute(key string) SingletagRenderer
}

type singletagRenderer struct {
	name               string
	openingtagRenderer openingtag.OpeningtagRenderer
}

func NewSingletagRenderer(name string) *singletagRenderer {
	v := new(singletagRenderer)
	v.openingtagRenderer = openingtag.NewCloseRenderer(name)
	return v
}

func (v *singletagRenderer) Render(writer io.Writer) error {
	return v.openingtagRenderer.Render(writer)
}

func (v *singletagRenderer) SetAttribute(key, value string) SingletagRenderer {
	v.openingtagRenderer.SetAttribute(key, value)
	return v
}

func (v *singletagRenderer) RemoveAttribute(key string) SingletagRenderer {
	v.openingtagRenderer.RemoveAttribute(key)
	return v
}
