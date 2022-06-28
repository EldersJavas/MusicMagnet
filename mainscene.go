// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"math/rand"
)

type MainScene struct {
	audioContext  *audio.Context
	bgmPlayer     *audio.Player
	seStartPlayer *audio.Player
	seEndPlayer   *audio.Player
}

func (s *MainScene) Update(sceneSwitcher SceneSwitcher) error {
	//TODO implement me
	return nil
}

func (s *MainScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{247, 201, 216, 0xff})
	sw, sh := screen.Size()
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
		Tr("Ab"),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 30
		y := sh - 100
		text.Draw(screen, line, f, x, y, color.RGBA{236, 228, 197, 0xff})
	}
}
