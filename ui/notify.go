package ui

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/tomsquest/fynodoro/pomodoro"
	"log"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	title := fmt.Sprintf("%s done", kind)
	message := fmt.Sprintf("You just finished a %s pomodoro.", kind)
	err := beeep.Notify(title, message, "/usr/share/pixmaps/fynodoro.png")
	if err != nil {
		log.Printf("unable to notify: %v\n", err)
	}

	err = playNotificationSound()
	if err != nil {
		log.Printf("unable to play notification sound: %v\n", err)
	}
}
