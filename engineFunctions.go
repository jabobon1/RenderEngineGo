package main

import (
	"math"
)

func Distance(v1 Vector3D, v2 Vector3D) float64 {
	dx := v2.X - v1.X
	dy := v2.Y - v1.Y
	dz := v2.Z - v1.Z

	return math.Sqrt(
		math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2))
}
func RadiansToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

func DegreesToRadians(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func AngleBetween(v1, v2 Vector3D) Vector3D {
	angle_xy := math.Atan2(v2.Y, v2.X) - math.Atan2(v1.Y, v1.X)
	angle_yz := math.Atan2(v2.Z, math.Sqrt(v2.X*v2.X+v2.Y*v2.Y)) - math.Atan2(v1.Z, math.Sqrt(v1.X*v1.X+v1.Y*v1.Y))
	angle_xz := math.Atan2(v2.Z, v2.X) - math.Atan2(v1.Z, v1.X)

	return Vector3D{
		RadiansToDegrees(angle_xy),
		RadiansToDegrees(angle_yz),
		RadiansToDegrees(angle_xz),
	}
}

func Normalize(v Vector3D) Vector3D {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vector3D{v.X / length, v.Y / length, v.Z / length}
}

func Add(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Sub(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func Cross(v1, v2 Vector3D) Vector3D {
	x := v1.Y*v2.Z - v1.Z*v2.Y
	y := v1.Z*v2.X - v1.X*v2.Z
	z := v1.X*v2.Y - v1.Y*v2.X
	return Vector3D{x, y, z}
}

func Dot(v1, v2 Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func XRotationMatrix(angle float64) Matrix4x4 {
	rad := DegreesToRadians(angle)
	sin, cos := math.Sin(rad), math.Cos(rad)
	return Matrix4x4{
		{1, 0, 0, 0},
		{0, cos, -sin, 0},
		{0, sin, cos, 0},
		{0, 0, 0, 1},
	}
}

func YRotationMatrix(angle float64) Matrix4x4 {
	rad := DegreesToRadians(angle)
	sin, cos := math.Sin(rad), math.Cos(rad)
	return Matrix4x4{
		{cos, 0, sin, 0},
		{0, 1, 0, 0},
		{-sin, 0, cos, 0},
		{0, 0, 0, 1},
	}
}
func ZRotationMatrix(angle float64) Matrix4x4 {
	rad := DegreesToRadians(angle)
	sin, cos := math.Sin(rad), math.Cos(rad)

	return Matrix4x4{
		{cos, -sin, 0, 0},
		{sin, cos, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func XYZRotationMatrix(aX, aY, aZ float64) Matrix4x4 {
	x := DegreesToRadians(aX)
	y := DegreesToRadians(aY)
	z := DegreesToRadians(aZ)

	xSin, xCos := math.Sin(x), math.Cos(x)
	ySin, yCos := math.Sin(y), math.Cos(y)
	zSin, zCos := math.Sin(z), math.Cos(z)

	return Matrix4x4{
		{yCos * zCos, xSin*ySin*zCos - xCos*zSin, xCos*ySin*zCos + xSin*zSin, 0},
		{yCos * zSin, xSin*ySin*zSin + xCos*zCos, xCos*ySin*zSin - xSin*zCos, 0},
		{-ySin, xSin * yCos, xCos * yCos, 0},
		{0, 0, 0, 1},
	}
}
