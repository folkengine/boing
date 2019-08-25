package main

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
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

const Sonar1 = "/files/jillys-sonar.wav"
const Klaxon1 = "/files/danielnieto7__alert.wav"
const Boing1 = "/files/juskiddink__boing.wav"

func Boing() {
	Play(Boing1)
}

func Klaxon() {
	Play(Klaxon1)
}

func Ping() {
	Play(Sonar1)
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

func main() {
	Boing()
	Ping()
	Klaxon()
}