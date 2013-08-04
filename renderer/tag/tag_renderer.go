package tag

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type tagRenderer struct {
	name    string
	content renderer.Renderer
}

func NewTagRenderer(name string) *tagRenderer {
	v := new(tagRenderer)
	v.name = name
	return v
}

func (v *tagRenderer) SetContent(renderer renderer.Renderer) {
	v.content = renderer
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
