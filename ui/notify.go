package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/tomsquest/fynodoro/pomodoro"
	"log"
	"os/exec"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	title := fmt.Sprintf("%s done", kind)
	content := fmt.Sprintf("You just finished a %s pomodoro.", kind)
	notification := fyne.NewNotification(title, content)
	fyne.CurrentApp().SendNotification(notification)

	script := "/home/tom/Dev/fynodoro/assets/notify.sh"
	cmd := exec.Command("/bin/sh", script)
	if err := cmd.Run(); err != nil {
		log.Printf("unable to run script: %v\n", err)
	}
}
