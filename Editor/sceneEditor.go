package main

import (
	"fmt"
	"renderEngineGo/pkg"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	FONT_SIZE = 24
)

type Button struct {
	Rect       sdl.Rect
	Text       string
	Color      sdl.Color
	HoverColor sdl.Color
	TextColor  sdl.Color
	IsHovered  bool
}

type SceneEditor struct {
	Renderer         *sdl.Renderer
	Window           *sdl.Window
	Buttons          []Button
	Objects          []pkg.GameObject3D
	Font             *ttf.Font
	AvailableFigures map[string]func(pkg.Vector3D) pkg.GameObject3D
	ShowFiguresList  bool
	AddGameObj       func(pkg.GameObject3D)
	SelectedObject   *pkg.GameObject3D
	EditMode         string // "position", "rotation", "scale"
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
			Color:      sdl.Color{R: 70, G: 130, B: 180, A: 255},
			HoverColor: sdl.Color{R: 100, G: 149, B: 237, A: 255},
			TextColor:  sdl.Color{R: 255, G: 255, B: 255, A: 255},
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
		shadowRect := sdl.Rect{X: button.Rect.X + 2, Y: button.Rect.Y + 2, W: button.Rect.W, H: button.Rect.H}
		e.Renderer.SetDrawColor(0, 0, 0, 100)
		e.Renderer.FillRect(&shadowRect)

		if button.IsHovered {
			e.Renderer.SetDrawColor(button.HoverColor.R, button.HoverColor.G, button.HoverColor.B, button.HoverColor.A)
		} else {
			e.Renderer.SetDrawColor(button.Color.R, button.Color.G, button.Color.B, button.Color.A)
		}
		e.Renderer.FillRect(&button.Rect)

		e.Renderer.SetDrawColor(255, 255, 255, 255)
		e.Renderer.DrawRect(&button.Rect)

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

	e.Renderer.SetDrawColor(240, 240, 240, 255)
	e.Renderer.FillRect(&sdl.Rect{X: listX, Y: listY, W: listWidth, H: listHeight})

	e.Renderer.SetDrawColor(70, 130, 180, 255)
	e.Renderer.DrawRect(&sdl.Rect{X: listX, Y: listY, W: listWidth, H: listHeight})

	i := 0
	for figure := range e.AvailableFigures {
		e.drawText(figure, listX+10, listY+int32(i*30)+10, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		i++
	}
}

func (e *SceneEditor) handleMouseMotion(event *sdl.MouseMotionEvent) {
	for i := range e.Buttons {
		e.Buttons[i].IsHovered = event.X >= e.Buttons[i].Rect.X && event.X < e.Buttons[i].Rect.X+e.Buttons[i].Rect.W &&
			event.Y >= e.Buttons[i].Rect.Y && event.Y < e.Buttons[i].Rect.Y+e.Buttons[i].Rect.H
	}

	if e.SelectedObject != nil && e.EditMode != "" {
		fmt.Println("selected", e.EditMode)

		dx := float64(event.XRel) * 0.01
		dy := float64(event.YRel) * 0.01

		switch e.EditMode {
		case "position":
			e.SelectedObject.Position.X += dx
			e.SelectedObject.Position.Y -= dy
		case "rotation":
			e.SelectedObject.Rotation.X += dy
			e.SelectedObject.Rotation.Y += dx
		case "scale":
			e.SelectedObject.Size.X += dx
			e.SelectedObject.Size.Y += dy
			e.SelectedObject.Size.Z += (dx + dy) / 2
		}
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
					figureKeys := make([]string, 0, len(e.AvailableFigures))
					for k := range e.AvailableFigures {
						figureKeys = append(figureKeys, k)
					}
					getFigure := e.AvailableFigures[figureKeys[selectedIndex]]
					newObject := getFigure(pkg.Vector3D{1, 1, 1})
					e.AddGameObj(newObject)
					e.Objects = append(e.Objects, newObject)
					e.ShowFiguresList = false
				}
				return
			}
		}

		e.ShowFiguresList = false
		// Object selection
		for i := range e.Objects {
			obj := e.Objects[i]
			minX, maxX, minY, maxY := obj.GetMinMaxPointsOnScreen()
			fmt.Println(minX, maxX, minY, maxY, event.X, event.Y)

			// Implement object selection logic here
			// This is a simplified example; you may need to implement proper 3D picking
			if event.X >= int32(minX) && event.X < int32(maxX) &&
				event.Y >= int32(minY) && event.Y < int32(maxY) {
				e.SelectedObject = &e.Objects[i]
				fmt.Println("SELECTED", e.Objects[i])
				return
			}
		}

		e.SelectedObject = nil
	}
}

func (e *SceneEditor) handleButtonClick(buttonIndex int) {
	switch buttonIndex {
	case 0: // Add Object
		e.ShowFiguresList = !e.ShowFiguresList
		e.EditMode = ""
	case 1: // Change Position
		e.EditMode = "position"
	case 2: // Change Rotation
		e.EditMode = "rotation"
	case 3: // Change Scale
		e.EditMode = "scale"
	}
}

func (e *SceneEditor) DrawScene() {
	e.Renderer.SetDrawColor(245, 245, 245, 255)
	e.drawButtons()
	e.drawFiguresList()

	// Draw selection outline for the selected object
	if e.SelectedObject != nil {
		e.Renderer.SetDrawColor(255, 255, 0, 255) // Yellow outline
		rect := sdl.Rect{
			X: int32(e.SelectedObject.Position.X) - 2,
			Y: int32(e.SelectedObject.Position.Y) - 2,
			W: int32(e.SelectedObject.Size.X) + 4,
			H: int32(e.SelectedObject.Size.Y) + 4,
		}
		e.Renderer.DrawRect(&rect)
	}
}

func (e *SceneEditor) Cleanup() {
	e.Font.Close()
}
