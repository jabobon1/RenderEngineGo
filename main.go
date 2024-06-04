package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var FPS uint32 = 35
var SPEED uint32 = 5
var MOVE_SPEED float64 = 0.1

const (
	WIDTH      int32   = 2000
	HEIGHT     int32   = 1500
	FOW        float64 = 45.0 // Vertical field of view
	nearPlane          = 0.1
	farPlane           = 100.0
	zThreshold         = 0.001
)

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
			e.camera.changePosition(Vector3D{MOVE_SPEED, 0, 0})
		} else if t.Keysym.Sym == sdl.K_d && t.State == sdl.PRESSED {
			e.camera.changePosition(Vector3D{-MOVE_SPEED, 0, 0})
		} else if t.Keysym.Sym == sdl.K_w && t.State == sdl.PRESSED {
			e.camera.changePosition(Vector3D{0, -MOVE_SPEED, 0})
		} else if t.Keysym.Sym == sdl.K_s && t.State == sdl.PRESSED {
			e.camera.changePosition(Vector3D{0, MOVE_SPEED, 0})
		} else if t.Keysym.Sym == sdl.K_KP_PLUS && t.State == sdl.PRESSED {
			e.camera.changePosition(Vector3D{0, 0, -MOVE_SPEED})
		} else if t.Keysym.Sym == sdl.K_MINUS && t.State == sdl.PRESSED {
			e.camera.changePosition(Vector3D{0, 0, MOVE_SPEED})
		}
		// else if t.Keysym.Sym == sdl.K_a && t.State == sdl.PRESSED {
		// 	cube.position.X -= MOVE_SPEED
		// } else if t.Keysym.Sym == sdl.K_d && t.State == sdl.PRESSED {
		// 	cube.position.X += MOVE_SPEED
		// } else if t.Keysym.Sym == sdl.K_w && t.State == sdl.PRESSED {
		// 	cube.position.Y -= MOVE_SPEED
		// } else if t.Keysym.Sym == sdl.K_s && t.State == sdl.PRESSED {
		// 	cube.position.Y += MOVE_SPEED
		// }
	}
	return false
}

func (e MyGameEngine) update() {
	// Переопределенная реализация метода update()
	drawedUi := false
	for _, gamgeObj := range *e.gameObjects {
		gamgeObj.angles.updateAngles()
		e.camera.updateObject(&gamgeObj)
		gamgeObj.draw(e.renderer)
		if !drawedUi {
			err := drawUI(e.renderer, gamgeObj.angles)
			if err != nil {
				fmt.Println("Error drawing UI:", err)
				return
			}
			drawedUi = true
		}

	}
}

func main() {
	gameEngine, err := initGameEngine(nil, WIDTH, HEIGHT, FOW)
	if err != nil {
		fmt.Println("Error creating GameEngine:", err)
		return
	}
	defer gameEngine.Close()

	myEngine := MyGameEngine{*gameEngine}
	cube := getCube3D(1)
	cube3 := getCube3D(1)
	cube3.position.Z += 3
	cube4 := getCube3D(1)
	cube4.position.Z += 6
	// cube := getPyramid3D(1)
	// cube := getSphere3D(5, 180)
	// cube := getTorus3D(10, 2, 50, 20)

	// Call the getLandscape3D function with the provided parameters
	myEngine.addGameObj(cube)
	myEngine.addGameObj(cube3)
	myEngine.addGameObj(cube4)
	fmt.Println("myEngine", myEngine.gameObjects)
	run(&FPS, myEngine)
}
