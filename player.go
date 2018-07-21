package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

type morsePlayer struct {
	music     *beep.Music
	exchange  []string
	freqHertz float64
	vol       int
}

func newMorsePlayer() *morsePlayer {
	player := morsePlayer{
		music:     beep.NewMusic(""),
		exchange:  []string{},
		freqHertz: 500,
		vol:       80,
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

func (player *morsePlayer) buildDit() *[]int16 {
	return buildABeep(player.music, player.vol, 150, 1, player.freqHertz)
}

func (player *morsePlayer) buildDah() *[]int16 {
	return buildABeep(player.music, player.vol, 300, 1, player.freqHertz)
}
