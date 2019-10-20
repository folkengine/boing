package boing

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const Chicken3 = "files/chicken_3.wav"
const Klaxon1 = "files/danielnieto7__alert.wav"
const MouthHarp = "files/juskiddink__boing.wav"
const Sonar1 = "files/jillys-sonar.wav"
const Tone1 = "files/A-Tone-His_Self-1266414414.wav"

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

func FilePathInGoPath(relativePath string) (string, bool) {
	env, hasGoPath := os.LookupEnv("GOPATH")
	if !hasGoPath {
		return "", hasGoPath
	}
	return fmt.Sprintf("%s/%s", env, relativePath), hasGoPath
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

func Play(wave string) {
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
