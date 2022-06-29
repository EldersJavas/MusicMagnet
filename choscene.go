// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type ChooseScene struct {
	MX int
	MY int
}

func (s *ChooseScene) Update(sceneSwitcher SceneSwitcher) error {
	s.MX, s.MY = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if IsInPos(float64(s.MX), float64(s.MY), 20, 930, 230, 1040) {
			go sceneSwitcher.MainScene()
		}
		if IsInPos(float64(s.MX), float64(s.MY), 550, 260, 1290, 350) {
			go sceneSwitcher.PlayScene(0)
		}
		if IsInPos(float64(s.MX), float64(s.MY), 630, 420, 1230, 520) {
			go sceneSwitcher.PlayScene(1)
		}
		//if IsInPos(float64(s.MX), float64(s.MY), 20, 930, 230, 1040) {go sceneSwitcher.MainScene()}

		//if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){}
	}
	return nil
}

func (s *ChooseScene) Draw(screen *ebiten.Image) {
	sw, sh := InitSceneScreen(screen)

	lines := []string{Tr("ChMu")}
	for _, line := range lines {
		f := spaceAgeBig
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 150
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}

	lines = []string{
		"1. Bad Apple!!",
		"2. An Die Freude",
	}
	for n, line := range lines {
		f := spaceAgeSmall
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 350 + (n * 150)
		text.Draw(screen, line, f, x, y, color.CMYK{61, 49, 0, 0})
	}

	lines = []string{Tr("Back")}
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
