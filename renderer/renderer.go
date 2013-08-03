package renderer

import "io"

type Renderer interface {
	Render(writer io.Writer) error
}
