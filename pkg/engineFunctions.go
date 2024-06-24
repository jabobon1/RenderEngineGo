package pkg

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

func Magnitude(v Vector3D) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
func Normalize(v Vector3D) Vector3D {
	mag := Magnitude(v)
	return Vector3D{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag}
}

func Add(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}
func Add2D(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}
func Mult(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}
func Mult2D(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X * v2.X, v1.Y * v2.Y}
}

func Sub(v1, v2 Vector3D) Vector3D {
	return Vector3D{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z}
}

func Sub2D(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

func Cross(v1, v2 Vector3D) Vector3D {
	return Vector3D{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
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

type Line struct {
	vec1, vec2 Vector3D
}

func (line Line) Cross(v1, v2 Vector3D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X

}

func getIntersectionVector(line1, line2 Line) (Vector3D, bool) {
	sub1 := Sub(line1.vec2, line1.vec1)
	sub2 := Sub(line2.vec2, line2.vec1)

	cross_prod := line1.Cross(sub1, sub2)
	if cross_prod == 0 {
		return Vector3D{0, 0, 0}, false
	}
	vec1Sub := Sub(line2.vec1, line1.vec1)
	crossSub1 := line1.Cross(vec1Sub, sub2)

	t := crossSub1 / cross_prod

	interesctionVector := Add(
		line1.vec1,
		Mult(line1.vec2, Vector3D{t, t, t}), // multiplying by time
	)

	return interesctionVector, true
}

func Centroid(p1, p2, p3 Vector3D) Vector3D {
	return Vector3D{
		X: (p1.X + p2.X + p3.X) / 3,
		Y: (p1.Y + p2.Y + p3.Y) / 3,
		Z: (p1.Z + p2.Z + p3.Z) / 3,
	}
}
