// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type AboutScene struct {
	MX int
	MY int
}

func (s *AboutScene) Update(sceneSwitcher SceneSwitcher) error {
	s.MX, s.MY = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		if IsInPos(float64(s.MX), float64(s.MY), 20, 930, 230, 1040) {
			go sceneSwitcher.MainScene()
		}
	}
	return nil
}

func (s *AboutScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})
	sw, sh := screen.Size()
	lines := []string{
		Tr("Ab"),
		Tr("tt"),
		Tr("Web"),
		Tr("Lic"),
		Tr("THX"),
		Tr("Cp"),
	}
	for n, line := range lines {
		f := spaceAgeSmall
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X
		y := 100 + n*180
		text.Draw(screen, line, f, x, y, color.RGBA{1, 1, 1, 0xff})
	}

	lines = []string{
		Tr("Back"),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 30
		y := sh - 60
		text.Draw(screen, line, f, x, y, color.RGBA{236, 228, 197, 0xff})
	}
}
