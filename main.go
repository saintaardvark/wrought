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
	player.Print()
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
	// printMorse(initialGreeting(cx, rx))
	// playMorse(firstExchange(cx, rx))
	// playMorse(secondExchange(cx, rx))
	// playMorse(gnightBob(cx, rx))
	player.PrintCW()
	player.PrintText()
	player.PlayCW()

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

func initialGreeting(caller, receiver *ham.Ham) string {
	callerRepeat := fmt.Sprintf("%s %s %s", caller.Callsign, caller.Callsign, caller.Callsign)
	receiverRepeat := fmt.Sprintf("%s %s %s", receiver.Callsign, receiver.Callsign, receiver.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", callerRepeat, caller.Callsign, receiverRepeat)
	return msg
}

func firstExchange(caller, receiver *ham.Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FOR CALL BT UR RST 599 599 HR "
	msg = msg + qth(caller.Location) + " "
	msg = msg + name(caller.Name) + " "
	msg = msg + "HW CPY? "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func secondExchange(caller, receiver *ham.Ham) string {
	msg := de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FOR RPT SLD CPY FB UR RST 599 599 BT" + " "
	msg = msg + name(receiver.Name) + " "
	msg = msg + qth(receiver.Location) + " "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func gnightBob(caller, receiver *ham.Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FER FB QSO " + receiver.Name + " "
	msg = msg + "HP CU AGN BT VY 73 TO U ES URS SK" + " "
	msg = msg + de(receiver.Callsign, caller.Callsign) + "\n"
	// And now the reply
	msg = msg + de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FER QSO " + caller.Name + " "
	msg = msg + "BCNU BT VY 73 TO U ES URS SK" + " "
	msg = msg + de(caller.Callsign, receiver.Callsign)
	return msg
}

func name(name string) string {
	return fmt.Sprintf("NAME %s %s", name, name)
}

func qth(location string) string {
	return fmt.Sprintf("QTH %s %s", location, location)
}

func de(cx, rx string) string {
	return fmt.Sprintf("%s DE %s", cx, rx)
}

func kn(cx, rx string) string {
	return fmt.Sprintf("%s KN", de(cx, rx))
}

func printMorse(msg string) {
	fmt.Println(morse.EncodeITU(msg))
}
