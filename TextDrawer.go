package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func drawText(renderer *sdl.Renderer, text string, x, y int32, angle float64, color sdl.Color) (sdl.Rect, error) {
	// Load the font
	font, err := ttf.OpenFont("/usr/share/fonts/opentype/unifont/unifont.otf", 24)
	if err != nil {
		return sdl.Rect{0, 0, 0, 0}, fmt.Errorf("failed to load font: %v", err)
	}
	defer font.Close()

	// Create a surface with the text
	surface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		return sdl.Rect{0, 0, 0, 0}, fmt.Errorf("failed to render text: %v", err)
	}
	defer surface.Free()

	// Create a texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return sdl.Rect{0, 0, 0, 0}, fmt.Errorf("failed to create texture: %v", err)
	}
	defer texture.Destroy()

	// Calculate the position and size of the text
	rect := sdl.Rect{X: x, Y: y, W: surface.W, H: surface.H}

	renderer.CopyEx(texture, nil, &rect, angle, &sdl.Point{X: surface.W / 2, Y: surface.H / 2}, sdl.FLIP_NONE)

	// Copy the texture to the rendering target
	// err = renderer.Copy(texture, nil, &rect)
	// if err != nil {
	// 	return sdl.Rect{0, 0, 0, 0}, fmt.Errorf("failed to copy texture: %v", err)
	// }

	return rect, nil
}
