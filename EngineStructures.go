package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector2D struct {
	X, Y float64
}
type Vector3D struct {
	X, Y, Z float64
}
type Vector4D struct {
	X, Y, Z, W float64
}

type AngleVelocity struct {
	angleX, angleY, angleZ          float64
	angleXVel, angleYVel, angleZVel float64
	currentAxeIdx                   int
}

func (aV *AngleVelocity) updateAngles() {
	aV.angleX = math.Mod(aV.angleX+aV.angleXVel, 360)
	aV.angleY = math.Mod(aV.angleY+aV.angleYVel, 360)
	aV.angleZ = math.Mod(aV.angleZ+aV.angleZVel, 360)
}
func (aV *AngleVelocity) changeAxe() {
	aV.currentAxeIdx = (aV.currentAxeIdx + 1) % len(axes)

}

func (aV *AngleVelocity) getAxName() string {
	switch axes[aV.currentAxeIdx] {
	case xAxe:
		return "X"
	case yAxe:
		return "Y"
	case zAxe:
		return "Z"
	}
	return ""
}

func (aV *AngleVelocity) changeAngleVelociity(up bool) {
	var adder float64 = 1
	if !up {
		adder = -1
	}
	axe := aV.currentAxeIdx

	switch axe {
	case xAxe:
		aV.angleXVel += adder
	case yAxe:
		aV.angleYVel += adder
	case zAxe:
		aV.angleZVel += adder
	}
}

type GameObject3D struct {
	vertices        *[]Vector3D
	updatedVertices []Vector3D
	updatedNormals  []Vector3D
	normalMap       []Vector3D
	indicies        *[][]int
	angles          *AngleVelocity
	position        *Vector3D
	rotation        *Vector3D
	size            Vector3D
}

type Matrix4x4 [4][4]float64

func (m1 Matrix4x4) Multiply(m2 Matrix4x4) Matrix4x4 {
	result := Matrix4x4{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}

func (m Matrix4x4) MultiplyVector(v Vector4D) Vector4D {
	result := Vector4D{}

	result.X = m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W
	result.Y = m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W
	result.Z = m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W
	result.W = m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W

	return result
}

type Camera struct {
	position Vector3D
	rotation Vector3D
	fovRadV  float64
	fovRadH  float64
}

func (c *Camera) changePosition(newPosition Vector3D) {
	c.position = Add(c.position, newPosition)
}

func (c *Camera) updateObject(gameObject *GameObject3D) {
	// // // Combine the rotation matrices
	rotationMatrix := XYZRotationMatrix(
		gameObject.angles.angleX,
		gameObject.angles.angleY,
		gameObject.angles.angleZ)

	for i := range gameObject.updatedNormals {
		rotatedNomals := rotationMatrix.MultiplyVector(
			Vector4D{gameObject.normalMap[i].X,
				gameObject.normalMap[i].Y,
				gameObject.normalMap[i].Z, 1})

		gameObject.updatedNormals[i] = Normalize(Vector3D{rotatedNomals.X, rotatedNomals.Y, rotatedNomals.Z})
	}

	// // Combine the translation and rotation matrices
	// transformMatrix := projection_matrix.Multiply(rotationMatrix)

	for i, vertex := range *gameObject.vertices {
		// Convert the vertex to a 4D vector
		vec4D := Vector4D{vertex.X, vertex.Y, vertex.Z, 1}
		rotatedMatrix := rotationMatrix.MultiplyVector(vec4D)

		y := rotatedMatrix.Y - c.position.Y - gameObject.position.Y
		x := rotatedMatrix.X - c.position.X - gameObject.position.X
		z := rotatedMatrix.Z - c.position.Z - gameObject.position.Z

		if math.Abs(z) < zThreshold {
			if z < 0 {
				z = -zThreshold
			} else {
				z = zThreshold
			}
		}
		// Update the vertex in the updatedVertices slice
		gameObject.updatedVertices[i] = c.projectedXY(x, y, z)
	}

}
func (c *Camera) projectedXY(x, y, z float64) Vector3D {
	projectedX := x * c.fovRadH
	projectedY := y * c.fovRadV
	projectedX = (projectedX/z + 1.0) * float64(WIDTH) * 0.5
	projectedY = (1.0 - projectedY/z) * float64(HEIGHT) * 0.5

	return Vector3D{projectedX, projectedY, z}
}

// func getNotIntersectedLines(gameObjects *[]GameObject3D) {
// 	sortedGameObjects := make([][2]Vector3D, 0) //descending

// 	for _, gamgeObj := range *gameObjects {
// 		for _, edge := range *gamgeObj.indicies {
// 			p1 := gamgeObj.updatedVertices[edge[0]]
// 			p2 := gamgeObj.updatedVertices[edge[1]]

// 			if p1.X != p2.X || p1.Y != p2.Y {
// 				sortedGameObjects = append(sortedGameObjects, [2]Vector3D{p1, p2})
// 			}
// 		}
// 	}
// 	sort.Slice(sortedGameObjects, func(i, j int) bool {
// 		//descending
// 		return math.Min(sortedGameObjects[i][0].Z, sortedGameObjects[i][1].Z) > math.Min(sortedGameObjects[j][0].Z, sortedGameObjects[j][1].Z)
// 	})

// 	drawableLines := make([][2]Vector2D, 0)
// 	baseLine := Line(sortedGameObjects[0][0], sortedGameObjects[0][1])

// 	for i, line := range sortedGameObjects {
// 		if i >= len(sortedGameObjects) {
// 			break
// 		}
// 		intersection, intersected := getIntersectionVector()

// 	}

// }
func (c *Camera) drawObjects(renderer *sdl.Renderer, gameObjects *[]GameObject3D) {
	renderer.SetDrawColor(WHITE.R, WHITE.G, WHITE.B, WHITE.A)
	renderer.Clear()
	renderer.SetDrawColor(BLACK.R, BLACK.G, BLACK.B, BLACK.A)

	points := make([]sdl.Point, 0)

	for _, gameObj := range *gameObjects {
		// Define the camera direction as a normalized vector pointing from the camera's position
		// cameraDirection := Normalize(Sub(*gameObj.position, Add(c.position, Vector3D{0, 1, 0})))

		fmt.Printf("Camera pos: x %f, y %f, z %f\n", c.position.X, c.position.Y, c.position.Z)
		// fmt.Println("cameraDirection: ", cameraDirection)

		fmt.Printf("GameObject: vertices %d, indices %d\n", len(*gameObj.vertices), len(*gameObj.indicies))
		fmt.Printf("Angles: x %f, y %f, z %f\n", gameObj.angles.angleX, gameObj.angles.angleY, gameObj.angles.angleZ)
		fmt.Printf("GameObject pos: x %f, y %f, z %f\n", gameObj.position.X, gameObj.position.Y, gameObj.position.Z)

		for idxPol, polygons := range *gameObj.indicies {
			// we add normal here because object pos != face pos, here we can take face pos by add size * normal and add obj pos
			cameraDirection := Normalize(Sub(*gameObj.position, Add(c.position, Mult(gameObj.size, gameObj.updatedNormals[idxPol]))))

			// center := Centroid(gameObj.updatedVertices[polygons[0]], gameObj.updatedVertices[polygons[1]], gameObj.updatedVertices[polygons[2]])
			// facePos := Add(*gameObj.position, center)
			// cameraDirection := Normalize(Sub(facePos, c.position))
			fmt.Println("cameraDirection: ", cameraDirection)

			p1 := gameObj.updatedVertices[polygons[0]]
			p2 := gameObj.updatedVertices[polygons[1]]
			p3 := gameObj.updatedVertices[polygons[2]]

			// t1 := Sub(gameObj.updatedNormals[polygons[1]], gameObj.updatedNormals[polygons[0]])
			// t2 := Sub(gameObj.updatedNormals[polygons[2]], gameObj.updatedNormals[polygons[0]]) // Correct the second vector subtraction
			// normal := Normalize(Cross(t1, t2))
			normal := gameObj.updatedNormals[idxPol]
			res := Dot(cameraDirection, normal)

			fmt.Println("PS ", p1, p2, p3, normal, res)

			// Only draw the face if it's visible (res < 0 means the face is facing the camera)
			if res > 0 {
				points = append(points, sdl.Point{X: int32(p1.X), Y: int32(p1.Y)})
				points = append(points, sdl.Point{X: int32(p2.X), Y: int32(p2.Y)})

				points = append(points, sdl.Point{X: int32(p2.X), Y: int32(p2.Y)})
				points = append(points, sdl.Point{X: int32(p3.X), Y: int32(p3.Y)})

				points = append(points, sdl.Point{X: int32(p3.X), Y: int32(p3.Y)})
				points = append(points, sdl.Point{X: int32(p1.X), Y: int32(p1.Y)})
			}
		}
		fmt.Println("PS end")
	}

	if len(points) > 0 {
		if err := renderer.DrawLines(points); err != nil {
			fmt.Printf("Error drawing lines: %v\n", err)
		}
	}
	renderer.Present()
}
