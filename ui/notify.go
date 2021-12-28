package ui

import (
	"fmt"

	"github.com/tomsquest/fynodoro/pomodoro"

	"fyne.io/fyne/v2"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	title := fmt.Sprintf("%s done", kind)
	message := fmt.Sprintf("You just finished a %s pomodoro.", kind)
	fyne.CurrentApp().SendNotification(fyne.NewNotification(title, message))
}
