package main

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
