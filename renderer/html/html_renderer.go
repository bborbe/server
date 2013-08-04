package html

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/body"
	"github.com/bborbe/server/renderer/head"
	"io"
)

type htmlRenderer struct {
	body renderer.Renderer
	head renderer.Renderer
}

func NewHtmlRenderer() *htmlRenderer {
	v := new(htmlRenderer)
	v.body = body.NewBodyRenderer()
	v.head = head.NewHeadRenderer()
	return v
}

func (v *htmlRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<!doctype html>"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("<html>"))
	if err != nil {
		return err
	}
	err = v.head.Render(writer)
	if err != nil {
		return err
	}
	err = v.body.Render(writer)
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("</html>"))
	if err != nil {
		return err
	}
	return err
}
