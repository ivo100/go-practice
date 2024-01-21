// Demo code for a timer based update
package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	view2 *tview.Box
	app2  *tview.Application
)

func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	timeStr := time.Now().Format("Current time is 15:04:05")
	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	return 0, 0, 0, 0
}

func refresh() {
	tick := time.NewTicker(RefreshInterval)
	for {
		select {
		case <-tick.C:
			app2.Draw()
		}
	}
}

func main1() {
	app2 = tview.NewApplication()
	view2 = tview.NewBox().SetDrawFunc(drawTime)

	go refresh()
	if err := app2.SetRoot(view2, true).Run(); err != nil {
		panic(err)
	}
}
