package main

import "math"

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
