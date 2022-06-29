// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"io"
)

type Music struct {
	Name      string
	TestVoice io.Reader
	Voice     io.Reader
	Img       *ebiten.Image
	//Pu        []rune
	BPM  int
	Time int
	//LRC
}

func LoadMusic(file string) io.Reader {
	const sampleRate = 48000
	m, err := resourceFS.Open(file)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	return m
}

var MusicMap = map[string]*Music{
	"BadApple": {
		Name:      "Bad Apple!!",
		TestVoice: LoadMusic("Bad_Apple!!.mp3"),
		Voice:     LoadMusic("Bad_Apple!!.mp3"),
		Img:       nil,
		BPM:       140,
		Time:      319,
	},
	"AnDieFreude": {
		Name:      "An Die Freude",
		TestVoice: LoadMusic("AnDieFreude.mp3"),
		Voice:     LoadMusic("AnDieFreude.mp3"),
		Img:       nil,
		BPM:       125,
		Time:      35,
	},
}
