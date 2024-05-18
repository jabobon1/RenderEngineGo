package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println("Error initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Cube", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
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

	vertices, edges := getCube3D()
	angleX, angleY, angleZ := 0.0, 0.0, 0.0

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// Update rotation angles
		angleX += 1.0
		angleY += 1.0
		angleZ += 1.0

		// Rotate vertices
		rotatedVertices := make([]Point3D, len(vertices))
		for i, vertex := range vertices {
			rotated := vertex
			rotated.rotateX(angleX)
			rotated.rotateY(angleY)
			rotated.rotateZ(angleZ)
			rotatedVertices[i] = rotated
		}

		renderer.SetDrawColor(WHITE.r, WHITE.g, WHITE.b, WHITE.a)
		renderer.Clear()

		renderer.SetDrawColor(BLACK.r, BLACK.g, BLACK.b, BLACK.a)
		for _, edge := range edges {
			renderer.DrawLine(int32(rotatedVertices[edge[0]].X+400), int32(rotatedVertices[edge[0]].Y+300),
				int32(rotatedVertices[edge[1]].X+400), int32(rotatedVertices[edge[1]].Y+300))
		}

		renderer.Present()
		sdl.Delay(16)
	}

}
