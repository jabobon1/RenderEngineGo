package main

import (
	"fmt"
	"renderEngineGo/pkg"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	FONT_SIZE = 24 // Increased font size

)

type Button struct {
	Rect       sdl.Rect
	Text       string
	Color      sdl.Color
	HoverColor sdl.Color
	TextColor  sdl.Color
	IsHovered  bool
}

type SceneObject struct {
	Type     string
	Position pkg.Vector3D
	Rotation pkg.Vector3D
	Scale    pkg.Vector3D
}

type SceneEditor struct {
	Renderer         *sdl.Renderer
	Window           *sdl.Window
	Buttons          []Button
	Objects          []SceneObject
	Font             *ttf.Font
	AvailableFigures []string
	ShowFiguresList  bool
}

func NewSceneEditor() (*SceneEditor, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, fmt.Errorf("failed to initialize SDL: %v", err)
	}

	if err := ttf.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize TTF: %v", err)
	}

	window, err := sdl.CreateWindow("3D Scene Editor", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, fmt.Errorf("failed to create renderer: %v", err)
	}

	font, err := ttf.OpenFont("/usr/share/fonts/opentype/unifont/unifont.otf", FONT_SIZE)
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	editor := &SceneEditor{
		Renderer:         renderer,
		Window:           window,
		Font:             font,
		AvailableFigures: []string{"Cube", "Sphere", "Pyramid", "Cylinder"},
	}

	editor.initializeButtons()

	return editor, nil
}

func (e *SceneEditor) initializeButtons() {
	buttonTexts := []string{"Add Object", "Change Position", "Change Rotation", "Change Scale"}
	buttonWidth := int32(200)
	buttonHeight := int32(60)
	spacing := int32(10)
	totalWidth := int32(len(buttonTexts))*(buttonWidth+spacing) - spacing
	startX := WIDTH - totalWidth - spacing

	for i, text := range buttonTexts {
		e.Buttons = append(e.Buttons, Button{
			Rect:       sdl.Rect{X: startX + int32(i)*(buttonWidth+spacing), Y: spacing, W: buttonWidth, H: buttonHeight},
			Text:       text,
			Color:      sdl.Color{R: 70, G: 130, B: 180, A: 255},  // Steel Blue
			HoverColor: sdl.Color{R: 100, G: 149, B: 237, A: 255}, // Cornflower Blue
			TextColor:  sdl.Color{R: 255, G: 255, B: 255, A: 255}, // White
		})
	}
}

func (e *SceneEditor) drawText(text string, x, y int32, color sdl.Color) error {
	surface, err := e.Font.RenderUTF8Blended(text, color)
	if err != nil {
		return fmt.Errorf("failed to render text: %v", err)
	}
	defer surface.Free()

	texture, err := e.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("failed to create texture: %v", err)
	}
	defer texture.Destroy()

	rect := sdl.Rect{X: x, Y: y, W: surface.W, H: surface.H}
	e.Renderer.Copy(texture, nil, &rect)

	return nil
}

func (e *SceneEditor) drawButtons() {
	for _, button := range e.Buttons {
		// Draw button shadow
		shadowRect := sdl.Rect{X: button.Rect.X + 2, Y: button.Rect.Y + 2, W: button.Rect.W, H: button.Rect.H}
		e.Renderer.SetDrawColor(0, 0, 0, 100)
		e.Renderer.FillRect(&shadowRect)

		// Draw button
		if button.IsHovered {
			e.Renderer.SetDrawColor(button.HoverColor.R, button.HoverColor.G, button.HoverColor.B, button.HoverColor.A)
		} else {
			e.Renderer.SetDrawColor(button.Color.R, button.Color.G, button.Color.B, button.Color.A)
		}
		e.Renderer.FillRect(&button.Rect)

		// Draw button border
		e.Renderer.SetDrawColor(255, 255, 255, 255)
		e.Renderer.DrawRect(&button.Rect)

		// Draw button text
		textWidth, _, _ := e.Font.SizeUTF8(button.Text)
		textX := button.Rect.X + (button.Rect.W-int32(textWidth))/2
		textY := button.Rect.Y + (button.Rect.H-int32(e.Font.Height()))/2
		e.drawText(button.Text, textX, textY, button.TextColor)
	}
}

func (e *SceneEditor) drawFiguresList() {
	if !e.ShowFiguresList {
		return
	}

	listX := e.Buttons[0].Rect.X
	listY := e.Buttons[0].Rect.Y + e.Buttons[0].Rect.H + 5
	listWidth := e.Buttons[0].Rect.W
	listHeight := int32(len(e.AvailableFigures)*30 + 10)

	// Draw list background
	e.Renderer.SetDrawColor(240, 240, 240, 255)
	e.Renderer.FillRect(&sdl.Rect{X: listX, Y: listY, W: listWidth, H: listHeight})

	// Draw list border
	e.Renderer.SetDrawColor(70, 130, 180, 255)
	e.Renderer.DrawRect(&sdl.Rect{X: listX, Y: listY, W: listWidth, H: listHeight})

	for i, figure := range e.AvailableFigures {
		e.drawText(figure, listX+10, listY+int32(i*30)+10, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	}
}

func (e *SceneEditor) handleEvents() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.MouseMotionEvent:
			e.handleMouseMotion(t)
		case *sdl.MouseButtonEvent:
			e.handleMouseClick(t)
		}
	}
	return true
}

func (e *SceneEditor) handleMouseMotion(event *sdl.MouseMotionEvent) {
	for i := range e.Buttons {
		e.Buttons[i].IsHovered = event.X >= e.Buttons[i].Rect.X && event.X < e.Buttons[i].Rect.X+e.Buttons[i].Rect.W &&
			event.Y >= e.Buttons[i].Rect.Y && event.Y < e.Buttons[i].Rect.Y+e.Buttons[i].Rect.H
	}
}

func (e *SceneEditor) handleMouseClick(event *sdl.MouseButtonEvent) {
	if event.Type == sdl.MOUSEBUTTONDOWN {
		for i, button := range e.Buttons {
			if event.X >= button.Rect.X && event.X < button.Rect.X+button.Rect.W &&
				event.Y >= button.Rect.Y && event.Y < button.Rect.Y+button.Rect.H {
				e.handleButtonClick(i)
				return
			}
		}

		if e.ShowFiguresList {
			listX := e.Buttons[0].Rect.X
			listY := e.Buttons[0].Rect.Y + e.Buttons[0].Rect.H + 5
			listWidth := e.Buttons[0].Rect.W
			listHeight := int32(len(e.AvailableFigures)*30 + 10)

			if event.X >= listX && event.X < listX+listWidth &&
				event.Y >= listY && event.Y < listY+listHeight {
				selectedIndex := (event.Y - listY - 5) / 30
				if selectedIndex >= 0 && int(selectedIndex) < len(e.AvailableFigures) {
					e.addObject(e.AvailableFigures[selectedIndex])
					e.ShowFiguresList = false
				}
				return
			}
		}

		e.ShowFiguresList = false
	}
}

func (e *SceneEditor) handleButtonClick(buttonIndex int) {
	switch buttonIndex {
	case 0: // Add Object
		e.ShowFiguresList = !e.ShowFiguresList
	case 1:
		// Implement change position logic
	case 2:
		// Implement change rotation logic
	case 3:
		// Implement change scale logic
	}
}

func (e *SceneEditor) addObject(objectType string) {
	newObject := SceneObject{
		Type:     objectType,
		Position: pkg.Vector3D{X: 0, Y: 0, Z: 0},
		Rotation: pkg.Vector3D{X: 0, Y: 0, Z: 0},
		Scale:    pkg.Vector3D{X: 1, Y: 1, Z: 1},
	}
	e.Objects = append(e.Objects, newObject)
	fmt.Printf("Added %s to the scene\n", objectType)
}

func (e *SceneEditor) Run() {
	running := true
	for running {
		running = e.handleEvents()

		e.Renderer.SetDrawColor(245, 245, 245, 255) // Light gray background
		e.Renderer.Clear()

		e.drawButtons()
		e.drawFiguresList()

		e.Renderer.Present()
	}
}

func (e *SceneEditor) Cleanup() {
	e.Font.Close()
	e.Renderer.Destroy()
	e.Window.Destroy()
	ttf.Quit()
	sdl.Quit()
}
