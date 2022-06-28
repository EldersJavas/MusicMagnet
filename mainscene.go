// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type MainScene struct {
	Title         string
	Playbutton    string
	audioContext  *audio.Context
	bgmPlayer     *audio.Player
	seStartPlayer *audio.Player
	seEndPlayer   *audio.Player
}

func NewMainScene() *MainScene {
	return &MainScene{Title: "Music Magnet", Playbutton: "play"}
}

func (s *MainScene) Draw(screen *ebiten.Image) {

}
