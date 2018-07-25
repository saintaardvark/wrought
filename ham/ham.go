package ham

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	callsignFile = "/home/aardvark/.qrq/toplist"
	citiesFile   = "./data/world-cities/data/world-cities.csv"
)

// Ham is a ham
type Ham struct {
	Callsign string
	Location string
	Name     string
}

// NewHam returns a pointer to a newly generated Ham struct
func NewHam() *Ham {
	ham := Ham{
		Callsign: getRandomCallsign(),
		Location: getRandomCity(),
		Name:     "JANE",
	}
	return &ham
}

func getRandomCallsign() string {
	cs := readCallsigns()
	return (*cs)[rand.Intn(len(*cs))]
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
