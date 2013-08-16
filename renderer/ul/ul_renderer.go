package ul

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/li"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type UlRenderer interface {
	renderer.Renderer
	Add(li li.LiRenderer)
	SetAttribute(key, value string)
	RemoveAttribute(key string)
	AddClass(class string)
	RemoveClass(class string)
}

type ulRenderer struct {
	renderer     tag.TagRenderer
	listRenderer list.ListRenderer
}

func NewUlRenderer() *ulRenderer {
	r := new(ulRenderer)

	listRenderer := list.NewListRenderer()

	ulRenderer := tag.NewTagRenderer("ul")
	ulRenderer.SetContent(listRenderer)

	r.renderer = ulRenderer
	r.listRenderer = listRenderer
	return r
}

func (r *ulRenderer) Render(writer io.Writer) error {
	return r.renderer.Render(writer)
}

func (r *ulRenderer) Add(li li.LiRenderer) {
	r.listRenderer.Add(li)
}

func (r *ulRenderer) SetAttribute(key, value string) {
	r.renderer.SetAttribute(key, value)
}

func (r *ulRenderer) RemoveAttribute(key string) {
	r.renderer.RemoveAttribute(key)
}

func (r *ulRenderer) AddClass(class string) {
	r.renderer.AddClass(class)
}

func (r *ulRenderer) RemoveClass(class string) {
	r.renderer.RemoveClass(class)
}
