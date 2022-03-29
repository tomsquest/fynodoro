package ui

import (
	"bytes"
	"github.com/faiface/beep"
	"log"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func playNotificationSound() {
	sound := bytes.NewReader(AssetBingWav.StaticContent)

	go func() {
		//streamer, format, err := mp3.Decode(f)
		streamer, format, err := wav.Decode(sound)
		if err != nil {
			log.Fatal(err)
		}
		_ = streamer.Close()

		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			log.Fatal(err)
		}

		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done
	}()
}
