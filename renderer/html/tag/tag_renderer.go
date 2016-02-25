package tag

import (
	"io"

	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/html/openingtag"
)

type TagRenderer interface {
	renderer.Renderer
	SetContent(renderer renderer.Renderer)
	SetAttribute(key, value string)
	RemoveAttribute(key string)
	AddClass(class string)
	RemoveClass(class string)
}

type tagRenderer struct {
	name               string
	content            renderer.Renderer
	openingtagRenderer openingtag.OpeningtagRenderer
}

func NewTagRenderer(name string) *tagRenderer {
	v := new(tagRenderer)
	v.name = name
	v.openingtagRenderer = openingtag.NewOpenRenderer(name)
	return v
}

func (v *tagRenderer) SetContent(renderer renderer.Renderer) {
	v.content = renderer
}

func (v *tagRenderer) SetAttribute(key, value string) {
	v.openingtagRenderer.SetAttribute(key, value)
}

func (v *tagRenderer) RemoveAttribute(key string) {
	v.openingtagRenderer.RemoveAttribute(key)
}

func (v *tagRenderer) Render(writer io.Writer) error {
	var err error
	err = v.openingtagRenderer.Render(writer)
	if err != nil {
		return err
	}
	if v.content != nil {
		err = v.content.Render(writer)
		if err != nil {
			return err
		}
	}
	_, err = writer.Write([]byte("</"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(v.name))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(">"))
	if err != nil {
		return err
	}
	return err
}

func (v *tagRenderer) AddClass(class string) {
	v.openingtagRenderer.AddClass(class)
}

func (v *tagRenderer) RemoveClass(class string) {
	v.openingtagRenderer.RemoveClass(class)
}
