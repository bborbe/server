package body

import (
	"io"
)

type bodyRenderer struct {
}

func NewBodyRenderer() *bodyRenderer {
	v := new(bodyRenderer)
	return v
}

func (v *bodyRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<body>"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("</body>"))
	if err != nil {
		return err
	}
	return err
}
