package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func drawUI(renderer *sdl.Renderer, angleVelocityObj *AngleVelocity) error {
	var x int32 = 10
	var y int32 = 10
	var offset int32 = 5

	rect, err := drawText(renderer, fmt.Sprintf("FPS: %d", FPS), x, y, BLACK)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset
	rect, err = drawText(renderer, fmt.Sprintf("Velocity: %d", SPEED), x, y, BLACK)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset
	rect, err = drawText(
		renderer,
		fmt.Sprintf("Angle X: %.2f, Y: %.2f, Z: %.2f", angleVelocityObj.angleX, angleVelocityObj.angleY, angleVelocityObj.angleZ),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}

	y = rect.Y + rect.H + offset
	rect, err = drawText(
		renderer,
		fmt.Sprintf("Angle Velocity: %.2f, Y: %.2f, Z: %.2f", angleVelocityObj.angleXVel, angleVelocityObj.angleYVel, angleVelocityObj.angleZVel),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}
	y = rect.Y + rect.H + offset

	currentAx := angleVelocityObj.getAxName()

	rect, err = drawText(
		renderer,
		fmt.Sprintf("currentAx: %s", currentAx),
		x,
		y,
		BLACK,
	)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return fmt.Errorf("failed to drawUI: %v", err)
	}

	return nil
}
