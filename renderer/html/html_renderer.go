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

type htmlRenderer struct {
	renderer renderer.Renderer
	body     placeholder.PlaceholderRenderer
	head     placeholder.PlaceholderRenderer
}

func NewHtmlRenderer() *htmlRenderer {
	v := new(htmlRenderer)
	headPlaceholder := placeholder.NewPlaceholderRenderer()
	headPlaceholder.SetRenderer(head.NewHeadRenderer())
	bodyPlaceholder := placeholder.NewPlaceholderRenderer()
	bodyPlaceholder.SetRenderer(body.NewBodyRenderer())
	html := tag.NewTagRenderer("html")
	html.SetContent(list.NewListRenderer(headPlaceholder, bodyPlaceholder))
	doctype := content.NewContentRenderer("<!doctype html>")
	v.renderer = list.NewListRenderer(doctype, html)
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
