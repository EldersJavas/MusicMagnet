// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"sync"
	"time"
)

type PlayScene struct {
	MuN          int
	MuC          Music
	Start        bool
	PointC       int
	audioContext *audio.Context
	PlayMusic    *audio.Player
	ClickMusic   *audio.Player
	Once1        sync.Once
	MX, MY       int
}

func (s *PlayScene) Update(sceneSwitcher SceneSwitcher) error {
	s.MX, s.MY = ebiten.CursorPosition()
	if s.Start != true {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			if IsInPos(float64(s.MX), float64(s.MY), 1700, 0, 1920, 120) {
				s.PlayMusic.Play()
				s.Start = true
			}
		}
	}

	s.Once1.Do(func() {
		const sampleRate = 44100
		switch s.MuN {
		case 0:
			s.MuC = MusicMap["BadApple"]
			LogP(s.MuC)
			decoded, err := mp3.DecodeWithSampleRate(sampleRate, s.MuC.Voice)
			LogP(decoded, err)
			loop := audio.NewInfiniteLoop(decoded, decoded.Length())
			p, err := s.audioContext.NewPlayer(loop)
			LogP(p, err)
			s.PlayMusic = p
			s.PlayMusic.Pause()
		case 1:
			s.MuC = MusicMap["AnDieFreude"]
			LogP(s.MuC)
			decoded, err := wav.DecodeWithSampleRate(sampleRate, s.MuC.Voice)
			LogP(decoded)
			loop := audio.NewInfiniteLoop(decoded, decoded.Length())
			p, err := s.audioContext.NewPlayer(loop)
			LogP(err)
			s.PlayMusic = p
			s.PlayMusic.Pause()
		}
		for true {
			if s.Start {
				time.Sleep(time.Second)
				s.P(1)
			}
		}
	})

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
