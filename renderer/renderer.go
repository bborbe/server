package renderer

import "io"

type Renderer interface {
	Render(writer io.Writer) error
}

type Attribute interface {
	SetAttribute(key, value string)
	RemoveAttribute(key string)
}

type Class interface {
	AddClass(class string)
	RemoveClass(class string)
}

type RendererWithClass interface {
	Renderer
	Class
}
