package table

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/content"
	"github.com/bborbe/server/renderer/tablecell"
	"github.com/bborbe/server/renderer/tablerow"
)

func TestImplementsRenderer(t *testing.T) {
	v := NewTableRenderer()
	var i (*renderer.Renderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsTableRenderer(t *testing.T) {
	v := NewTableRenderer()
	var i (*TableRenderer) = nil
	err := AssertThat(v, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	var err error
	v := NewTableRenderer()
	writer := bytes.NewBufferString("")
	err = v.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Startswith("<table"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Endswith("</table>"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompleteTable(t *testing.T) {
	var err error

	cell := tablecell.NewTablecellRenderer()
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContentString("hello world")
	cell.SetContent(contentRenderer)

	row := tablerow.NewTablerowRenderer()
	row.AddCell(cell)

	table := NewTableRenderer()
	table.AddRow(row)

	writer := bytes.NewBufferString("")
	err = table.Render(writer)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("<table><tr><td>hello world</td></tr></table>"))
	if err != nil {
		t.Fatal(err)
	}

}
