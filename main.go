package main

import (
	"fmt"
	"math/rand"
	"time"
	"wrought/ham"

	"github.com/dbatbold/beep"
)

var (
	me = ham.Ham{
		Callsign: "VA7UNX",
		Location: "NEW WESTMINSTER BC CANADA",
		Name:     "HUGH",
	}
)

func main() {
	rand.Seed(time.Now().Unix())
	cx := &me
	rx := ham.NewHam()
	player := newMorsePlayer()
	if err := beep.OpenSoundDevice("default"); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}

	if err := beep.InitSoundDevice(); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}
	defer beep.CloseSoundDevice()
	player.exchange = append(player.exchange, initialGreeting(cx, rx))
	player.exchange = append(player.exchange, firstExchange(cx, rx))
	player.exchange = append(player.exchange, secondExchange(cx, rx))
	player.exchange = append(player.exchange, gnightBob(cx, rx))
	player.PrintText()
	player.PlayCW()
	// player.PrintCW()

}
