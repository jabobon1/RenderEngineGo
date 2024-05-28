package main

import "math"

func getCube3D(size float64) GameObject3D {
	vertices := []Vector3D{
		{-size, -size, -size},
		{size, -size, -size},
		{size, size, -size},
		{-size, size, -size},
		{-size, -size, size},
		{size, -size, size},
		{size, size, size},
		{-size, size, size},
	}
	rotatedVertices := make([]Vector3D, len(vertices))

	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Front face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Back face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Connecting edges
	}
	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{10, 10, 10}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}

	return GameObject3D{&vertices, rotatedVertices, &edges, &angleVelocity, &position, &rotation}
}

func getPyramid3D(size float64) GameObject3D {
	vertices := []Vector3D{
		{-size, -size, -size},
		{size, -size, -size},
		{0, size, 0},
		{0, -size, size},
	}
	rotatedVertices := make([]Vector3D, len(vertices))
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 0}, // Base edges
		{0, 3}, {1, 3}, {2, 3}, // Side edges
	}
	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{10, 10, 10}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
	return GameObject3D{&vertices, rotatedVertices, &edges, &angleVelocity, &position, &rotation}
}

func getSphere3D(radius float64, segments int) GameObject3D {
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
	edges := make([][]int, 0)
	for i := 0; i < segments; i++ {
		for j := 0; j < segments; j++ {
			a := i*(segments+1) + j
			b := a + 1
			c := (i+1)*(segments+1) + j
			d := c + 1
			edges = append(edges, []int{a, b}, []int{a, c}, []int{b, d}, []int{c, d})
		}
	}
	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{10, 10, 10}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
	return GameObject3D{&vertices, rotatedVertices, &edges, &angleVelocity, &position, &rotation}
}

func getTorus3D(majorRadius, minorRadius float64, majorSegments, minorSegments int) GameObject3D {
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
	edges := make([][]int, 0)

	for i := 0; i < majorSegments; i++ {
		for j := 0; j < minorSegments; j++ {
			a := i*minorSegments + j
			b := (i*minorSegments + (j+1)%minorSegments) % len(vertices)
			c := ((i+1)%majorSegments)*minorSegments + j
			d := ((i+1)%majorSegments)*minorSegments + (j+1)%minorSegments
			edges = append(edges, []int{a, b}, []int{a, c}, []int{b, d}, []int{c, d})
		}
	}

	angleVelocity := AngleVelocity{0, 0, 0, 0, 0, 0, 0}
	position := Vector3D{10, 10, 10}
	rotation := Vector3D{angleVelocity.angleX, angleVelocity.angleY, angleVelocity.angleZ}
	return GameObject3D{&vertices, rotatedVertices, &edges, &angleVelocity, &position, &rotation}
}
