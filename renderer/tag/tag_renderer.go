package tag

import (
	"io"
)

type tagRenderer struct {
	name string
}

func NewTagRenderer(name string) *tagRenderer {
	v := new(tagRenderer)
	v.name = name
	return v
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
