package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/haton14/drawexcel/drawio"
	"github.com/haton14/drawexcel/excel"
)

func main() {
	drawioFilePath := flag.String("input", "", "drawio file path")
	flag.Parse()
	if *drawioFilePath == "" {
		fmt.Println("required -input.")
		return
	}
	drawioFile, err := os.Open(*drawioFilePath)
	if err != nil {
		fmt.Printf("cannot open file: %v\n", err)
		return
	}
	defer drawioFile.Close()
	buf, err := io.ReadAll(drawioFile)
	if err != nil {
		fmt.Printf("cannot read file: %v\n", err)
		return
	}
	drawio, err := drawio.Read(buf)
	if err != nil {
		fmt.Printf("cannot read drawio file: %v\n", err)
		return
	}
	format := excel.Convert(drawio)
	excel.Write(format)

}
