package boing

import (
	"errors"
	"fmt"
	"github.com/electronicpanopticon/gobrick"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const Chickens = "files/chickens.wav"
const Chicken2 = "files/chicken2.wav"
const Chicken3 = "files/chicken_3.wav"
const Chicken4 = "files/chicken_4.wav"
const ChickensInBarn = "files/Chickens_in_Barn_02.wav"
const Rooster = "files/384188__inspectorj__rooster-crowing-a_R.wav"
const Klaxon1 = "files/danielnieto7__alert.wav"
const MouthHarp = "files/juskiddink__boing.wav"
const Sonar1 = "files/jillys-sonar.wav"
const Tone1 = "files/A-Tone-His_Self-1266414414.wav"

func Boing() {
	Play(MouthHarp)
}

func Chicken() {
	Play(Chickens, Chicken2, Chicken3, Chicken4, ChickensInBarn, Rooster)
}

func Crow() {
	Play(Rooster)
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

// TODO: Cleanup sig
func FilePathInGoPath(relativePath string) (string, bool) {
	gopath := gobrick.GetGOPATH()
	return fmt.Sprintf("%s/%s", gopath, relativePath), true
}

func FileInGoPath(relativePath string) (*os.File, error) {
	fp, hasGoPath := FilePathInGoPath(relativePath)
	if !hasGoPath {
		myError := errors.New("GOPATH not set")
		return nil, myError
	}
	return os.Open(fp)
}

func relativeFilePath(filename string) string {
	return fmt.Sprintf("src/github.com/folkengine/boing/%s", filename)
}

func Play(waves ...string) {

	rand.Seed(time.Now().Unix())
	wave := waves[rand.Intn(len(waves))]

	f, err := FileInGoPath(relativeFilePath(wave))
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
