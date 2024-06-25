package main

import (
	"fmt"
	"renderEngineGo/pkg"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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
	sceneEditor *SceneEditor
}

// HandleKeyBoardPress implements pkg.GameObjectInterface.
func (e MyGameEngine) HandleKeyBoardPress(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.MouseMotionEvent:
		e.sceneEditor.handleMouseMotion(t)
	case *sdl.MouseButtonEvent:
		e.sceneEditor.handleMouseClick(t)
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.K_a && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{-MOVE_SPEED, 0, 0})
		} else if t.Keysym.Sym == sdl.K_d && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{MOVE_SPEED, 0, 0})
		} else if t.Keysym.Sym == sdl.K_w && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{0, -MOVE_SPEED, 0})
		} else if t.Keysym.Sym == sdl.K_s && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{0, MOVE_SPEED, 0})
		} else if t.Keysym.Sym == sdl.K_KP_PLUS && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{0, 0, -MOVE_SPEED})
		} else if t.Keysym.Sym == sdl.K_MINUS && t.State == sdl.PRESSED {
			e.Camera.ChangePosition(pkg.Vector3D{0, 0, MOVE_SPEED})
		}
	}

	return false
}
func (e MyGameEngine) Update() {
	e.Renderer.SetDrawColor(pkg.WHITE.R, pkg.WHITE.G, pkg.WHITE.B, pkg.WHITE.A)
	e.Renderer.Clear()
	for _, gamgeObj := range *e.GameObjects {
		gamgeObj.Angles.UpdateAngles()
		e.Camera.UpdateObject(&gamgeObj)
	}
	// Ensure GameObjects is safe to access
	e.Camera.DrawObjects(e.Renderer, e.GameObjects)
	e.sceneEditor.DrawScene()

	e.Renderer.Present()

}

func main() {
	gameEngine, err := pkg.InitGameEngine(nil, WIDTH, HEIGHT, FOW)
	if err != nil {
		fmt.Println("Error creating GameEngine:", err)
		return
	}
	font, err := ttf.OpenFont("/usr/share/fonts/opentype/unifont/unifont.otf", FONT_SIZE)
	if err != nil {
		return
	}

	myEngine := MyGameEngine{*gameEngine, nil}

	editor := &SceneEditor{
		Renderer: gameEngine.Renderer,
		Window:   gameEngine.Window,
		Font:     font,
		AvailableFigures: map[string]func(pkg.Vector3D) pkg.GameObject3D{
			"Cube": pkg.GetCube3D,
		},
		// AvailableFigures: []string{"Cube", "Sphere", "Pyramid", "Cylinder"},
		// GameEngine: myEngine,
		AddGameObj: myEngine.AddGameObj,
	}
	myEngine.sceneEditor = editor

	editor.initializeButtons()

	defer gameEngine.Close()
	defer editor.Cleanup()

	fps := uint32(32)
	pkg.Run(&fps, myEngine)
	// editor.Run()
}
