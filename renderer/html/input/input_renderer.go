package input

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/html/singletag"
)

type InputRenderer interface {
	renderer.Renderer
	SetAttribute(key, value string)
	RemoveAttribute(key string)
	AddClass(class string)
	RemoveClass(class string)
}

type inputRenderer struct {
	renderer singletag.SingletagRenderer
}

func NewInputRenderer() *inputRenderer {
	r := new(inputRenderer)
	inputRenderer := singletag.NewSingletagRenderer("input")
	r.renderer = inputRenderer
	return r
}

func (r *inputRenderer) Render(writer io.Writer) error {
	return r.renderer.Render(writer)
}

func (r *inputRenderer) SetAttribute(key, value string) {
	r.renderer.SetAttribute(key, value)
}

func (r *inputRenderer) RemoveAttribute(key string) {
	r.renderer.RemoveAttribute(key)
}

func (r *inputRenderer) AddClass(class string) {
	r.renderer.AddClass(class)
}

func (r *inputRenderer) RemoveClass(class string) {
	r.renderer.RemoveClass(class)
}
