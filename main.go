package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"time"
)

func main() {
	myApp := app.NewWithID("com.tomquest.fynodoro")
	myApp.Settings().SetTheme(&myTheme{})

	myWin := myApp.NewWindow("Fynodoro")
	myWin.SetIcon(resourceIconPng)
	myWin.CenterOnScreen()

	myCanvas := myWin.Canvas()

	timerDuration := 25 * 60 * time.Second
	go updateTimer(myCanvas, timerDuration)

	myWin.ShowAndRun()
}

func updateTimer(myCanvas fyne.Canvas, timerDuration time.Duration) {
	displayTimer(myCanvas, timerDuration)

	timer := time.NewTimer(timerDuration)
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				timerDuration -= time.Second
				displayTimer(myCanvas, timerDuration)
			case <-timer.C:
				return
			}
		}
	}()
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}

func displayTimer(myCanvas fyne.Canvas, timeLeft time.Duration) {
	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	text := canvas.NewText(formatDuration(timeLeft), red)
	text.TextStyle.Bold = true
	text.TextSize = 100
	myCanvas.SetContent(text)
}
