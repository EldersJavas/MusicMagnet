// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Hajime Hoshi

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"os"
	"sync"
	"time"
)

type SceneSwitcher interface {
	ChoScene()
	MainScene()
	PlayScene(n int)
	LoadScene()
	AboutScene()
}

type Scene interface {
	Update(sceneSwitcher SceneSwitcher) error
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene     Scene
	nextScene Scene
	Mplay     GPlayer
	Once      sync.Once
}

func (g *Game) Update() error {
	if g.nextScene != nil {
		g.scene = g.nextScene
		g.nextScene = nil
	}
	if err := g.scene.Update(g); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func (g *Game) MainScene() {
	g.Once.Do(func() { go g.Mplay.BGMBegain() })

	time.Sleep(time.Millisecond * 100)
	g.nextScene = &MainScene{}
}
func (g *Game) AboutScene() {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &AboutScene{}
}
func (g *Game) PlayScene(n int) {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &PlayScene{MuN: n}
}
func (g *Game) ChoScene() {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &ChooseScene{}
}
func (g *Game) LoadScene() {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &LoadingScene{}
}

//🎶♫♬♪♩¶♯♮♭🎵🎼🎶🧲⏱
func main() {
	ebiten.SetWindowSize(1024, 576)
	ebiten.SetWindowTitle("♪Music Magnet♪")
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	for _, arg := range os.Args {
		switch arg {
		case "-f":
			ebiten.SetFullscreen(true)
		case "-en":
			Langs = EN
		case "-zh":
			Langs = ZH
		}
	}
	//fmt.Printf("%v,%v,%v", os.Args, Langs, Tr("code"))
	g := &Game{
		scene: &SplashScene{},
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type GPlayer struct {
	audioContext *audio.Context
	bgmPlayer    *audio.Player
}

func (g *GPlayer) BGMStateChange() {
	if g.bgmPlayer.IsPlaying() {
		g.bgmPlayer.Pause()
	} else {
		g.bgmPlayer.Play()
	}
}
func (g *GPlayer) BGMBegain() error {
	const sampleRate = 44100
	g.audioContext = audio.NewContext(sampleRate)
	{
		f, err := ResourceFS.Open("bgm.ogg")
		if err != nil {
			return err
		}
		defer f.Close()

		decoded, err := vorbis.DecodeWithSampleRate(sampleRate, f)
		if err != nil {
			return err
		}
		loop := audio.NewInfiniteLoop(decoded, decoded.Length())
		p, err := g.audioContext.NewPlayer(loop)
		if err != nil {
			return err
		}
		g.bgmPlayer = p
		g.bgmPlayer.SetVolume(0.8)
		g.bgmPlayer.Play()
	}
	return nil
}
