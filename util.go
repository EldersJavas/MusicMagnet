// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

func IsInPos(x, y, x1, y1, x2, y2 float64) bool {
	if x >= x1 && x <= x2 {
		if y >= y1 && y <= y2 {
			return true
		}
	}
	return false
}

func InitSceneScreen(image *ebiten.Image) (int, int) {
	image.Fill(color.RGBA{247, 201, 216, 0xff})
	return image.Size()
}
func LogP(a ...any) {
	log.Println(a)
}
