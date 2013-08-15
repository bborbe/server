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
	SetHead(head renderer.Renderer) HtmlRenderer
	SetBody(body renderer.Renderer) HtmlRenderer
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
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("<!doctype html>\n")
	v.renderer = list.NewListRenderer(contentRenderer, html)
	return v
}

func (v *htmlRenderer) SetHead(head renderer.Renderer) HtmlRenderer {
	v.head.SetRenderer(head)
	return v
}

func (v *htmlRenderer) SetBody(body renderer.Renderer) HtmlRenderer {
	v.body.SetRenderer(body)
	return v
}

func (v *htmlRenderer) Render(writer io.Writer) error {
	var err error
	err = v.renderer.Render(writer)
	if err != nil {
		return err
	}
	return err
}
