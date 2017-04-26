package raytracer

import (
	"image/color"
)

type MeshType int
type LightType int

const (
	Plane MeshType = iota
	Sphere
	Cylinder
	Cone
)

const (
	Spot LightType = iota
)

type Pixel struct {
	Color  color.RGBA
	Bright float32
}

type Camera struct {
	Pos Vector
	Rot Vector
}

type Scene struct {
	Width  int
	Height int
	D      int
	Cam    Camera
	Meshs  []Mesh
	Lights []Light
}

type Mesh struct {
	Pos   Vector
	Rot   Vector
	Color Pixel
	R     float32
}

type Light struct {
	Pos   Vector
	Color Pixel
}

type Ray struct {
	Is      bool
	V       Vector
	K       float32
	Color   Pixel
	TopMesh *Mesh
	TopSpot *Light
}

type CalcRes struct {
	A, B, C float32
	Delta   float32
	K1, K2  float32
}
