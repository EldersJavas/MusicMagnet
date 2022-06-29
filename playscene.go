// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type PlayScene struct {
	audioContext *audio.Context
	PlayMusic    *audio.Player
	ClickMusic   *audio.Player
}

func (s *PlayScene) Update(sceneSwitcher SceneSwitcher) error {
	//TODO implement me
	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	sw, sh := InitSceneScreen(screen)

	lines := []string{
		Tr("tt"),
	}
	for _, line := range lines {
		f := spaceAgeBig
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 150 + sh - sh
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}
}
