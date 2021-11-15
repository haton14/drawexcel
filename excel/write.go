package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Write(format string) {
	f := excelize.NewFile()
	f.AddShape("Sheet1", "G6", format)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
