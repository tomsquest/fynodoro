package ui

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	"os"
	"sync"
)

var once sync.Once
var context *oto.Context

func playNotificationSound() {
	fileName := "/usr/share/fynodoro/notification.mp3"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	once.Do(func() {
		context, err = oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
		if err != nil {
			log.Fatal(err)
		}
	})

	go func() {
		player := context.NewPlayer()
		defer func() {
			err := player.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		if _, err := io.Copy(player, decoder); err != nil {
			log.Fatal(err)
		}
	}()
}
