package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector3D struct {
	X, Y, Z float64
}
type Vector4D struct {
	X, Y, Z, W float64
}

type AngleVelocity struct {
	angleX, angleY, angleZ          float64
	angleXVel, angleYVel, angleZVel float64
	currentAxeIdx                   int
}

func (aV *AngleVelocity) updateAngles() {
	aV.angleX = math.Mod(aV.angleX+aV.angleXVel, 360)
	aV.angleY = math.Mod(aV.angleY+aV.angleYVel, 360)
	aV.angleZ = math.Mod(aV.angleZ+aV.angleZVel, 360)
}
func (aV *AngleVelocity) changeAxe() {
	aV.currentAxeIdx = (aV.currentAxeIdx + 1) % len(axes)

}

func (aV *AngleVelocity) getAxName() string {
	switch axes[aV.currentAxeIdx] {
	case xAxe:
		return "X"
	case yAxe:
		return "Y"
	case zAxe:
		return "Z"
	}
	return ""
}

func (aV *AngleVelocity) changeAngleVelociity(up bool) {
	var adder float64 = 1
	if !up {
		adder = -1
	}
	axe := aV.currentAxeIdx

	switch axe {
	case xAxe:
		aV.angleXVel += adder
	case yAxe:
		aV.angleYVel += adder
	case zAxe:
		aV.angleZVel += adder
	}
}

type GameObject3D struct {
	vertices        *[]Vector3D
	updatedVertices []Vector3D
	edges           *[][]int
	angles          *AngleVelocity
	position        *Vector3D
	rotation        *Vector3D
}

func (gameObj *GameObject3D) draw(renderer *sdl.Renderer) {
	for _, edge := range *gameObj.edges {
		p1 := gameObj.updatedVertices[edge[0]]
		p2 := gameObj.updatedVertices[edge[1]]

		if p1.X != p2.X || p1.Y != p2.Y {
			renderer.DrawLine(int32(p1.X), int32(p1.Y), int32(p2.X), int32(p2.Y))
		}

	}
}

type Matrix4x4 [4][4]float64

func (m1 Matrix4x4) Multiply(m2 Matrix4x4) Matrix4x4 {
	result := Matrix4x4{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}

func (m Matrix4x4) MultiplyVector(v Vector4D) Vector4D {
	result := Vector4D{}

	result.X = m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W
	result.Y = m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W
	result.Z = m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W
	result.W = m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W

	return result
}
