// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Hajime Hoshi

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
	"time"
)

type SceneSwitcher interface {
	ChoScene()
	MainScene()
	PlayScene()
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
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &MainScene{}
}
func (g *Game) AboutScene() {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &AboutScene{}
}
func (g *Game) PlayScene() {
	time.Sleep(time.Millisecond * 100)
	g.nextScene = &PlayScene{}
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
	ebiten.SetMaxTPS(120)
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
	fmt.Printf("%v,%v,%v", os.Args, Langs, Tr("code"))
	g := &Game{
		scene: &SplashScene{},
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
