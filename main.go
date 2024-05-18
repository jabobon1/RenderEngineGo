package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Point3D struct {
	X, Y, Z float64
}

func rotateX(point Point3D, angle float64) Point3D {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)
	return Point3D{
		X: point.X,
		Y: point.Y*cos - point.Z*sin,
		Z: point.Y*sin + point.Z*cos,
	}
}

func rotateY(point Point3D, angle float64) Point3D {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)
	return Point3D{
		X: point.X*cos + point.Z*sin,
		Y: point.Y,
		Z: -point.X*sin + point.Z*cos,
	}
}

func rotateZ(point Point3D, angle float64) Point3D {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)
	return Point3D{
		X: point.X*cos - point.Y*sin,
		Y: point.X*sin + point.Y*cos,
		Z: point.Z,
	}
}

func getVertices() []Point3D {
	vertices := []Point3D{
		{X: -100, Y: -100, Z: 100},
		{X: 100, Y: -100, Z: 100},
		{X: 100, Y: 100, Z: 100},
		{X: -100, Y: 100, Z: 100},
		{X: -100, Y: -100, Z: -100},
		{X: 100, Y: -100, Z: -100},
		{X: 100, Y: 100, Z: -100},
		{X: -100, Y: 100, Z: -100},
	}
	return vertices
}

func getEdges() [][]int {
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Front face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Back face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Connecting edges
	}
	return edges
}

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

	vertices := getVertices()
	edges := getEdges()
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
			rotated = rotateX(rotated, angleX)
			rotated = rotateY(rotated, angleY)
			rotated = rotateZ(rotated, angleZ)
			rotatedVertices[i] = rotated
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 255)
		for _, edge := range edges {
			renderer.DrawLine(int32(rotatedVertices[edge[0]].X+400), int32(rotatedVertices[edge[0]].Y+300),
				int32(rotatedVertices[edge[1]].X+400), int32(rotatedVertices[edge[1]].Y+300))
		}

		renderer.Present()
		sdl.Delay(16)
	}

}
