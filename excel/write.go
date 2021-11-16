package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Write(format string) {
	f := excelize.NewFile()
	f.AddShape("Sheet1", "G6", `
	{
		"type": "rect",
		"color":
		{
			"line": "#4286F4",
			"fill": "#8eb9ff"
		},
		"width": 180,
		"height": 180
	}`)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
