package main

import (
	"github.com/fogleman/gg"
	"image/color"
)

const (
	H  = 200
	W  = 400
	BW = 20
)

var (
	dc    = gg.NewContext(W, H)
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func rect(x, y, w, h float64) {
	dc.NewSubPath()
	dc.MoveTo(x, y)
	dc.LineTo(x+w, y)
	dc.LineTo(x+w, y-h)
	dc.LineTo(x, y-h)
	dc.ClosePath()
}

func bar(x, y, h float64, col color.Color) {
	//h = H - h
	//esi.Println("x ", x, " y ", y, " h ", h)
	rect(x, y, BW, h)
	dc.SetColor(col)
	dc.Fill()
	dc.Stroke()
}

func hline(x, y, w float64) {
	dc.SetLineWidth(1)
	dc.DrawLine(x, y, w, y)
	dc.Stroke()
}

func vline(x, y, h float64) {
	dc.SetLineWidth(1)
	dc.DrawLine(x, y, x, h)
	dc.Stroke()
}

func main() {
	dc.SetColor(white)
	hline(0, 0, W-1)
	vline(0, 0, H-1)
	hline(0, H-1, W-1)
	vline(W-1, 0, H-1)
	//dc.DrawRectangle(20, 20, W-40, H-40)
	//dc.Stroke()
	bot := H - 2.
	left := 40.
	x := left
	bar(x, bot, 20, red)
	x += BW + 2
	bar(x, bot, 30, green)
	x += BW + 2
	bar(x, bot, 35, green)
	dc.SavePNG("out.png")
}
