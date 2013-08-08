package openingtag

import (
	"github.com/bborbe/server/renderer"
	"io"
	"strings"
)

type OpeningtagRenderer interface {
	renderer.Renderer
	renderer.Attribute
	renderer.Class
}

type openingtagRenderer struct {
	closed     bool
	name       string
	attributes map[string]string
	classes    map[string]bool
}

func newRenderer(name string, closed bool) *openingtagRenderer {
	r := new(openingtagRenderer)
	r.name = name
	r.closed = closed
	r.attributes = make(map[string]string)
	r.classes = make(map[string]bool)
	return r
}

func NewOpenRenderer(name string) *openingtagRenderer {
	return newRenderer(name, false)
}

func NewCloseRenderer(name string) *openingtagRenderer {
	return newRenderer(name, true)
}

func (v *openingtagRenderer) GetAttribute(key string) string {
	return v.attributes[key]
}

func (v *openingtagRenderer) SetAttribute(key, value string) {
	v.attributes[key] = value
}

func (v *openingtagRenderer) RemoveAttribute(key string) {
	delete(v.attributes, key)
}

func (v *openingtagRenderer) AddClass(class string) {
	if len(class) > 0 {
		v.classes[class] = true
	}
}

func (v *openingtagRenderer) RemoveClass(class string) {
	delete(v.classes, class)
}

func (v *openingtagRenderer) generateAttributes() {
	parts := strings.Split(v.GetAttribute("class"), " ")
	for _, part := range parts {
		v.AddClass(part)
	}
	size := len(v.classes)
	if size > 0 {
		classes := make([]string, size)
		pos := 0
		for k, _ := range v.classes {
			classes[pos] = k
			pos++
		}
		v.SetAttribute("class", strings.Join(classes, " "))
	} else {
		v.RemoveAttribute("class")
	}
}

func (v *openingtagRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(v.name))
	if err != nil {
		return err
	}
	v.generateAttributes()
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
	if v.closed {
		_, err = writer.Write([]byte("/"))
		if err != nil {
			return err
		}
	}
	_, err = writer.Write([]byte(">"))
	if err != nil {
		return err
	}
	return err
}
