// Created by EldersJavas(EldersJavas&gmail.com)

package main

const (
	EN = iota
	ZH
)

var Langs = EN
var I18nEN = map[string]string{
	"code":  "en",
	"tt":    "♪Music Magnet♪",
	"Play":  "Play!",
	"ChMu":  "Choose Music",
	"Ab":    "About",
	"Load":  "Loading",
	"THX":   "Spcial Thank Hajime Hoshi",
	"Great": "Great",
	"Miss":  "Missed",
	"Point": "Point",
	"Lic":   "Apache 2.0",
	"Cp":    "EbitenPot 2021-2022",
	"Web":   "github.com/EldersJavas/MusicMagnet",
	"Back":  "Back",
}
var I18nZH = map[string]string{
	"code":  "zh",
	"tt":    "♪音乐磁力♪",
	"Play":  "开始游戏",
	"ChMu":  "选择音乐",
	"Ab":    "关于",
	"Load":  "加载",
	"THX":   "鸣谢 Hajime Hoshi",
	"Great": "好",
	"Miss":  "失误",
	"Point": "得分",
	"Lic":   "Apache 2.0",
	"Cp":    "EbitenPot 2021-2022",
	"Web":   "github.com/EldersJavas/MusicMagnet",
	"Back":  "返回",
}

func Tr(str string) string {
	switch Langs {
	case EN:
		return I18nEN[str]
	case ZH:
		return I18nZH[str]
	}
	return "No String"
}

func GetLang() string {
	switch Langs {
	case EN:
		return "EN"
	case ZH:
		return "ZH"
	}
	return ""
}
func ChangeLang() {
	if Langs < 1 {
		Langs++
	} else {
		Langs = 0
	}

}
