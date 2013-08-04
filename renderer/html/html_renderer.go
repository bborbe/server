package html

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/body"
	"github.com/bborbe/server/renderer/content"
	"github.com/bborbe/server/renderer/head"
	"github.com/bborbe/server/renderer/list"
	"github.com/bborbe/server/renderer/placeholder"
	"github.com/bborbe/server/renderer/tag"
	"io"
)

type HtmlRenderer interface {
	renderer.Renderer
	SetHead(head renderer.Renderer)
	SetBody(body renderer.Renderer)
}

type htmlRenderer struct {
	renderer renderer.Renderer
	body     placeholder.PlaceholderRenderer
	head     placeholder.PlaceholderRenderer
}

func NewHtmlRenderer() *htmlRenderer {
	v := new(htmlRenderer)
	v.head = placeholder.NewPlaceholderRenderer()
	v.head.SetRenderer(head.NewHeadRenderer())
	v.body = placeholder.NewPlaceholderRenderer()
	v.body.SetRenderer(body.NewBodyRenderer())
	html := tag.NewTagRenderer("html")
	html.SetContent(list.NewListRenderer(v.head, v.body))
	v.renderer = list.NewListRenderer(content.NewContentRenderer("<!doctype html>"), html)
	return v
}

func (v *htmlRenderer) SetHead(head renderer.Renderer) {
	v.head.SetRenderer(head)
}

func (v *htmlRenderer) SetBody(body renderer.Renderer) {
	v.body.SetRenderer(body)
}

func (v *htmlRenderer) Render(writer io.Writer) error {
	var err error
	err = v.renderer.Render(writer)
	if err != nil {
		return err
	}
	return err
}
