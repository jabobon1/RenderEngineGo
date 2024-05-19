package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var FPS uint32 = 60
var SPEED uint32 = 5
var MOVE_SPEED float64 = 10

func updateFPS(up bool) {
	if up {
		FPS += SPEED
	} else if FPS >= SPEED {
		FPS -= SPEED
	}
	if FPS > 240 {
		FPS = 240 // Cap the FPS at 240
	} else if FPS <= 0 {
		FPS = 1
	}
	fmt.Printf("FPS increased to %d\n", FPS)
}

type MyGameEngine struct {
	GameEngine
}

func (e MyGameEngine) handleKeyBoardPress(event sdl.Event) bool {
	cube := e.getGameObject()

	switch t := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.K_UP && t.State == sdl.PRESSED {
			updateFPS(true)
		} else if t.Keysym.Sym == sdl.K_DOWN && t.State == sdl.PRESSED {
			updateFPS(false)
		} else if t.Keysym.Sym == sdl.K_SPACE && t.State == sdl.PRESSED {
			cube.angles.changeAxe()
		} else if t.Keysym.Sym == sdl.K_LEFT && t.State == sdl.PRESSED {
			cube.angles.changeAngleVelociity(false)
		} else if t.Keysym.Sym == sdl.K_RIGHT && t.State == sdl.PRESSED {
			cube.angles.changeAngleVelociity(true)
		} else if t.Keysym.Sym == sdl.K_a && t.State == sdl.PRESSED {
			cube.position.X -= MOVE_SPEED
		} else if t.Keysym.Sym == sdl.K_d && t.State == sdl.PRESSED {
			cube.position.X += MOVE_SPEED
		} else if t.Keysym.Sym == sdl.K_w && t.State == sdl.PRESSED {
			cube.position.Y -= MOVE_SPEED
		} else if t.Keysym.Sym == sdl.K_s && t.State == sdl.PRESSED {
			cube.position.Y += MOVE_SPEED
		}
	}
	return false
}

func (e MyGameEngine) update() {
	// Переопределенная реализация метода update()
	var cube *GameObject3D = e.getGameObject()

	cube.rotate()

	e.renderer.SetDrawColor(WHITE.R, WHITE.G, WHITE.B, WHITE.A)
	e.renderer.Clear()

	e.renderer.SetDrawColor(BLACK.R, BLACK.G, BLACK.B, BLACK.A)
	cube.draw(e.renderer)
	err := drawUI(e.renderer, cube.angles)
	if err != nil {
		fmt.Println("Error drawing UI:", err)
		return
	}
	e.renderer.Present()
}

func main() {
	gameEngine, err := initGameEngine(nil, 1500, 1200)
	if err != nil {
		fmt.Println("Error creating GameEngine:", err)
		return
	}
	defer gameEngine.Close()

	myEngine := MyGameEngine{*gameEngine}
	cube := getCube3D(300)
	myEngine.addGameObj(cube)
	fmt.Println("myEngine", myEngine.gameObjects)
	run(myEngine, FPS)
}
