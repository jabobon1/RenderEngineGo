package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var FPS uint32 = 35
var SPEED uint32 = 5
var MOVE_SPEED float64 = 0.1

const (
	WIDTH      int32   = 1200
	HEIGHT     int32   = 700
	fov        float64 = 45.0 // Vertical field of view
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

type Camera struct {
	position    Vector3D
	rotation    Vector3D
	previousPos Vector3D
	previousRot Vector3D
}

func (c *Camera) changePosition(newPosition Vector3D) {
	c.previousPos = c.position
	c.position = Add(c.position, newPosition)
}

func (c *Camera) updateObject(gameObject *GameObject3D) {
	fmt.Println("gameObject ", gameObject.vertices)
	fmt.Println("c ", c.position)

	aspectRatio := float64(WIDTH) / float64(HEIGHT)
	fovRadV := 1.0 / math.Tan(fov*0.5/180.0*math.Pi)
	fovRadH := fovRadV / aspectRatio

	// projection_matrix := Matrix4x4{
	// 	{aspect_ratio * focus, 0, 0, 0},
	// 	{0, focus, 0, 0},
	// 	{0, 0, z_far / (z_far - z_near), -(z_far * z_near) / (z_far - z_near)},
	// 	{0, 0, 0, 1},
	// }

	rad := gameObject.angles.angleX * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)

	// Create rotation matrices based on the camera's updated rotation angles
	rotateX := Matrix4x4{
		{1, 0, 0, 0},
		{0, cos, -sin, 0},
		{0, sin, cos, 0},
		{0, 0, 0, 1},
	}

	rad = gameObject.angles.angleY * math.Pi / 180.0
	sin, cos = math.Sin(rad), math.Cos(rad)

	rotateY := Matrix4x4{
		{cos, 0, sin, 0},
		{0, 1, 0, 0},
		{-sin, 0, cos, 0},
		{0, 0, 0, 1},
	}

	rad = gameObject.angles.angleZ * math.Pi / 180.0
	sin, cos = math.Sin(rad), math.Cos(rad)

	rotateZ := Matrix4x4{
		{cos, -sin, 0, 0},
		{sin, cos, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	// // // Combine the rotation matrices
	rotationMatrix := rotateX.Multiply(rotateY).Multiply(rotateZ)

	// fmt.Println("rotationMatrix ", rotationMatrix)

	// // Combine the translation and rotation matrices
	// transformMatrix := projection_matrix.Multiply(rotationMatrix)

	for i, vertex := range *gameObject.vertices {
		// Convert the vertex to a 4D vector
		vec4D := Vector4D{vertex.X, vertex.Y, vertex.Z, 1}
		rotatedMatrix := rotationMatrix.MultiplyVector(vec4D)

		fmt.Println("vec4D ", vec4D)
		fmt.Println("rotatedMatrix ", rotatedMatrix)
		y := rotatedMatrix.Y - c.position.Y
		x := rotatedMatrix.X - c.position.X
		z := rotatedMatrix.Z - c.position.Z

		projectedX := x * fovRadH
		projectedY := y * fovRadV

		if math.Abs(z) < zThreshold {
			if z < 0 {
				z = -zThreshold
			} else {
				z = zThreshold
			}
		}

		projectedX = (projectedX/z + 1.0) * float64(WIDTH) * 0.5
		projectedY = (1.0 - projectedY/z) * float64(HEIGHT) * 0.5

		fmt.Println("vertex ", vertex)
		fmt.Println("projectedX ", projectedX, "projectedY ", projectedY)

		// Apply the transformation matrix to the vertex
		// transformedVec4D := transformMatrix.MultiplyVector(vec4D)

		// Convert the transformed 4D vector back to a 3D vector
		// transformedVertex := Vector3D{transformedVec4D.X * 70 / vertex.Z, transformedVec4D.Y * 70 / vertex.Z, transformedVec4D.Z}
		// fmt.Println("transformedVertex:", transformedVertex)
		// Update the vertex in the updatedVertices slice
		gameObject.updatedVertices[i] = Vector3D{projectedX, projectedY, 1}
	}

	// angle_xyz := AngleBetween(c.position, *gameObject.position)
	// angle_xyz = Add(angle_xyz, Vector3D{gameObject.angles.angleX, gameObject.angles.angleY, gameObject.angles.angleZ})
	// fmt.Println("c.position ", c.position, " gameObject.position ", gameObject.position)
	// fmt.Println("angle_xyz ", angle_xyz)

	// for i, vertex := range *gameObject.vertices {
	// 	vertex.rotate(
	// 		angle_xyz.X,
	// 		angle_xyz.Y,
	// 		angle_xyz.Z,
	// 	)
	// 	newVec := Sub(vertex, c.position)
	// 	gameObject.updatedVertices[i] = newVec
	// }

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
	var cube *GameObject3D = e.getGameObject()
	// cube.rotate()
	cube.angles.updateAngles()
	e.camera.updateObject(cube)

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
	gameEngine, err := initGameEngine(nil, WIDTH, HEIGHT)
	if err != nil {
		fmt.Println("Error creating GameEngine:", err)
		return
	}
	defer gameEngine.Close()

	myEngine := MyGameEngine{*gameEngine}
	cube := getCube3D(1)
	myEngine.addGameObj(cube)
	fmt.Println("myEngine", myEngine.gameObjects)
	run(&FPS, myEngine)
}
