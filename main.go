package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"wrought/ham"
	"wrought/morsePlayer"
	"wrought/qso"

	"github.com/dbatbold/beep"
	"github.com/urfave/cli"
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
	player := morsePlayer.NewMorsePlayer()
	app := cli.NewApp()
	app.Name = "wrought"
	app.Usage = "CW trainer"

	if err := beep.OpenSoundDevice("default"); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}

	if err := beep.InitSoundDevice(); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}
	defer beep.CloseSoundDevice()
	qso := qso.BuildQSO(cx, rx, player)

	app.Commands = []cli.Command{
		{
			Name:  "play",
			Usage: "play a qso",
			Action: func(c *cli.Context) error {
				qso.PlayCW(player)
				return nil
			},
		},
		{
			Name:  "print",
			Usage: "print a qso",
			Action: func(c *cli.Context) error {
				qso.PrintText()
				return nil
			},
		},
		{
			Name:  "half",
			Usage: "Play remote half of conversation with pauses, so you can practice keying",
			Action: func(c *cli.Context) error {
				qso.PlayRemoteHalf()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
