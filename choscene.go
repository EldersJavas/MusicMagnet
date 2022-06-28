// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type ChooseScene struct {
	audioContext  *audio.Context
	bgmPlayer     *audio.Player
	seStartPlayer *audio.Player
	seEndPlayer   *audio.Player
	MX            int
	MY            int
}

func (s *ChooseScene) Update(sceneSwitcher SceneSwitcher) error {
	s.MX, s.MY = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if IsInPos(float64(s.MX), float64(s.MY), 20, 930, 230, 1040) {
			go sceneSwitcher.MainScene()
		}

	}
	//TODO implement me
	return nil
}

func (s *ChooseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})
	sw, sh := screen.Size()
	lines := []string{
		Tr("Back"),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 30
		y := sh - 60
		text.Draw(screen, line, f, x, y, color.RGBA{236, 228, 197, 0xff})
	}

	//Pos debug
	lines = []string{fmt.Sprintf("x:%v,Y:%v", s.MX, s.MY)}
	for _, line := range lines {
		f := DebugFace
		r := text.BoundString(f, line)
		x := sw - r.Max.X - 100
		y := sh - 60
		text.Draw(screen, line, f, x, y, color.RGBA{1, 1, 1, 0xff})
	}
}
