package ui

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
	"log"
	"os/exec"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	myPref := pref.Load()

	if myPref.EnableNotificationPopup {
		title := fmt.Sprintf("%s done", kind)
		content := fmt.Sprintf("You just finished a %s pomodoro.", kind)
		err := beeep.Notify(title, content, "/usr/share/pixmaps/fynodoro.png")
		if err != nil {
			log.Printf("unable to display notification bubble: %v\n", err)
		}
	}

	if myPref.NotificationScript != "" {
		cmd := exec.Command("/bin/sh", myPref.NotificationScript)
		if err := cmd.Run(); err != nil {
			log.Printf("unable to run script: %v\n", err)
		}
	}
}
