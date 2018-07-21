package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

const (
	ditLength   = 150
	dahLength   = 300
	letterPause = 3 * ditLength
)

type morsePlayer struct {
	music     *beep.Music
	exchange  []string
	freqHertz float64
	vol       int
	samples   []*[]int16
}

func newMorsePlayer() *morsePlayer {
	player := morsePlayer{
		music:     beep.NewMusic(""),
		exchange:  []string{},
		freqHertz: 500,
		vol:       80,
		samples:   []*[]int16{},
	}
	return &player
}

func (player *morsePlayer) PrintCW() {
	for _, s := range player.exchange {
		fmt.Println(morse.EncodeITU(strings.ToLower(s)))
	}
}

func (player *morsePlayer) PrintText() {
	for _, s := range player.exchange {
		fmt.Println(s)
	}
}
func (player *morsePlayer) PlayCW() {
	for _, s := range player.exchange {
		for _, letter := range strings.Split(s, "") {
			if letter == "-" {
				player.buildDah()
			} else if letter == "." {
				player.buildDit()
			} else if letter == " " {
				time.Sleep(time.Duration(200 * time.Millisecond))
			}
		}
	}
}

func (player *morsePlayer) buildDit() {
	newSamples := buildABeep(player.music, player.vol, ditLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildDah() {
	newSamples := buildABeep(player.music, player.vol, dahLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildPauseBetweenLetters() {
	newSamples := buildABeep(player.music, 0, letterPause, 1, 0.0)
	player.samples = append(player.samples, newSamples)
}
