// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"sync"
	"time"
)

type PlayScene struct {
	Start        bool
	PointC       int
	audioContext *audio.Context
	PlayMusic    *audio.Player
	ClickMusic   *audio.Player
	Once1        sync.Once
}

func (s *PlayScene) Update(sceneSwitcher SceneSwitcher) error {
	s.Once1.Do(func() {
		if s.Start {
			time.Sleep(time.Second)
			s.P(1)
		}
	})
	//TODO implement me
	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	sw, sh := InitSceneScreen(screen)
	sw = sw
	sh = sh
	lines := []string{
		fmt.Sprintf(Tr("Point"), s.PointC),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 20
		y := 100
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}
	lines = []string{
		Tr("Pause"),
	}
	for _, line := range lines {
		f := spaceAgeSmall
		//r := text.BoundString(f, line)
		x := 20
		y := 100
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}

}

func (g *PlayScene) P(a int) {
	g.PointC = g.PointC + a
}
