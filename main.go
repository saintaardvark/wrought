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
	fmt.Printf("%s\n", morse.EncodeITU(initialGreeting(cx, rx)))
	fmt.Println(rx, "DE", cx, "TNX FOR CALL BT UR RST 599 599 HR QTH", cxLocation, cxLocation, "NAME", cxName, cxName, "HW CPY?", rx, "DE", cx, "KN")
	fmt.Println(cx, "DE", rx, "TNX FOR RPT SLD CPY FB UR RST 599 599 BT NAME", rxName, rxName, "QTH", rxLocation, rxLocation)
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
