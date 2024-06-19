package main

func getRectangle3D(size Vector3D) GameObject3D {
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
	updatedNormals := make([]Vector3D, len(normalMap))

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{0, 0, 0}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	return GameObject3D{&vertices, rotatedVertices, updatedNormals, normalMap, &indices, &angleVelocity, &position, &rotation, size}
}

func getCube3D(size float64) GameObject3D {
	return getRectangle3D(Vector3D{size, size, size})
}

// func getPyramid3D(size float64) GameObject3D {
// 	vertices := []Vector3D{
// 		{-size, -size, -size},
// 		{size, -size, -size},
// 		{0, size, 0},
// 		{0, -size, size},
// 	}
// 	rotatedVertices := make([]Vector3D, len(vertices))
// 	rotatedVertices3D := make([]Vector3D, len(vertices))

// 	indices := [][]int{
// 		{0, 1, 2}, // Base triangle 1
// 		{0, 2, 3}, // Base triangle 2
// 		{0, 1, 3}, // Side triangle 1
// 		{1, 2, 3}, // Side triangle 2
// 	}

// 	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
// 	position := Vector3D{10, 10, 10}
// 	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
// 	return GameObject3D{&vertices, rotatedVertices, rotatedVertices3D, &indices, &angleVelocity, &position, &rotation}
// }

// func getSphere3D(radius float64, segments int) GameObject3D {
// 	vertices := make([]Vector3D, 0)
// 	for i := 0; i <= segments; i++ {
// 		phi := float64(i) * math.Pi / float64(segments)
// 		for j := 0; j <= segments; j++ {
// 			theta := float64(j) * 2 * math.Pi / float64(segments)
// 			x := radius * math.Sin(phi) * math.Cos(theta)
// 			y := radius * math.Cos(phi)
// 			z := radius * math.Sin(phi) * math.Sin(theta)
// 			vertices = append(vertices, Vector3D{x, y, z})
// 		}
// 	}
// 	rotatedVertices := make([]Vector3D, len(vertices))

// 	indices := make([][]int, 0)
// 	for i := 0; i < segments; i++ {
// 		for j := 0; j < segments; j++ {
// 			a := i*(segments+1) + j
// 			b := a + 1
// 			c := (i+1)*(segments+1) + j
// 			d := c + 1
// 			indices = append(indices, []int{a, b, c}, []int{b, d, c})
// 		}
// 	}

// 	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
// 	position := Vector3D{10, 10, 10}
// 	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
// 	return GameObject3D{&vertices, rotatedVertices, &indices, &angleVelocity, &position, &rotation}
// }

// func getTorus3D(majorRadius, minorRadius float64, majorSegments, minorSegments int) GameObject3D {
// 	vertices := make([]Vector3D, 0)

// 	for i := 0; i < majorSegments; i++ {
// 		majorAngle := float64(i) * 2 * math.Pi / float64(majorSegments)
// 		for j := 0; j < minorSegments; j++ {
// 			minorAngle := float64(j) * 2 * math.Pi / float64(minorSegments)
// 			x := (majorRadius + minorRadius*math.Cos(minorAngle)) * math.Cos(majorAngle)
// 			y := (majorRadius + minorRadius*math.Cos(minorAngle)) * math.Sin(majorAngle)
// 			z := minorRadius * math.Sin(minorAngle)
// 			vertices = append(vertices, Vector3D{x, y, z})
// 		}
// 	}
// 	rotatedVertices := make([]Vector3D, len(vertices))

// 	indices := make([][]int, 0)
// 	for i := 0; i < majorSegments; i++ {
// 		for j := 0; j < minorSegments; j++ {
// 			a := i*minorSegments + j
// 			b := (i*minorSegments + (j+1)%minorSegments) % len(vertices)
// 			c := ((i+1)%majorSegments)*minorSegments + j
// 			d := ((i+1)%majorSegments)*minorSegments + (j+1)%minorSegments
// 			indices = append(indices, []int{a, b, c}, []int{b, d, c})
// 		}
// 	}

// 	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
// 	position := Vector3D{10, 10, 10}
// 	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
// 	return GameObject3D{&vertices, rotatedVertices, &indices, &angleVelocity, &position, &rotation}
// }
