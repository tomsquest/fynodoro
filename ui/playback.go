package ui

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"os"
)

func playNotificationSound() error {
	fileName := "/usr/share/fynodoro/notification.mp3"
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("unable to read %s: %w", fileName, err)
	}

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		return fmt.Errorf("unable to decode %s: %w", fileName, err)
	}

	context, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return fmt.Errorf("unable to create oto context: %w", err)
	}
	<-ready

	go func() {
		player := context.NewPlayer(decoder)
		defer func() {
			err = player.Close()
		}()

		player.Play()
	}()
	if err != nil {
		return fmt.Errorf("unable to play: %w", err)
	}

	return nil
}
