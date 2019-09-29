package boing

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func getDir() string {
	env, hasGoPath := os.LookupEnv("GOPATH")
	if !hasGoPath {
		panic("GOPATH not set.")
	}
	return fmt.Sprintf("%s/src/github.com/folkengine/boing", env)
}

const Chicken3 = "/files/chicken_3.wav"
const Klaxon1 = "/files/danielnieto7__alert.wav"
const MouthHarp = "/files/juskiddink__boing.wav"
const Sonar1 = "/files/jillys-sonar.wav"
const Tone1 = "/files/A-Tone-His_Self-1266414414.wav"

func Boing() {
	Play(MouthHarp)
}

func Chicken() {
	Play(Chicken3)
}
func Klaxon() {
	Play(Klaxon1)
}

func Ping() {
	Play(Sonar1)
}

func Tone() {
	Play(Tone1)
}

func Play(wave string) {
	filepath := fmt.Sprintf("%s%s", getDir(), wave)
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
