package drawio

import (
	"encoding/xml"
)

type Mxfile struct {
	Diagram Diagram `xml:"diagram"`
}

type Diagram struct {
	MxGraphModel MxGraphModel `xml:"mxGraphModel"`
}

type MxGraphModel struct {
	Root Root `xml:"root"`
}

type Root struct {
	MxCells []MxCell `xml:"mxCell"`
}
type MxCell struct {
	ID         int        `xml:"id,attr"`
	Style      string     `xml:"style,attr"`
	MxGeometry MxGeometry `xml:"mxGeometry"`
}

type MxGeometry struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

func Read(buf []byte) (Mxfile, error) {
	drawio := Mxfile{}
	err := xml.Unmarshal(buf, &drawio)
	return drawio, err

}
