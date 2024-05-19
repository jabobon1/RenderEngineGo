package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GameEngine struct {
	gameObjects *[]GameObject3D
	window      *sdl.Window
	renderer    *sdl.Renderer
}

func (g *GameEngine) Close() {
	// выполнить очистку или освобождение ресурсов
	g.renderer.Destroy()
	g.window.Destroy()
	sdl.Quit()
}

func initGameEngine(gameObjects *[]GameObject3D, width, heigh int32) (*GameEngine, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println("Error initializing SDL:", err)
		return nil, err
	}

	window, err := sdl.CreateWindow("Cube", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, heigh, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error creating window:", err)
		sdl.Quit()
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error creating renderer:", err)
		window.Destroy()
		sdl.Quit()
		return nil, err
	}
	if err := ttf.Init(); err != nil {
		fmt.Println("Error initializing SDL_ttf:", err)
		renderer.Destroy()
		window.Destroy()
		sdl.Quit()
		return nil, err

	}
	if gameObjects == nil {
		gObjs := make([]GameObject3D, 0)
		gameObjects = &gObjs
	}

	return &GameEngine{gameObjects, window, renderer}, nil

}

type GameObject3DAbs interface {
	update()
}

func (e *GameEngine) update() {
	fmt.Println("Original update has been called")

}

func (e *GameEngine) getGameObject() *GameObject3D {
	for _, gObg := range *e.gameObjects {
		return &gObg
	}
	return nil

}

func (e *GameEngine) addGameObj(gameObject GameObject3D) {
	*e.gameObjects = append(*e.gameObjects, gameObject)

}

type GameObjectInterface interface {
	update()
	handleKeyBoardPress(event sdl.Event) bool
}

func run(e GameObjectInterface, fps uint32) {
	for {
		e.update()
		stop := false
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			stop = e.handleKeyBoardPress(event)
		}
		if stop {
			return
		}
		sdl.Delay(1000 / fps)
	}
}
