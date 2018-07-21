package main

import (
	"fmt"
	"strings"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

type morsePlayer struct {
	music    *beep.Music
	exchange []string
}

func newMorsePlayer() *morsePlayer {
	player := morsePlayer{
		music:    beep.NewMusic(""),
		exchange: []string{},
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
