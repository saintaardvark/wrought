package main

import (
	"fmt"
	"strings"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

const (
	ditLength   = 150
	dahLength   = 300
	letterPause = 3 * ditLength
	wordPause   = 7 * ditLength
	freq        = 500
	volume      = 80
)

var (
	prosigns = map[string]bool{
		"CQ": true,
	}
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
		freqHertz: beep.HertzToFreq(freq),
		vol:       volume,
		samples:   []*[]int16{},
	}
	return &player
}

func (player *morsePlayer) Print() {
	fmt.Printf("Player: freq: %f, vol: %d\n", player.freqHertz, player.vol)
}

func (player *morsePlayer) PrintCW() {
	for _, s := range player.exchange {
		fmt.Println(morse.EncodeITU(strings.ToLower(s)))
	}
}

func (player *morsePlayer) CW() string {
	var cw string
	for _, s := range player.exchange {
		cw = fmt.Sprintf("%s\n%s", cw, morse.EncodeITU(strings.ToLower(s)))
	}
	return cw
}

func (player *morsePlayer) PrintText() {
	for _, s := range player.exchange {
		fmt.Println(s)
	}
}

func (player *morsePlayer) PlayCW() {
	player.buildCWSamplesCW()
	for _, sample := range player.samples {
		justPlayBeep(player.music, sample)
	}
}

func (player *morsePlayer) buildCWSamplesCW() {
	//	(morse.EncodeITU(strings.ToLower(s))) // arghhh, was not encoding!
	cw := player.CW()
	//	for _, s := range cw() {
	// FIXME: Trying to detect prosign *after* CW has been encoded...
	for _, word := range strings.Split(cw, " ") {
		fmt.Printf("[FIXME] Building %s\n", word)
		if prosigns[word] == true {
			player.buildProsign(word)
		} else {
			player.buildWord(word)
		}
		player.buildPause(false)
	}
}

func (player *morsePlayer) buildWord(word string) {
	for _, s := range strings.Split(word, "") {
		if s == "-" {
			player.buildDah()
		} else if s == "." {
			player.buildDit()
		} else if s == " " {
			// time.Sleep(time.Duration(200 * time.Millisecond))
			player.buildPause(false)
		}
	}
}

func (player *morsePlayer) buildProsign(prosign string) {
	for _, s := range strings.Split(prosign, "") {
		if s == "-" {
			player.buildDah()
		} else if s == "." {
			player.buildDit()
		}
	}
}

func (player *morsePlayer) buildDit() {
	newSamples := buildABeep(player.vol, ditLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildDah() {
	newSamples := buildABeep(player.vol, dahLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildPause(word bool) {
	var newSamples *[]int16
	if word == true {
		newSamples = buildABeep(0, wordPause, 1, 0.0)
	} else {
		newSamples = buildABeep(0, letterPause, 1, 0.0)
	}
	player.samples = append(player.samples, newSamples)
}
