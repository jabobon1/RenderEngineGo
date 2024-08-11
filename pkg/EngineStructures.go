package pkg

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector interface {
	Add(v Vector) Vector
	Sub(v Vector) Vector
	Mult(v Vector) Vector
}

type Vector2D struct {
	X, Y float64
}

func (v1 Vector2D) Add(v Vector) Vector {
	v2 := v.(Vector2D)
	return Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vector2D) Sub(v Vector) Vector {
	v2 := v.(Vector2D)
	return Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}
func (v1 Vector2D) Mult(v Vector) Vector {
	v2 := v.(Vector2D)
	return Vector2D{v1.X * v2.X, v1.Y * v2.Y}
}

type Vector3D struct {
	X, Y, Z float64
}

func (v1 Vector3D) Add(v Vector) Vector {
	v2 := v.(Vector3D)
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector3D) Sub(v Vector) Vector {
	v2 := v.(Vector3D)
	return Vector3D{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 Vector3D) Mult(v Vector) Vector {
	v2 := v.(Vector3D)
	return Vector3D{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

type Vector4D struct {
	X, Y, Z, W float64
}

type AngleVelocity struct {
	angleX, angleY, angleZ          float64
	angleXVel, angleYVel, angleZVel float64
	currentAxeIdx                   int
}

func (aV *AngleVelocity) UpdateAngles() {
	aV.angleX = math.Mod(aV.angleX+aV.angleXVel, 360)
	aV.angleY = math.Mod(aV.angleY+aV.angleYVel, 360)
	aV.angleZ = math.Mod(aV.angleZ+aV.angleZVel, 360)
}
func (aV *AngleVelocity) ChangeAxe() {
	aV.currentAxeIdx = (aV.currentAxeIdx + 1) % len(axes)

}

func (aV *AngleVelocity) GetAxName() string {
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

func (aV *AngleVelocity) ChangeAngleVelociity(up bool) {
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
	colorMap        []sdl.Color
	indicies        *[][]int
	Angles          *AngleVelocity
	Position        *Vector3D
	Rotation        *Vector3D
	Size            Vector3D
}

func (gm GameObject3D) GetMinMaxPointsOnScreen() (float64, float64, float64, float64) {
	minX, maxX, minY, maxY := float64(math.Inf(1)), float64(0), float64(math.Inf(1)), float64(0)

	for _, vert := range gm.updatedVertices {
		minX = math.Min(vert.X, minX)
		minY = math.Min(vert.Y, minY)

		maxX = math.Max(vert.X, maxX)
		maxY = math.Max(vert.Y, maxY)
	}

	return minX, maxX, minY, maxY

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

func (c *Camera) ChangePosition(newPosition Vector3D) {
	c.position = Add(c.position, newPosition)
}

func (c *Camera) UpdateObject(gameObject *GameObject3D) {
	// // // Combine the rotation matrices
	rotationMatrix := XYZRotationMatrix(
		gameObject.Angles.angleX,
		gameObject.Angles.angleY,
		gameObject.Angles.angleZ)

	for i := range gameObject.updatedNormals {
		rotatedNomals := rotationMatrix.MultiplyVector(
			Vector4D{-gameObject.normalMap[i].X,
				-gameObject.normalMap[i].Y,
				-gameObject.normalMap[i].Z,
				1})

		gameObject.updatedNormals[i] = Normalize(Vector3D{rotatedNomals.X, rotatedNomals.Y, rotatedNomals.Z})
	}

	// // Combine the translation and rotation matrices
	// transformMatrix := projection_matrix.Multiply(rotationMatrix)

	for i, vertex := range *gameObject.vertices {
		// Convert the vertex to a 4D vector
		vec4D := Vector4D{vertex.X, vertex.Y, vertex.Z, 1}
		rotatedMatrix := rotationMatrix.MultiplyVector(vec4D)

		y := rotatedMatrix.Y + c.position.Y - gameObject.Position.Y
		x := rotatedMatrix.X + c.position.X - gameObject.Position.X
		z := rotatedMatrix.Z + c.position.Z - gameObject.Position.Z

		// if math.Abs(z) < zThreshold {
		// 	if z < 0 {
		// 		z = -zThreshold
		// 	} else {
		// 		z = zThreshold
		// 	}
		// }
		// Update the vertex in the updatedVertices slice
		gameObject.updatedVertices[i] = c.projectedXY(x, y, z)
	}

}

func (c *Camera) HandleCameraMove(t sdl.KeyboardEvent, moveSpeed float64) bool {
	if t.Keysym.Sym == sdl.K_a && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{-moveSpeed, 0, 0})
		return true
	} else if t.Keysym.Sym == sdl.K_d && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{moveSpeed, 0, 0})
		return true
	} else if t.Keysym.Sym == sdl.K_w && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{0, moveSpeed, 0})
		return true

	} else if t.Keysym.Sym == sdl.K_s && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{0, -moveSpeed, 0})
		return true
	} else if t.Keysym.Sym == sdl.K_KP_PLUS && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{0, 0, moveSpeed})
		return true
	} else if t.Keysym.Sym == sdl.K_MINUS && t.State == sdl.PRESSED {
		c.ChangePosition(Vector3D{0, 0, -moveSpeed})
		return true
	}
	return false
}

const (
	WIDTH  int32 = 2000
	HEIGHT int32 = 1500
)

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

type LinePoints struct {
	p1, p2 sdl.Point
}

func getLine(p1, p2 sdl.Point) LinePoints {
	if p1.X < p2.X {
		return LinePoints{p1, p2}
	} else if p1.X == p2.X {
		if p1.Y < p2.Y {
			return LinePoints{p1, p2}
		}
		return LinePoints{p2, p1}
	}
	return LinePoints{p2, p1}

}

func vertexFromPoint(p Vector3D, color sdl.Color) sdl.Vertex {
	return sdl.Vertex{
		Position: sdl.FPoint{X: float32(p.X), Y: float32(p.Y)},
		Color:    color,
		TexCoord: sdl.FPoint{X: 0, Y: 0},
	}

}

func (c *Camera) DrawObjects(renderer *sdl.Renderer, gameObjects *[]GameObject3D) {
	renderer.SetDrawColor(BLACK.R, BLACK.G, BLACK.B, BLACK.A)

	colorPoints := make([][]sdl.Vertex, 0)
	uniquePoints := make(map[LinePoints]int)

	for _, gameObj := range *gameObjects {
		fmt.Println("gameObj", gameObj.Position, "Camera", c.position)

		for idxPol, polygons := range *gameObj.indicies {
			// we add normal here because object pos != face pos, here we can take face pos by add size * normal and add obj pos
			// cameraDirection := Normalize(Sub(*gameObj.Position, Add(c.position, Mult(gameObj.Size, gameObj.updatedNormals[idxPol]))))
			pos1 := *gameObj.Position
			pos2 := Add(c.position, Mult(gameObj.Size, gameObj.updatedNormals[idxPol]))

			cameraDirection := Normalize(Sub(pos1, pos2))

			p1 := gameObj.updatedVertices[polygons[0]]
			p2 := gameObj.updatedVertices[polygons[1]]
			p3 := gameObj.updatedVertices[polygons[2]]

			normal := gameObj.updatedNormals[idxPol]
			res := Dot(cameraDirection, normal)
			fmt.Println("cameraDirection", cameraDirection, "normal", normal, "res Dot", res)

			points := []sdl.Point{
				{X: int32(p1.X), Y: int32(p1.Y)},
				{X: int32(p2.X), Y: int32(p2.Y)},
				{X: int32(p3.X), Y: int32(p3.Y)},
			}

			if res > 0 {
				for i := range points {
					_p1 := points[i]
					_p2 := points[0]
					if i != len(points)-1 {
						_p2 = points[i+1]
					}

					line := getLine(_p1, _p2)
					uniquePoints[line]++
				}

				color := gameObj.colorMap[idxPol]
				vertices := []sdl.Vertex{
					vertexFromPoint(p1, color),
					vertexFromPoint(p2, color),
					vertexFromPoint(p3, color),
				}
				colorPoints = append(colorPoints, vertices)

			}
		}
	}

	for line := range uniquePoints {
		if err := renderer.DrawLine(line.p1.X, line.p1.Y, line.p2.X, line.p2.Y); err != nil {
			fmt.Printf("Error drawing lines: %v\n", err)
		}
	}

	for _, vertices := range colorPoints {
		renderer.RenderGeometry(nil, vertices, nil)
	}

}
