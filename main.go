package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/haton14/drawexcel/drawio"
	"github.com/haton14/drawexcel/excel"
)

func main() {
	drawioFilePath := flag.String("input", "", "drawio file path")
	flag.Parse()
	if *drawioFilePath == "" {
		log.Fatalf("input required ")
	}
	drawioFile, err := os.Open(*drawioFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer drawioFile.Close()
	buf, err := io.ReadAll(drawioFile)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	drawio, err := drawio.Read(buf)
	if err != nil {
		log.Fatalf("failed to read fdrawio file: %s", err)
	}
	format := excel.Convert(drawio)
	excel.Write(format)
}
