package main

import (
	"bufio"
	"fmt"
	"math"
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

	if err := beep.OpenSoundDevice("default"); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}

	if err := beep.InitSoundDevice(); err != nil {
		fmt.Printf("Can't open sound device: %s\n", err.Error())
	}
	defer beep.CloseSoundDevice()
	fmt.Printf(initialGreeting(cx, rx))
	playMorse("VA7UNX")
	// printMorse(initialGreeting(cx, rx))
	fmt.Println(firstExchange(cx, rx))
	playMorse(firstExchange(cx, rx))
	fmt.Println(secondExchange(cx, rx))
	playMorse(secondExchange(cx, rx))
	fmt.Println(gnightBob(cx, rx))
	playMorse(gnightBob(cx, rx))
}

func playMorse(s string) {
	fmt.Println("[FIXME] About to encode " + s)
	stuff := morse.EncodeITU(strings.ToLower(s))
	fmt.Println("[FIXME] " + stuff)
	for _, letter := range strings.Split(stuff, "") {
		if letter == "-" {
			dah()
		} else if letter == "." {
			dit()
		} else if letter == " " {
			time.Sleep(time.Duration(200 * time.Millisecond))
		}
	}
}

func dit() {
	doABeep(150)
}

func dah() {
	doABeep(300)
}

func doABeep(duration int) {
	freqHertz := 500.0
	vol := 80
	var foo string

	freq := beep.HertzToFreq(freqHertz)
	music := beep.NewMusic(foo)
	playBeep(music, vol, duration, 1, freq)
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

// Taken from github.com/dbatbold/beep; 2-term BSD license
// Thanks, dbatbold!
func playBeep(music *beep.Music, volume, duration, count int, freq float64) {
	bar := beep.SampleAmp16bit * (float64(volume) / 100.0)
	samples := int(beep.SampleRate64 * (float64(duration) / 1000.0))
	rest := 0
	if count > 1 {
		rest = (beep.SampleRate / 20) * 4 // 200ms
	}
	buf := make([]int16, samples+rest)
	var last int16
	var fade = 1024
	if samples < fade {
		fade = 1
	}
	for i := range buf {
		if i < samples-fade {
			buf[i] = int16(bar * math.Sin(float64(i)*freq))
			last = buf[i]
		} else {
			if last > 0 {
				last -= 31
			} else {
				last += 31
			}
			buf[i] = last
		}
	}
	beep.InitSoundDevice()
	for i := 0; i < count; i++ {
		go music.Playback(buf, buf)
		music.WaitLine()
	}
	beep.FlushSoundBuffer()
}
