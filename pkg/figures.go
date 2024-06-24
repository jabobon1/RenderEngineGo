package pkg

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func GetCube3D(size Vector3D) GameObject3D {
	vertices := []Vector3D{
		{-size.X, -size.Y, -size.Z}, // 0
		{size.X, -size.Y, -size.Z},  // 1
		{size.X, size.Y, -size.Z},   // 2
		{-size.X, size.Y, -size.Z},  // 3
		{-size.X, -size.Y, size.Z},  // 4
		{size.X, -size.Y, size.Z},   // 5
		{size.X, size.Y, size.Z},    // 6
		{-size.X, size.Y, size.Z},   // 7
	}
	rotatedVertices := make([]Vector3D, len(vertices))

	indices := [][]int{
		{4, 5, 6}, {4, 6, 7}, // Front face
		{0, 1, 2}, {0, 2, 3}, // Back  face
		{0, 1, 5}, {0, 5, 4}, // Top face
		{3, 2, 6}, {3, 6, 7}, // Buttom face
		{0, 4, 7}, {0, 7, 3}, // Right face
		{1, 5, 6}, {1, 6, 2}, // Left face
	}

	normalMap := []Vector3D{
		{0, 0, -1}, {0, 0, -1}, // Front face
		{0, 0, 1}, {0, 0, 1}, // Back face
		{0, 1, 0}, {0, 1, 0}, // Top face
		{0, -1, 0}, {0, -1, 0}, // Buttom face
		{1, 0, 0}, {1, 0, 0}, // Right face
		{-1, 0, 0}, {-1, 0, 0}, // Left face
	}

	colorMap := []sdl.Color{
		{R: 255, G: 0, B: 0, A: 255}, {R: 255, G: 0, B: 0, A: 255},
		{R: 200, G: 0, B: 0, A: 255}, {R: 200, G: 0, B: 0, A: 255},
		{R: 150, G: 0, B: 0, A: 255}, {R: 150, G: 0, B: 0, A: 255},
		{R: 100, G: 0, B: 0, A: 255}, {R: 100, G: 0, B: 0, A: 255},
		{R: 100, G: 0, B: 0, A: 255}, {R: 100, G: 0, B: 0, A: 255},
		{R: 100, G: 0, B: 0, A: 255}, {R: 100, G: 0, B: 0, A: 255},
	}

	updatedNormals := make([]Vector3D, len(normalMap))

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{0, 0, 0}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	return GameObject3D{&vertices, rotatedVertices, updatedNormals, normalMap, colorMap, &indices, &angleVelocity, &position, &rotation, size}
}

func GetPyramid3D(size float64) GameObject3D {
	vertices := []Vector3D{
		{-size, -size, -size}, // 0
		{size, -size, -size},  // 1
		{0, size, 0},          // 2
		{0, -size, size},      // 3
	}
	rotatedVertices := make([]Vector3D, len(vertices))

	indices := [][]int{
		{0, 1, 2}, // Side triangle 1
		{1, 3, 2}, // Side triangle 2
		{3, 0, 2}, // Side triangle 3
		{0, 1, 3}, // Base triangle
	}

	normalMap := []Vector3D{
		{-1, 1, -1}, // Side triangle 1
		{1, 1, -1},  // Side triangle 2
		{0, 1, 1},   // Side triangle 3
		{0, -1, 0},  // Base triangle
	}

	colorMap := []sdl.Color{
		{R: 255, G: 0, B: 0, A: 255},
		{R: 200, G: 0, B: 0, A: 255},
		{R: 150, G: 0, B: 0, A: 255},
		{R: 100, G: 0, B: 0, A: 255},
	}

	updatedNormals := make([]Vector3D, len(normalMap))

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{0, 0, 0}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	return GameObject3D{&vertices, rotatedVertices, updatedNormals, normalMap, colorMap, &indices, &angleVelocity, &position, &rotation, Vector3D{size, size, size}}
}

func GetSphere3D(radius float64, segments int) GameObject3D {
	vertices := make([]Vector3D, 0)
	for i := 0; i <= segments; i++ {
		phi := float64(i) * math.Pi / float64(segments)
		for j := 0; j <= segments; j++ {
			theta := float64(j) * 2 * math.Pi / float64(segments)
			x := radius * math.Sin(phi) * math.Cos(theta)
			y := radius * math.Cos(phi)
			z := radius * math.Sin(phi) * math.Sin(theta)
			vertices = append(vertices, Vector3D{x, y, z})
		}
	}
	rotatedVertices := make([]Vector3D, len(vertices))

	indices := make([][]int, 0)
	for i := 0; i < segments; i++ {
		for j := 0; j < segments; j++ {
			a := i*(segments+1) + j
			b := a + 1
			c := (i+1)*(segments+1) + j
			d := c + 1
			indices = append(indices, []int{a, b, c}, []int{b, d, c})
		}
	}

	normalMap := make([]Vector3D, len(indices))
	for i, triangle := range indices {
		v0, v1, v2 := vertices[triangle[0]], vertices[triangle[1]], vertices[triangle[2]]
		normalMap[i] = calculateNormal(v0, v1, v2)
	}

	colorMap := make([]sdl.Color, len(indices))
	for i := range colorMap {
		colorMap[i] = sdl.Color{R: uint8(128 + i%128), G: 0, B: uint8(128 - i%128), A: 255}
	}

	updatedNormals := make([]Vector3D, len(normalMap))

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{0, 0, 0}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	return GameObject3D{&vertices, rotatedVertices, updatedNormals, normalMap, colorMap, &indices, &angleVelocity, &position, &rotation, Vector3D{radius * 2, radius * 2, radius * 2}}
}

func GetTorus3D(majorRadius, minorRadius float64, majorSegments, minorSegments int) GameObject3D {
	vertices := make([]Vector3D, 0)

	for i := 0; i < majorSegments; i++ {
		majorAngle := float64(i) * 2 * math.Pi / float64(majorSegments)
		for j := 0; j < minorSegments; j++ {
			minorAngle := float64(j) * 2 * math.Pi / float64(minorSegments)
			x := (majorRadius + minorRadius*math.Cos(minorAngle)) * math.Cos(majorAngle)
			y := (majorRadius + minorRadius*math.Cos(minorAngle)) * math.Sin(majorAngle)
			z := minorRadius * math.Sin(minorAngle)
			vertices = append(vertices, Vector3D{x, y, z})
		}
	}
	rotatedVertices := make([]Vector3D, len(vertices))

	indices := make([][]int, 0)
	for i := 0; i < majorSegments; i++ {
		for j := 0; j < minorSegments; j++ {
			a := i*minorSegments + j
			b := (i*minorSegments + (j+1)%minorSegments) % len(vertices)
			c := ((i+1)%majorSegments)*minorSegments + j
			d := ((i+1)%majorSegments)*minorSegments + (j+1)%minorSegments
			indices = append(indices, []int{a, b, c}, []int{b, d, c})
		}
	}

	normalMap := make([]Vector3D, len(indices))
	for i, triangle := range indices {
		v0, v1, v2 := vertices[triangle[0]], vertices[triangle[1]], vertices[triangle[2]]
		normalMap[i] = calculateNormal(v0, v1, v2)
	}

	colorMap := make([]sdl.Color, len(indices))
	for i := range colorMap {
		colorMap[i] = sdl.Color{R: uint8(128 + i%128), G: uint8(i % 256), B: uint8(128 - i%128), A: 255}
	}

	updatedNormals := make([]Vector3D, len(normalMap))

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{0, 0, 0}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	size := Vector3D{(majorRadius + minorRadius) * 2, (majorRadius + minorRadius) * 2, minorRadius * 2}
	return GameObject3D{&vertices, rotatedVertices, updatedNormals, normalMap, colorMap, &indices, &angleVelocity, &position, &rotation, size}
}

func calculateNormal(v0, v1, v2 Vector3D) Vector3D {
	u := Vector3D{v1.X - v0.X, v1.Y - v0.Y, v1.Z - v0.Z}
	v := Vector3D{v2.X - v0.X, v2.Y - v0.Y, v2.Z - v0.Z}

	normal := Vector3D{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}

	length := math.Sqrt(normal.X*normal.X + normal.Y*normal.Y + normal.Z*normal.Z)
	if length != 0 {
		normal.X /= length
		normal.Y /= length
		normal.Z /= length
	}

	return normal
}
