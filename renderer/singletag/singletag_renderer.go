package singletag

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/openingtag"
	"io"
)

type SingletagRenderer interface {
	renderer.Renderer
	renderer.Attribute
	renderer.Class
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

func (v *singletagRenderer) SetAttribute(key, value string) {
	v.openingtagRenderer.SetAttribute(key, value)
}

func (v *singletagRenderer) RemoveAttribute(key string) {
	v.openingtagRenderer.RemoveAttribute(key)
}

func (v *singletagRenderer) AddClass(class string) {
	v.openingtagRenderer.AddClass(class)
}

func (v *singletagRenderer) RemoveClass(class string) {
	v.openingtagRenderer.RemoveClass(class)
}
