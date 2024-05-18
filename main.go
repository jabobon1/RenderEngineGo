package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var FPS uint32 = 60
var SPEED uint32 = 5
var MOVE_SPEED float64 = 10

const (
	xAxe int = 0
	yAxe int = 1
	zAxe int = 2
)

var axes = [3]int{xAxe, yAxe, zAxe}

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

func drawUI(renderer *sdl.Renderer, angleVelocityObj *AngleVelocity) error {
	var x int32 = 10
	var y int32 = 10
	var offset int32 = 5

	rect, err := drawText(renderer, fmt.Sprintf("FPS: %d", FPS), x, y, BLACK)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset
	rect, err = drawText(renderer, fmt.Sprintf("Velocity: %d", SPEED), x, y, BLACK)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset
	rect, err = drawText(
		renderer,
		fmt.Sprintf("Angle X: %.2f, Y: %.2f, Z: %.2f", angleVelocityObj.angleX, angleVelocityObj.angleY, angleVelocityObj.angleZ),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}

	y = rect.Y + rect.H + offset
	rect, err = drawText(
		renderer,
		fmt.Sprintf("Angle Velocity: %.2f, Y: %.2f, Z: %.2f", angleVelocityObj.angleXVel, angleVelocityObj.angleYVel, angleVelocityObj.angleZVel),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset

	currentAx := angleVelocityObj.getAxName()

	rect, err = drawText(
		renderer,
		fmt.Sprintf("currentAx: %s", currentAx),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}

	return nil
}

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println("Error initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Cube", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1500, 1200, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error creating window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	if err := ttf.Init(); err != nil {
		fmt.Println("Error initializing SDL_ttf:", err)
		return
	}
	defer ttf.Quit()
	cube := getCube3D(300)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				return
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
		}
		cube.rotate()

		renderer.SetDrawColor(WHITE.R, WHITE.G, WHITE.B, WHITE.A)
		renderer.Clear()

		renderer.SetDrawColor(BLACK.R, BLACK.G, BLACK.B, BLACK.A)
		cube.draw(renderer)
		err := drawUI(renderer, cube.angles)
		if err != nil {
			fmt.Println("Error drawing UI:", err)
			return
		}
		renderer.Present()
		sdl.Delay(1000 / FPS)
	}

}
