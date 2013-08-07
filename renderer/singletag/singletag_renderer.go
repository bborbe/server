package singletag

import (
	"github.com/bborbe/server/renderer"
	"io"
)

type SingletagRenderer interface {
	renderer.Renderer
}

type singletagRenderer struct {
	name string
}

func NewSingletagRenderer(name string) *singletagRenderer {
	v := new(singletagRenderer)
	v.name = name
	return v
}

func (v *singletagRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(v.name))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("/>"))
	if err != nil {
		return err
	}
	return err
}
