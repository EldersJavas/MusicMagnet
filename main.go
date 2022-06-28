// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Hajime Hoshi

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

type SceneSwitcher interface {
	SwitchToGameScene()
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

func (g *Game) SwitchToGameScene() {
	g.nextScene = &GameScene{}
}

//🎶♫♬♪♩¶♯♮♭🎵🎼🎶🧲⏱
func main() {
	ebiten.SetWindowSize(1024, 576)
	ebiten.SetWindowTitle("♪Music Magnet♪")
	ebiten.SetMaxTPS(120)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if os.Args[0] == "-f" {
		ebiten.SetFullscreen(true)
	}
	fmt.Sprintln(os.Args)
	g := &Game{
		scene: &SplashScene{},
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
