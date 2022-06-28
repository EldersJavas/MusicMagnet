// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"image/color"
)

type ChooseScene struct {
	audioContext  *audio.Context
	bgmPlayer     *audio.Player
	seStartPlayer *audio.Player
	seEndPlayer   *audio.Player
}

func (s *ChooseScene) Update(sceneSwitcher SceneSwitcher) error {
	//TODO implement me
	return nil
}

func (s *ChooseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})

}
