package tag

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type TagRenderer interface {
renderer.Renderer
	SetAttribute(key, value string)
	RemoveAttribute(key string)
	SetContent(renderer renderer.Renderer)
}

type tagRenderer struct {
	name    string
	content renderer.Renderer
	attributes map[string]string
}

func NewTagRenderer(name string) *tagRenderer {
	v := new(tagRenderer)
	v.name = name
	v.attributes = make(map[string]string)
	return v
}

func (v *tagRenderer) SetContent(renderer renderer.Renderer) {
	v.content = renderer
}

func (v *tagRenderer) SetAttribute(key, value string) {
	v.attributes[key] = value
}

func (v *tagRenderer) RemoveAttribute(key string) {
	delete(v.attributes, key)
}

func (v *tagRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(v.name))
	if err != nil {
		return err
	}
	for k, v := range v.attributes {
		_, err = writer.Write([]byte(" "))
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte(k))
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte("=\""))
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte(v))
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = writer.Write([]byte(">"))
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
