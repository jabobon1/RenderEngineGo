package main

func getCube3D(size float64) ([]Point3D, [][]int) {
	vertices := []Point3D{
		{X: -size, Y: -size, Z: size},
		{X: size, Y: -size, Z: size},
		{X: size, Y: size, Z: size},
		{X: -size, Y: size, Z: size},
		{X: -size, Y: -size, Z: -size},
		{X: size, Y: -size, Z: -size},
		{X: size, Y: size, Z: -size},
		{X: -size, Y: size, Z: -size},
	}
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Front face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Back face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Connecting edges
	}
	return vertices, edges
}
