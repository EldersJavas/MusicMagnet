// Created by EldersJavas(EldersJavas&gmail.com)

package main

func IsInPos(x, y, x1, y1, x2, y2 float64) bool {
	if x >= x1 && x <= x2 {
		if y >= y1 && y <= y2 {
			return true
		}
	}
	return false
}
