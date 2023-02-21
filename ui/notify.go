package ui

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/tomsquest/fynodoro/pomodoro"
	"log"
	"os/exec"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	title := fmt.Sprintf("%s done", kind)
	message := fmt.Sprintf("You just finished a %s pomodoro.", kind)
	err := beeep.Notify(title, message, "/usr/share/pixmaps/fynodoro.png")
	if err != nil {
		log.Printf("unable to display notification bubble: %v\n", err)
	}

	script := "/home/tom/Dev/fynodoro/assets/notify.sh"
	cmd := exec.Command("/bin/sh", script)
	if err := cmd.Run(); err != nil {
		log.Printf("unable to run script: %v\n", err)
	}
}
