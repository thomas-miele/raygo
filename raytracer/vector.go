package raytracer

import "math"

// Point is 3D coordinates position
type Point struct {
	X, Y, Z float64
}

// Vector interface for 3D coordinates and 3D vector
type Vector Point

// RayLight is a ray of light
type RayLight struct {
	Position  Point
	Direction Point
}

// Coordinates for toto foobar
type Coordinates interface {
	plus(right Point) Point
	minus(right Point) Point
	cross(right Point) Point
	dot(right Point) float64
	times(k float64) Point
	mag() float64
	norm() Point
}

// plus Addition de deux coordonnées 3D par axe
func (left Point) plus(right Point) Point {
	return Point{left.X + right.X, left.Y + right.Y, left.Z + right.Z}
}

// minus Soustraction de deux coordonnées 3D par axe
func (left Point) minus(right Point) Point {
	return Point{left.X - right.X, left.Y - right.Y, left.Z - right.Z}
}

func (left Point) cross(right Point) Point {
	x := (left.Y * right.Z) - (left.Z * right.Y)
	y := (left.Z * right.X) - (left.X * right.Z)
	z := (left.X * right.Y) - (left.Y * right.X)
	return Point{x, y, z}
}

// dot Addition des multiplications de 2 coordonnées 3D par axe
func (left Point) dot(right Point) float64 {
	return (left.X * right.X) + (left.Y * right.Y) + (left.Z * right.Z)
}

// times Multiplication par k d'une coordonnée 3D par axe
func (left Point) times(k float64) Point {
	return Point{k * left.X, k * left.Y, k * left.Z}
}

// mag Racine carré de la somme des multiplications de 2 coordonées 3D par axe
func (left Point) mag() float64 {
	formula := (left.X * left.X) + (left.Y * left.Y) + (left.Z * left.Z)
	sqrt := math.Sqrt(formula)
	return sqrt
}

// norm Normaliser une coordonée 3D
func (left Point) norm() Point {
	var vmag = left.mag()
	var div float64

	if vmag == 0.0 {
		div = math.MaxFloat64
	} else {
		div = 1 * vmag
	}
	return left.times(div)
}
