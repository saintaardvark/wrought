package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

// List of prosigns
// List of callsigns
// List of countries, states, cities
// List of phrases

// Ham is a ham
type Ham struct {
	Callsign string
	Location string
	Name     string
}

var (
	callsignFile = "/home/aardvark/.qrq/toplist"
	citiesFile   = "./data/world-cities/data/world-cities.csv"
	me           = Ham{
		Callsign: "VA7UNX",
		Location: "NEW WESTMINSTER BC CANADA",
		Name:     "HUGH",
	}
)

func main() {
	rand.Seed(time.Now().Unix())
	cx := &me
	rx := newHam()
	player := newMorsePlayer()
	if err := beep.OpenSoundDevice("default"); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}

	if err := beep.InitSoundDevice(); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}
	defer beep.CloseSoundDevice()
	// playMorse("VA7UNX")
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
	fmt.Println("[FIXME] About to encode " + s)
	cw := morse.EncodeITU(strings.ToLower(s))
	fmt.Println("[FIXME] " + cw)
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

func initialGreeting(caller, receiver *Ham) string {
	callerRepeat := fmt.Sprintf("%s %s %s", caller.Callsign, caller.Callsign, caller.Callsign)
	receiverRepeat := fmt.Sprintf("%s %s %s", receiver.Callsign, receiver.Callsign, receiver.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", callerRepeat, caller.Callsign, receiverRepeat)
	return msg
}

func firstExchange(caller, receiver *Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FOR CALL BT UR RST 599 599 HR "
	msg = msg + qth(caller.Location) + " "
	msg = msg + name(caller.Name) + " "
	msg = msg + "HW CPY? "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func secondExchange(caller, receiver *Ham) string {
	msg := de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FOR RPT SLD CPY FB UR RST 599 599 BT" + " "
	msg = msg + name(receiver.Name) + " "
	msg = msg + qth(receiver.Location) + " "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func gnightBob(caller, receiver *Ham) string {
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

func readCallsigns() *[]string {
	var callsigns []string
	file, err := os.Open(callsignFile)
	if err != nil {
		fmt.Printf("Can't read %s: %s", callsignFile, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cs := strings.SplitN(scanner.Text(), " ", 2)
		// fmt.Printf("%+s\n", cs[0])
		callsigns = append(callsigns, cs[0])
	}
	// fmt.Printf("%+s\n", callsigns)
	return &callsigns
}

func getRandomCallsign() string {
	cs := readCallsigns()
	return (*cs)[rand.Intn(len(*cs))]
}

func readCities() *[]string {
	var cities []string
	file, err := os.Open(citiesFile)
	if err != nil {
		fmt.Printf("Can't read %s: %s", citiesFile, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cityData := strings.SplitN(scanner.Text(), ",", 4)
		// fmt.Printf("%+s\n", cs[0])
		cities = append(cities, fmt.Sprintf("%s %s %s", cityData[0], cityData[2], cityData[1]))
	}
	// fmt.Printf("%+s\n", callsigns)
	return &cities

}

func getRandomCity() string {
	cities := readCities()
	return (*cities)[rand.Intn(len(*cities))]
}

func newHam() *Ham {
	ham := Ham{
		Callsign: getRandomCallsign(),
		Location: getRandomCity(),
		Name:     "JANE",
	}
	return &ham
}
