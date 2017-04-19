package main

import (
	"fmt"

	"image"
	"image/color"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/paulsmith/gogeos/geos"
)

func main() {
	initGeos()
	initGraphics()
}

func initGeos() {
	line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	buf, _ := line.Buffer(2.5)
	fmt.Println(buf)
}

func initGraphics() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	drawExistingImage(gc, "hello.png")
	drawShape(gc)

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}

func drawExistingImage(gc *draw2dimg.GraphicContext, location string) {
	infile, _ := os.Open(location)
	defer infile.Close()
	src, _, _ := image.Decode(infile)
	gc.DrawImage(src)
}

func drawShape(gc *draw2dimg.GraphicContext) {
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()
}
