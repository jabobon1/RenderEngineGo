package main

import (
	"fmt"
	"renderEngineGo/pkg"

	"github.com/veandco/go-sdl2/sdl"
)

var FPS uint32 = 35
var SPEED uint32 = 5
var MOVE_SPEED float64 = 0.1

const (
	WIDTH      int32   = 2000
	HEIGHT     int32   = 1500
	FOW        float64 = 45.0 // Vertical field of view
	nearPlane          = 0.1
	farPlane           = 100.0
	zThreshold         = 0.001
)

type MyGameEngine struct {
	pkg.GameEngine
}

// HandleKeyBoardPress implements pkg.GameObjectInterface.
func (m MyGameEngine) HandleKeyBoardPress(event sdl.Event) bool {
	fmt.Println("HandleKeyBoardPress")
	return false
}

// Update implements pkg.GameObjectInterface.
func (m MyGameEngine) Update() {
	fmt.Println("Updated")
}

// func main() {
// 	gameEngine, err := pkg.InitGameEngine(nil, WIDTH, HEIGHT, FOW)
// 	if err != nil {
// 		fmt.Println("Error creating GameEngine:", err)
// 		return
// 	}
// 	defer gameEngine.Close()

// 	myEngine := MyGameEngine{*gameEngine}
// 	cube := pkg.GetCube3D(1)

// 	myEngine.AddGameObj(cube)
// 	pkg.Run(&FPS, myEngine)
// }

func main() {
	editor, err := NewSceneEditor()
	if err != nil {
		fmt.Printf("Error creating scene editor: %v\n", err)
		return
	}
	defer editor.Cleanup()

	editor.Run()
}
