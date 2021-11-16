package excel

import (
	"fmt"

	"github.com/haton14/drawexcel/drawio"
)

type coordinates struct {
	x int
	y int
}

const rectangle_template string = `
{
	"type": "%s",
	"color":
	{
		"line": "#4286F4",
		"fill": "#8eb9ff"
	},
	"width": %d,
	"height": %d
}`

func Convert(drawio drawio.Mxfile) string {
	for _, mxcell := range drawio.Diagram.MxGraphModel.Root.MxCells {
		fmt.Println("ID:", mxcell.ID)
		fmt.Println("Style:", mxcell.Style)
		fmt.Println("X:", mxcell.MxGeometry.X)
		fmt.Println("Y:", mxcell.MxGeometry.Y)
		fmt.Println("Width:", mxcell.MxGeometry.Width)
		fmt.Println("Height:", mxcell.MxGeometry.Height)

		object := convertStyle(mxcell.Style)
		if object == "rect" {
			return fmt.Sprintf(rectangle_template, object, mxcell.MxGeometry.Width, mxcell.MxGeometry.Height)
		}
	}
	return ""
}

func convertStyle(style string) string {
	switch style {
	case "whiteSpace=wrap;html=1;aspect=fixed;":
		return "rect"
	}
	return ""
}
