// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"io"
)

type Music struct {
	Name      string
	TestVoice []byte
	Voice     []byte
	Img       *ebiten.Image
	//Pu        []rune
	BPM  int
	Time int
	//LRC
}

func init() {
	const sampleRate = 48000
	audioContext = audio.NewContext(sampleRate)
	m, err := resourceFS.Open("AnDieFreude.wav")
	if err != nil {
		panic(err)
	}
	defer m.Close()

	bs, err := io.ReadAll(m)
	if err != nil {
		panic(err)
	}
	decoded, err := 
	if err != nil {
		return err
	}



}

var MusicMap = map[string]Music{
	"BadApple": {
		Name:      "Bad Apple!!",
		TestVoice: nil,
		Voice:     nil,
		Img:       nil,
		BPM:       140,
		Time:      319,
	},
	"AnDieFreude": {
		Name:      "An Die Freude",
		TestVoice: nil,
		Voice:     nil,
		Img:       nil,
		BPM:       125,
		Time:      35,
	},
}
