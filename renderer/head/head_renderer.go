package head

import (
	"io"
)

type headRenderer struct {
}

func NewHeadRenderer() *headRenderer {
	v := new(headRenderer)
	return v
}

func (v *headRenderer) Render(writer io.Writer) error {
	var err error
	_, err = writer.Write([]byte("<head>"))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte("</head>"))
	if err != nil {
		return err
	}
	return err
}
