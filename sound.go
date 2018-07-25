package main

import (
	"fmt"
	"math"

	"github.com/dbatbold/beep"
)

func doABeep(duration int) {
	freqHertz := 500.0
	vol := 80

	freq := beep.HertzToFreq(freqHertz)
	// Just send empty string
	music := beep.NewMusic("")
	playBeep(music, vol, duration, 1, freq)
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
	// fmt.Printf("[FIXME] buf: %+v\n", buf)
	beep.InitSoundDevice()
	for i := 0; i < count; i++ {
		go music.Playback(buf, buf)
		music.WaitLine()
	}
	beep.FlushSoundBuffer()
}

func buildABeep(volume, duration, count int, freq float64) *[]int16 {
	fmt.Printf("[FIXME] buildABeep: vol: %d\n", volume)
	fmt.Printf("[FIXME] buildABeep: duration: %d\n", duration)
	fmt.Printf("[FIXME] buildABeep: freq: %f\n", freq)
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
	// fmt.Printf("[FIXME] buildABeep: buf: %+v\n", buf)
	return &buf
}

func justPlayBeep(music *beep.Music, buf *[]int16) {
	fmt.Println("[FIXME] Made it to justPlayBeep")
	// fmt.Printf("[FIXME] *buf: %+v\n", *buf)
	// assigning buf directly does not affect anything
	// buf2 := *buf
	beep.InitSoundDevice()
	go music.Playback(*buf, *buf)
	music.WaitLine()
	beep.FlushSoundBuffer()
}
