package types

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

type Vector3i struct {
	X, Y, Z int
}

type Position Vector3i
type Rotation Vector3i

type Color struct {
	R, G, B int8
	Bright  float32
}

type Camera struct {
	Pos Position
	Rot Rotation
}

type Scene struct {
	Width, Height int
	D             int
	Cam           Camera
	Meshs         []Mesh
	Lights        []Light
}

type Mesh struct {
	Pos   Position
	Rot   Rotation
	Color Color
	R     float32
}

type Light struct {
	Pos   Position
	Color Color
}
