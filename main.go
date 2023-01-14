package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("03:04:05")
	clock.SetText(formattedTime)
}

func main() {
	a := app.New()
	win := a.NewWindow("Go Clock")

	clock := widget.NewLabel("")

	updateTime(clock)

	win.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	win.ShowAndRun()
}
