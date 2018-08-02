package main

import (
	"fmt"
	"strings"

	"github.com/dbatbold/beep"
	"github.com/martinlindhe/morse"
)

const (
	ditLength     = 150
	dahLength     = 300
	letterPause   = 3 * ditLength
	wordPause     = 7 * ditLength
	sentencePause = 10 * ditLength
	freq          = 500
	volume        = 80
)

var (
	prosigns = map[string]bool{
		"CQ": true,
		"KN": true,
		"BT": true,
	}
)

type morsePlayer struct {
	music     *beep.Music
	exchange  []string
	freqHertz float64
	vol       int
	samples   []*[]int16
}

func newMorsePlayer() *morsePlayer {
	player := morsePlayer{
		music:     beep.NewMusic(""),
		exchange:  []string{},
		freqHertz: beep.HertzToFreq(freq),
		vol:       volume,
		samples:   []*[]int16{},
	}
	return &player
}

func (player *morsePlayer) Print() {
	fmt.Printf("Player: freq: %f, vol: %d\n", player.freqHertz, player.vol)
}

func (player *morsePlayer) PrintCW() {
	for _, s := range player.exchange {
		fmt.Println(morse.EncodeITU(strings.ToLower(s)))
	}
}

func (player *morsePlayer) CW() string {
	var cw string
	for _, s := range player.exchange {
		cw = fmt.Sprintf("%s\n%s", cw, morse.EncodeITU(strings.ToLower(s)))
	}
	return cw
}

func (player *morsePlayer) PrintText() {
	for _, s := range player.exchange {
		fmt.Println(s)
	}
}

func (player *morsePlayer) PlayCW() {
	for _, exch := range player.exchange {
		player.buildCWSamplesRecursive(exch)
		player.buildSentencePause()
	}
	for _, sample := range player.samples {
		justPlayBeep(player.music, sample)
	}
}

func (player *morsePlayer) PlayRemoteHalf() {
	return
}

// Rewrite this to be recursive
func (player *morsePlayer) buildCWSamplesRecursive(s string) {
	if strings.Contains(s, " ") {
		for _, w := range strings.Split(s, " ") {
			player.buildCWSamplesRecursive(w)
		}
		return
	}
	if prosigns[s] == true {
		player.buildProsign(s)
	} else {
		player.buildWord(s)
	}
	player.buildWordPause()
}

func (player *morsePlayer) buildWord(word string) {
	m := morse.EncodeITU(strings.ToLower(word))
	for _, s := range strings.Split(m, "") {
		if s == "-" {
			player.buildDah()
		} else if s == "." {
			player.buildDit()
		} else {
			player.buildLetterPause()
		}
	}
	player.buildWordPause()
}

func (player *morsePlayer) buildProsign(prosign string) {
	m := morse.EncodeITU(strings.ToLower(prosign))
	for _, s := range strings.Split(m, "") {
		if s == "-" {
			player.buildDah()
		} else if s == "." {
			player.buildDit()
		}
	}
}

func (player *morsePlayer) buildDit() {
	newSamples := buildABeep(player.vol, ditLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildDah() {
	newSamples := buildABeep(player.vol, dahLength, 1, player.freqHertz)
	player.samples = append(player.samples, newSamples)
}

func (player *morsePlayer) buildWordPause() {
	player.buildPause(wordPause)
}

func (player *morsePlayer) buildSentencePause() {
	player.buildPause(sentencePause)
}

func (player *morsePlayer) buildLetterPause() {
	player.buildPause(letterPause)
}

func (player *morsePlayer) buildPause(pause int) {
	var newSamples *[]int16
	newSamples = buildABeep(0, pause, 1, 0.0)
	player.samples = append(player.samples, newSamples)
}

func doABeep(duration int) {
	// Just send empty string
	music := beep.NewMusic("")
	playBeep(music, volume, duration, 1, beep.HertzToFreq(freq))
}

func dit() {
	doABeep(150)
}

func dah() {
	doABeep(300)
}

// Taken from github.com/dbatbold/beep; 2-term BSD license
// Thanks, dbatbold!
func playBeep(music *beep.Music, volume, duration, count int, freq float64) {
	buf := buildABeep(volume, duration, count, freq)
	justPlayBeep(music, buf)
}

func buildABeep(volume, duration, count int, freq float64) *[]int16 {
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
	return &buf
}

func justPlayBeep(music *beep.Music, buf *[]int16) {
	beep.InitSoundDevice()
	go music.Playback(*buf, *buf)
	music.WaitLine()
	beep.FlushSoundBuffer()
}
