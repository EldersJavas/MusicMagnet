// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"math/rand"
	"sync"
	"time"
)

type PlayScene struct {
	MuN    int
	MuC    Music
	Start  bool
	PointC int
	//Player PPlayer
	Once1  sync.Once
	MX, MY int
	Load   bool
	B      []Ball
}
type PPlayer struct {
	audioContext *audio.Context
	Music        *audio.Player
}

type Ball struct {
	State int
	X, Y  int
	Im    *ebiten.Image
}

const (
	W = iota
	A
	S
	D
)

func (g *PlayScene) AddBall(bpm int) {
	for true {
		time.Sleep(time.Second * 3)
		r := rand.Intn(4)
		p := ebiten.NewImage(32, 32)
		p.Fill(color.RGBA{1, 1, 1, 255})
		Xa, Ya := 0, 0
		switch r {
		case A:
			Xa = 500
			Ya = 500
		case D:
			Xa = 500
			Ya = 500
		case S:
			Xa = 500
			Ya = 500
		case W:
			Xa = 500
			Ya = 500
		}

		b := Ball{State: r, Im: p, X: Xa, Y: Ya}
		g.B = append(g.B, b)

	}

}

func (s *PlayScene) Update(sceneSwitcher SceneSwitcher) error {
	go func() {
		s.AddBall(120)
	}()
	for i, ball := range s.B {
		switch ball.State {
		case A:
			ball.X -= 30
		case D:
			ball.X += 30
		case S:
			ball.Y += 30
		case W:
			ball.Y -= 30
		}

		if ball.X <= 10 || ball.X >= 1800 || ball.Y <= 10 || ball.Y >= 1000 {
			s.B = append(s.B[:i], s.B[i+1:]...)
		}
	}

	s.MX, s.MY = ebiten.CursorPosition()
	if s.Start == false {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			if IsInPos(float64(s.MX), float64(s.MY), 800, 310, 1100, 470) {
				Mplay.MusicPlay.Play()
				s.Start = true
			}
		}

	}
	if s.Start == true {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			if IsInPos(float64(s.MX), float64(s.MY), 1700, 0, 1920, 120) {
				Mplay.MusicPlay.Pause()
				s.Start = false
			}
		}
	}

	if s.Load == false {
		const sampleRate = 44100
		switch s.MuN {
		case 0:
			s.MuC = MusicMap["BadApple"]
			LogP(s.MuC)
			decoded, err := wav.DecodeWithSampleRate(sampleRate, LoadMusic("Bad_Apple!!.wav"))
			LogP(decoded, err)
			loop := audio.NewInfiniteLoop(decoded, decoded.Length())
			p, err := Mplay.AudioContext.NewPlayer(loop)
			LogP(p, err)
			Mplay.MusicPlay = p
			Mplay.MusicPlay.Pause()
		case 1:
			s.MuC = MusicMap["AnDieFreude"]
			LogP(s.MuC)
			decoded, err := wav.DecodeWithSampleRate(sampleRate, LoadMusic("AnDieFreude.wav"))
			LogP(decoded)
			loop := audio.NewInfiniteLoop(decoded, decoded.Length())
			p, err := Mplay.AudioContext.NewPlayer(loop)
			LogP(err)
			Mplay.MusicPlay = p
			Mplay.MusicPlay.Pause()
		}
		Mplay.bgmPlayer.Pause()
		go func() {
			for true {
				if s.Start {
					time.Sleep(time.Second)
					s.P(1)
				}
			}
		}()

		s.Load = true
	}

	return nil
}

func (s *PlayScene) Draw(screen *ebiten.Image) {

	sw, sh := InitSceneScreen(screen)
	for _, ball := range s.B {
		op1 := &ebiten.DrawImageOptions{}
		op1.GeoM.Translate(float64(ball.X), float64(ball.Y))

		screen.DrawImage(ball.Im, op1)

	}

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
		r := text.BoundString(f, line)
		x := sw - r.Max.X
		y := 100
		text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
	}
	if !s.Start {
		lines = []string{
			Tr("Start"),
		}
		for _, line := range lines {
			f := spaceAgeBig
			r := text.BoundString(f, line)
			x := (sw-r.Dx())/2 - r.Min.X
			y := sh/2 - 100
			text.Draw(screen, line, f, x, y, color.RGBA{236, 109, 136, 0xff})
		}
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

func (g *PlayScene) P(a int) {
	g.PointC = g.PointC + a
}
