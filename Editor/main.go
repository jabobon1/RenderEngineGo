package main

import (
	"fmt"
)

func main() {
	gameEngine, err := initGameEngine(nil, WIDTH, HEIGHT, FOW)
	if err != nil {
		fmt.Println("Error creating GameEngine:", err)
		return
	}
	defer gameEngine.Close()

	myEngine := MyGameEngine{*gameEngine}
	cube := getCube3D(1)
	// cube3 := getCube3D(1)
	// cube3.position.Z += 3
	// cube := getRectangle3D(Vector3D{2, 1, 3})
	// cube4.position.Z += 6
	// cube := getPyramid3D(1)
	// cube := getSphere3D(5, 22)
	// cube := getTorus3D(10, 2, 50, 20)

	// Call the getLandscape3D function with the provided parameters
	myEngine.addGameObj(cube)
	// myEngine.addGameObj(cube3)
	// myEngine.addGameObj(cube4)
	fmt.Println("myEngine", myEngine.gameObjects)
	run(&FPS, myEngine)
}
