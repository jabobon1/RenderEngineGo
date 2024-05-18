package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Point3D struct {
	X, Y, Z float64
}

func (point *Point3D) rotateX(angle float64) {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)

	y := point.Y*cos - point.Z*sin
	z := point.Y*sin + point.Z*cos

	point.Y = y
	point.Z = z

}

func (point *Point3D) rotateY(angle float64) {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)

	x := point.X*cos + point.Z*sin
	z := -point.X*sin + point.Z*cos

	point.X = x
	point.Z = z

}

func (point *Point3D) rotateZ(angle float64) {
	rad := angle * math.Pi / 180.0
	sin, cos := math.Sin(rad), math.Cos(rad)

	x := point.X*cos - point.Y*sin
	y := point.X*sin + point.Y*cos

	point.X = x
	point.Y = y

}

func (point *Point3D) rotate(x, y, z float64) {
	point.rotateX(x)
	point.rotateY(y)
	point.rotateZ(z)

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
	vertices        *[]Point3D
	updatedVertices []Point3D
	edges           *[][]int
	angles          *AngleVelocity
	position        *Point3D
}

func (gameObj *GameObject3D) rotate() {
	gameObj.angles.updateAngles()
	// Rotate vertices
	for i, vertex := range *gameObj.vertices {
		vertex.rotate(
			gameObj.angles.angleX,
			gameObj.angles.angleY,
			gameObj.angles.angleZ,
		)
		gameObj.updatedVertices[i] = vertex
	}
}
func (gameObj *GameObject3D) draw(renderer *sdl.Renderer) {
	for _, edge := range *gameObj.edges {
		renderer.DrawLine(
			int32(gameObj.updatedVertices[edge[0]].X+gameObj.position.X),
			int32(gameObj.updatedVertices[edge[0]].Y+gameObj.position.Y),
			int32(gameObj.updatedVertices[edge[1]].X+gameObj.position.X),
			int32(gameObj.updatedVertices[edge[1]].Y+gameObj.position.Y),
		)
	}
}
