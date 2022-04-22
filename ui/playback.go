package ui

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"os"
	"sync"
)

var once sync.Once
var context *oto.Context

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

	once.Do(func() {
		context, err = oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
	})
	if err != nil {
		return fmt.Errorf("unable to create oto context: %w", err)
	}

	go func() {
		player := context.NewPlayer()
		defer func() {
			err = player.Close()
		}()

		_, err = io.Copy(player, decoder)
	}()
	if err != nil {
		return fmt.Errorf("unable to play: %w", err)
	}

	return nil
}
