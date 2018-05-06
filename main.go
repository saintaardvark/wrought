package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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
	fmt.Printf(initialGreeting(cx, rx))
	// printMorse(initialGreeting(cx, rx))
	fmt.Println(firstExchange(rx, cx))
	fmt.Println(secondExchange(rx, cx))
	fmt.Println(gnightBob(cx, rx))
}

func newHam() *Ham {
	ham := Ham{
		Callsign: getRandomCallsign(),
		Location: getRandomCity(),
		Name:     "JANE",
	}
	return &ham
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
	msg = msg + kn(receiver.Callsign, caller.Callsign)
	return msg
}

func secondExchange(caller, receiver *Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FOR RPT SLD CPY FB UR RST 599 599 BT" + " "
	msg = msg + name(receiver.Name) + " "
	msg = msg + qth(receiver.Location) + " "
	msg = msg + kn(receiver.Callsign, caller.Callsign)
	return msg
}

func gnightBob(caller, receiver *Ham) string {
	msg := de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FER FB QSO " + receiver.Name + " "
	msg = msg + "HP CU AGN BT VY 73 TO U ES URS SK" + " "
	msg = msg + de(caller.Callsign, receiver.Callsign)
	msg = de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FER FB QSO " + receiver.Name + " "
	msg = msg + "HP CU AGN BT VY 73 TO U ES URS SK" + " "
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
		city, country, subcountry, _ := strings.SplitN(scanner.Text(), ",", 4)
		// fmt.Printf("%+s\n", cs[0])
		cities = append(cities, fmt.Sprintf("%s %s %s", city, country, subcountry))
	}
	// fmt.Printf("%+s\n", callsigns)
	return &cities

}

func getRandomCity() string {
	cities := readCities()
	return (*cities)[rand.Intn(len(*cities))]
}
