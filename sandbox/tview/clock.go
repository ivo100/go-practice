// Demo code for a timer based update
package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

const RefreshInterval = 500 * time.Millisecond

var (
	view *tview.Modal
	app  *tview.Application
)

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func updateTime() {
	for {
		time.Sleep(RefreshInterval)
		app.QueueUpdateDraw(func() {
			view.SetText(currentTimeString())
		})
	}
}

func main() {
	app = tview.NewApplication()

	view = tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				fmt.Println("Quit")
				app.Stop()
			}
		})
	/*

		button := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
			app.Stop()
		})
		button.SetBorder(true).SetRect(0, 0, 22, 3)

		flex := tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Left"), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle(currentTimeString()), 0, 3, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
		if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	*/

	go updateTime()
	if err := app.SetRoot(view, false).Run(); err != nil {
		panic(err)
	}
}
