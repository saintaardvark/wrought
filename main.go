package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"wrought/ham"
	"wrought/morsePlayer"

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
	buildQSO(cx, rx, player)

	app.Commands = []cli.Command{
		{
			Name:  "play",
			Usage: "play a qso",
			Action: func(c *cli.Context) error {
				player.PlayCW()
				return nil
			},
		},
		{
			Name:  "print",
			Usage: "print a qso",
			Action: func(c *cli.Context) error {
				player.PrintText()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
