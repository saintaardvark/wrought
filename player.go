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
	player.Print()
	player.buildCWSamplesCW()
	player.Print()
	fmt.Printf("[FIXME] Length of samples: %d\n", len(player.samples))
	for _, sample := range player.samples {
		player.Print()
		fmt.Printf("[FIXME] Playing sample of length %d\n", len(*sample))
		// THis does not make a difference:
		// justPlayBeep(beep.NewMusic(""), sample)
		justPlayBeep(player.music, sample)
	}

}

func (player *morsePlayer) buildCWSamplesCW() {
	fmt.Printf("[FIXME] buildCWSamples: ")
	player.Print()
	fmt.Printf("\n")
	//	(morse.EncodeITU(strings.ToLower(s))) // arghhh, was not encoding!
	cw := player.CW()
	//	for _, s := range cw() {
	for _, s := range strings.Split(cw, "") {
		fmt.Printf("[FIXME] buildCWSamples: S: %s\n", s)
		if s == "-" {
			player.buildDah()
		} else if s == "." {
			player.buildDit()
		} else if s == " " {
			// time.Sleep(time.Duration(200 * time.Millisecond))
			player.buildPauseBetweenLetters()
		}
	}
}

func (player *morsePlayer) buildDit() {
	fmt.Printf("[FIXME] buildDit: ")
	player.Print()
	fmt.Printf("\n")
	newSamples := buildABeep(player.music, player.vol, ditLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildDah() {
	fmt.Printf("[FIXME] buildDah: ")
	player.Print()
	fmt.Printf("\n")
	newSamples := buildABeep(player.music, player.vol, dahLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildPauseBetweenLetters() {
	fmt.Printf("[FIXME] buildPause: ")
	player.Print()
	fmt.Printf("\n")
	newSamples := buildABeep(player.music, 0, letterPause, 1, 0.0)
	player.samples = append(player.samples, newSamples)
}
