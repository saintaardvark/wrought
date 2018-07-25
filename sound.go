package main

import (
	"math"

	"github.com/dbatbold/beep"
)

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
