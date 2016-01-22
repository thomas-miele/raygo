package main

import (
	"image/color"
)

type meshType int
type lightType int

const (
	plane meshType = iota
	sphere
	cylinder
	cone
)

const (
	spot lightType = iota
)

type vector struct {
	x, y, z float32
}

type rotation vector
type position vector

type pixel struct {
	color  color.RGBA
	bright float32
}

type camera struct {
	pos position
	rot rotation
}

type Scene struct {
	width  int
	height int
	d      int
	cam    camera
	meshs  []mesh
	lights []light
}

type mesh struct {
	pos   position
	rot   rotation
	color pixel
	R     float32
}

type light struct {
	pos   position
	color pixel
}

type Ray struct {
	is         bool
	vx, vy, vz float32
	k          float32
	color      pixel
	topMesh    *mesh
	topSpot    *light
}

type calcRes struct {
	a, b, c float32
	delta   float32
	k1, k2  float32
}
