package main

import (
	"github.com/evandrojr/string-interpolation/esi"
	"github.com/fogleman/gg"
	"image/color"
)

const (
	H  = 200
	W  = 400
	BW = 10  // bar width
	BG = 10. // horizontal gap between bars
)

var (
	dc    = gg.NewContext(W, H)
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

func rect(x, y, w, h float64) {
	dc.NewSubPath()
	y1 := H - y
	//dc.MoveTo(x, y)
	//dc.LineTo(x+w, y)
	//dc.LineTo(x+w, y+h)
	//dc.LineTo(x, y+h)
	//dc.ClosePath()

	dc.MoveTo(x, y)
	dc.LineTo(x+w, y)
	dc.LineTo(x+w, y1-h)
	dc.LineTo(x, y)
	dc.ClosePath()
}

func bar(x, y, h float64, col color.Color) {
	//h = H - h
	esi.Println("x ", x, " y ", y, " h ", h)
	rect(x, y, BW, y+h)
	dc.SetColor(col)
	dc.Fill()
}

func hline(x, y, w float64) {
	dc.SetLineWidth(1)
	dc.DrawLine(x, y, w, y)
}

func vline(x, y, h float64) {
	dc.SetLineWidth(1)
	dc.DrawLine(x, y, x, h)
}

func main() {
	// x, y is top left corner, y grows down
	// .-->
	// |
	// v
	dc.SetColor(white)
	//hline(0, 0, W-1)
	//vline(0, 0, H-1)
	//hline(0, H-1, W-1)
	//vline(W-1, 0, H-1)
	//dc.Stroke()

	rect(1, 1, W-1, H-1)
	dc.DrawRectangle(0, 0, W-1, H-1)
	dc.Stroke()

	//left := 40.
	//step := BW + BG
	//y := H * 0.75
	//x := left
	//bar(x, y, 20, red)
	//x += step
	////y += 10
	//bar(x, y, 30, green)
	//x += step
	//y += 10
	//bar(x, y, 35, green)
	dc.SavePNG("out.png")
}
