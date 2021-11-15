package main

import (
	"fmt"
	"io"
	"os"

	"github.com/haton14/drawexcel/drawio"
	"github.com/haton14/drawexcel/excel"
)

func main() {
	drawioFile, err := os.Open("tmp.drawio")
	if err != nil {
		fmt.Printf("cannot open file: %v", err)
		return
	}
	defer drawioFile.Close()
	buf, err := io.ReadAll(drawioFile)
	if err != nil {
		fmt.Printf("cannot read file: %v", err)
		return
	}
	drawio, err := drawio.Read(buf)
	if err != nil {
		fmt.Printf("cannot read drawio file: %v", err)
		return
	}
	format := excel.Convert(drawio)
	excel.Write(format)

}
