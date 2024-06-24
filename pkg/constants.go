package pkg

import "github.com/veandco/go-sdl2/sdl"

var WHITE = sdl.Color{R: 255, G: 255, B: 255, A: 255}
var BLACK = sdl.Color{R: 0, G: 0, B: 0, A: 255}

const (
	xAxe int = 0
	yAxe int = 1
	zAxe int = 2
)

var axes = [3]int{xAxe, yAxe, zAxe}
