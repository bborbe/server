package tag

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/openingtag"
	"io"
)

type TagRenderer interface {
	renderer.Renderer
	SetContent(renderer renderer.Renderer) TagRenderer
	SetAttribute(key, value string) TagRenderer
	RemoveAttribute(key string) TagRenderer
	AddClass(class string) TagRenderer
	RemoveClass(class string) TagRenderer
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

func (v *tagRenderer) SetContent(renderer renderer.Renderer) TagRenderer {
	v.content = renderer
	return v
}

func (v *tagRenderer) SetAttribute(key, value string) TagRenderer {
	v.openingtagRenderer.SetAttribute(key, value)
	return v
}

func (v *tagRenderer) RemoveAttribute(key string) TagRenderer {
	v.openingtagRenderer.RemoveAttribute(key)
	return v
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

func (v *tagRenderer) AddClass(class string) TagRenderer {
	v.openingtagRenderer.AddClass(class)
	return v
}

func (v *tagRenderer) RemoveClass(class string) TagRenderer {
	v.openingtagRenderer.RemoveClass(class)
	return v
}
