package html

import (
	"io"
)

type htmlRenderer struct {
}

func NewHtmlRenderer() *htmlRenderer {
	v := new(htmlRenderer)
	return v
}

func (v *htmlRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<!doctype html>"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("<html><head></head><body>"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("</body></html>"))
	if err != nil {
		return err
	}
	return err
}
