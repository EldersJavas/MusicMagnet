// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type AboutScene struct {
	audioContext  *audio.Context
	bgmPlayer     *audio.Player
	seStartPlayer *audio.Player
	seEndPlayer   *audio.Player
}

func (s *AboutScene) Update(sceneSwitcher SceneSwitcher) error {
	//TODO implement me
	return nil
}

func (s *AboutScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})

	lines := []string{
		Tr("Ab"),
		Tr("Cp"),
		Tr("Ab"),
		Tr("Ab"),
		Tr("Ab"),
		Tr("Ab"),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 20
		y := 20
		text.Draw(screen, line, f, x, y, color.RGBA{172, 139, 191, 0xff})
	}
}
