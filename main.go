package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"wrought/ham"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

// List of prosigns
// List of callsigns
// List of countries, states, cities
// List of phrases

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
	playMorse("CQ")
	player.exchange = append(player.exchange, initialGreeting(cx, rx))
	player.exchange = append(player.exchange, firstExchange(cx, rx))
	player.exchange = append(player.exchange, gnightBob(cx, rx))
	player.PlayCW()
	player.PrintCW()
	player.PrintText()
}

func playMorse(s string) {
	cw := morse.EncodeITU(strings.ToLower(s))
	for _, letter := range strings.Split(cw, "") {
		if letter == "-" {
			dah()
		} else if letter == "." {
			dit()
		} else if letter == " " {
			time.Sleep(time.Duration(200 * time.Millisecond))
		}
	}
}

func printMorse(msg string) {
	fmt.Println(morse.EncodeITU(msg))
}
