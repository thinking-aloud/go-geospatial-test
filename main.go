package main

import (
	"fmt"

	"image"
	"image/color"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/paulsmith/gogeos/geos"
)

var gc *draw2dimg.GraphicContext

func main() {
	dest := initGraphics()

	drawImage("res/white.png")
	drawPolygon()

	geosTryout()
	clipShp("res/a.shp", "res/b.shp")

	// Save to file
	draw2dimg.SaveToPngFile("output.png", dest)
}

func initGraphics() *image.RGBA {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	gc = draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	return dest
}

func geosTryout() {
	line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	buf, _ := line.Buffer(2.5)
	fmt.Println(buf)
}

func clipShp(a string, b string) {
	var err error
	shpA, err := os.Open(a)
	if err != nil {
		fmt.Println(err)
	}
	defer shpA.Close()

	//imageA, _, _ := image.Decode(shpA)

}

func drawImage(location string) {
	infile, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		fmt.Println(err)
	}

	gc.DrawImage(src)
}

func drawPolygon() {
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()
}
