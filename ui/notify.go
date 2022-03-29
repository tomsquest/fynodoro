package ui

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/tomsquest/fynodoro/pomodoro"
)

func notifyPomodoroDone(kind pomodoro.Kind) {
	title := fmt.Sprintf("%s done", kind)
	message := fmt.Sprintf("You just finished a %s pomodoro.", kind)
	_ = beeep.Notify(title, message, "/usr/share/pixmaps/fynodoro.png")

	playNotificationSound()
}
