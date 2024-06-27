package pkg

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GameEngine struct {
	GameObjects *[]GameObject3D
	Window      *sdl.Window
	Renderer    *sdl.Renderer
	Camera      *Camera
}

func (g *GameEngine) Close() {
	// выполнить очистку или освобождение ресурсов
	g.Renderer.Destroy()
	g.Window.Destroy()
	sdl.Quit()
}

func InitGameEngine(gameObjects *[]GameObject3D, width, heigh int32, fov float64) (*GameEngine, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println("Error initializing SDL:", err)
		return nil, err
	}

	Window, err := sdl.CreateWindow("Cube", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, heigh, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error creating Window:", err)
		sdl.Quit()
		return nil, err
	}

	Renderer, err := sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error creating Renderer:", err)
		Window.Destroy()
		sdl.Quit()
		return nil, err
	}
	if err := ttf.Init(); err != nil {
		fmt.Println("Error initializing SDL_ttf:", err)
		Renderer.Destroy()
		Window.Destroy()
		sdl.Quit()
		return nil, err

	}
	if gameObjects == nil {
		gObjs := make([]GameObject3D, 0)
		gameObjects = &gObjs
	}

	fovRadV, fovRadH := getFovVH(float64(width), float64(heigh), fov)

	camera := Camera{Vector3D{0, 0, -10},
		Vector3D{0, 0, 0},
		fovRadV,
		fovRadH,
	}

	return &GameEngine{gameObjects, Window, Renderer, &camera}, nil

}

func getFovVH(width, heigh, fov float64) (float64, float64) {
	aspectRatio := float64(width) / float64(heigh)
	fovRadV := 1.0 / math.Tan(fov*0.5/180.0*math.Pi)
	fovRadH := fovRadV / aspectRatio
	return fovRadV, fovRadH
}

type GameObject3DAbs interface {
	Update()
}

func (e *GameEngine) Update() {
	fmt.Println("Original update has been called")

}

func (e *GameEngine) GetGameObject() *GameObject3D {
	for _, gObg := range *e.GameObjects {
		return &gObg
	}
	return nil

}

func (e GameEngine) AddGameObj(gameObject GameObject3D) {
	*e.GameObjects = append(*e.GameObjects, gameObject)

}

type GameObjectInterface interface {
	DrawObjects()
	AddGameObj(GameObject3D)
	Update()
	HandleKeyBoardPress(event sdl.Event) bool
}

func (e GameEngine) DrawObjects() {
	e.Camera.DrawObjects(e.Renderer, e.GameObjects)

}
func EngineUpdate(e GameObjectInterface) {
	e.Update()
	e.DrawObjects()

}

func Run(fps *uint32, e GameObjectInterface) {
	for {
		EngineUpdate(e)
		stop := false
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			stop = e.HandleKeyBoardPress(event)
		}
		if stop {
			return
		}
		sdl.Delay(1000 / *fps)
	}
}
