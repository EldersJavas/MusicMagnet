// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"math/rand"
)

type MainScene struct {
	MX int
	MY int
}

func (s *MainScene) Update(sceneSwitcher SceneSwitcher) error {
	s.MX, s.MY = ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		if IsInPos(float64(s.MX), float64(s.MY), 20, 930, 230, 1040) {
			go sceneSwitcher.AboutScene()
		}
		if IsInPos(float64(s.MX), float64(s.MY), 670, 720, 1250, 870) {
			go sceneSwitcher.ChoScene()
		}

	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if IsInPos(float64(s.MX), float64(s.MY), 1700, 0, 1920, 120) {
			ChangeLang()
		}
	}

	return nil
}

func (s *MainScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})
	sw, sh := screen.Size()
	//
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(float64(sw/2-rand.Intn(300)+rand.Intn(300)), float64(sh/2-rand.Intn(300)+rand.Intn(300)))
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(sw/2-rand.Intn(300)+rand.Intn(300)), float64(sh/2-rand.Intn(300)+rand.Intn(300)))

	screen.DrawImage(emoji.Image("🎵"), op1)
	screen.DrawImage(emoji.Image("🧲"), op2)

	lines := []string{
		Tr("Play"),
	}
	for _, line := range lines {
		f := spaceAgeBig
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X
		y := sh - 144 - 96
		text.Draw(screen, line, f, x, y, color.RGBA{172, 139, 191, 0xff})
	}
	lines = []string{
		Tr("tt"),
	}
	for _, line := range lines {
		f := spaceAgeBig
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 150
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}
	lines = []string{
		GetLang(),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		r := text.BoundString(f, line)
		x := sw - r.Max.X
		y := 100
		text.Draw(screen, line, f, x, y, color.RGBA{114, 139, 198, 0xff})

		// About 20,900-300,980
		lines = []string{
			Tr("Ab"),
		}
		for _, line := range lines {
			f := spaceAgeSmall
			//r := text.BoundString(f, line)
			x := 30
			y := sh - 60
			text.Draw(screen, line, f, x, y, color.RGBA{236, 228, 197, 0xff})

			//Pos debug
			lines = []string{fmt.Sprintf("x:%v,Y:%v", s.MX, s.MY)}
			for _, line := range lines {
				f := DebugFace
				r := text.BoundString(f, line)
				x := sw - r.Max.X - 100
				y := sh - 60
				text.Draw(screen, line, f, x, y, color.RGBA{1, 1, 1, 0xff})
			}
			//ebitenutil.DebugPrint(screen, )

		}
	}
}
