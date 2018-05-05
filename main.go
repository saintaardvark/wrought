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

var (
	callsignFile = "/home/aardvark/.qrq/toplist"
)

func main() {
	rand.Seed(time.Now().Unix())
	callsigns := readCallsigns()
	cx := getRandomCallsign(callsigns)
	cxLocation := "NEW WESTMINSTER BC CANADA"
	cxName := "HUGH"
	rx := getRandomCallsign(callsigns)
	rxLocation := "DALLAS TX USA"
	rxName := "JANE"
	fmt.Printf(initialGreeting(cx, rx))
	// printMorse(initialGreeting(cx, rx))
	fmt.Println(firstExchange(rx, cx, cxName, cxLocation))
	fmt.Println(secondExchange(rx, cx, rxName, rxLocation))
}

func firstExchange(rx, cx, cxName, cxLocation string) string {
	msg := fmt.Sprintf("%s DE %s ", rx, cx)
	msg = msg + "TNX FOR CALL BT UR RST 599 599 HR "
	msg = msg + qth(cxLocation) + " "
	msg = msg + name(cxName) + " "
	msg = msg + "HW CPY? "
	msg = msg + kn(rx, cx)
	return msg
}

func name(name string) string {
	return fmt.Sprintf("NAME %s %s", name, name)
}

func qth(location string) string {
	return fmt.Sprintf("QTH %s %s", location, location)
}

func kn(cx, rx string) string {
	return fmt.Sprintf("%s DE %s", cx, rx)
}

func secondExchange(rx, cx, rxName, rxLocation string) string {
	msg := fmt.Sprintf("%s DE %s ", cx, rx)
	msg = msg + "TNX FOR RPT SLD CPY FB UR RST 599 599 BT"
	msg = msg + name(rxName) + " "
	msg = msg + qth(rxLocation) + " "
	msg = msg + kn(rx, cx)
	return msg
}

func printMorse(msg string) {
	fmt.Println(morse.EncodeITU(msg))
}

func initialGreeting(cx, rx string) string {
	cxRepeat := fmt.Sprintf("%s %s %s", cx, cx, cx)
	rxRepeat := fmt.Sprintf("%s %s %s", rx, rx, rx)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", cxRepeat, cx, rxRepeat)
	return msg
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

func getRandomCallsign(cs *[]string) string {
	return (*cs)[rand.Intn(len(*cs))]
}
