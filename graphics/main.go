package main

import (
	"github.com/fogleman/gg"
	"image/color"
	"time"
)

const (
	H  = 300
	W  = 500
	BW = 16  // bar width
	BG = 4.  // horizontal gap between bars
	CI = 196 // color intensity
)

var (
	dc     = gg.NewContext(W, H)
	black  = color.RGBA{0, 0, 0, 255}
	white  = color.RGBA{CI, CI, CI, 255}
	red    = color.RGBA{CI, 0, 0, 255}
	green  = color.RGBA{0, CI, 0, 255}
	blue   = color.RGBA{0, 0, CI, 255}
	cyan   = color.RGBA{0, CI, CI, 255}
	tStart = time.Now()
)

type Candle struct {
	Time   time.Time
	High   float64
	Low    float64
	Open   float64
	Close  float64
	Volume float64
}

func DrawCandle(c Candle) {
	left := TimeToX(c.Time)
	center := left + BW/2.
	//esi.Println("DrawCandle open ", c.Open, " close ", c.Close, " high ", c.High, " low ", c.Low)
	if c.Open < c.Close {
		dc.SetColor(green)
		rect(left, c.Open, BW, c.Close-c.Open)
	} else {
		dc.SetColor(red)
		rect(left, c.Close, BW, c.Open-c.Close)
	}
	dc.Fill()
	dc.SetLineWidth(2)
	vline(center, c.Low, c.High-c.Low)

	if c.Volume > 0 {
		dc.SetColor(cyan)
		bar(left, 1, c.Volume)
	}
	dc.SetColor(white)
}

func rect(x, y, w, h float64) {
	dc.NewSubPath()
	y1 := H - y
	dc.MoveTo(x, y1)
	dc.LineTo(x+w, y1)
	dc.LineTo(x+w, y1-h)
	dc.LineTo(x, y1-h)
	dc.ClosePath()
}

func bar(x, y, h float64) {
	//esi.Println("x ", x, " y ", y, " h ", h)
	rect(x, y, BW, h)
	dc.Fill()
}

func hline(x, y, w float64) {
	dc.DrawLine(x, y, w, y)
	dc.Stroke()
}

func vline(x, y, h float64) {
	dc.DrawLine(x, H-y-h, x, H-y)
	dc.Stroke()
}

func main() {
	// gg context
	// x, y is top left corner, y grows down
	// .-->
	// |
	// v
	// reversed to x, y bottom left, y grows up (natural)

	dc.SetColor(white)
	rect(0, 0, W, H)
	dc.Stroke()
	c1 := Candle{
		Time:   tStart.Add(5 * time.Minute),
		High:   200.,
		Low:    100.,
		Open:   120,
		Close:  180,
		Volume: 50,
	}
	c2 := Candle{
		Time:   tStart.Add(10 * time.Minute),
		High:   210.,
		Low:    120.,
		Open:   140,
		Close:  130,
		Volume: 30,
	}
	c3 := Candle{
		Time:   tStart.Add(20 * time.Minute),
		High:   190.,
		Low:    100.,
		Open:   140,
		Close:  115,
		Volume: 25,
	}

	DrawCandle(c1)
	DrawCandle(c2)
	DrawCandle(c3)
	c1.Time = tStart.Add(25 * time.Minute)
	c1.Volume = 90
	DrawCandle(c1)
	dc.SavePNG("out.png")
}

func TimeToX(tm time.Time) float64 {
	diff := tm.Sub(tStart).Minutes() / 5
	x := diff * (BW + BG)
	return x
}
