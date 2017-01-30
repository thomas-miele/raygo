package raytracer

import (
	//"math"
	"image"
	"image/color"
)

const D = 100
const WinX = 720
const WinY = 480

func Raytracer(scene *Scene, img *image.RGBA) {
	var c color.RGBA

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c.R = 63
			c.G = 63
			c.B = 255
			c.A = 255
			//c = initCalc(scene, x, y)
			img.Set(x, y, c)
		}
	}
}

/*
func initCalc(scene *Scene, x, y int) color.RGBA {
	var ray Ray
	var x1, y1, z1 float32
	var pixel color.RGBA

	pixel.R = 255
	pixel.G = 255
	pixel.B = 255
	pixel.A = 255

	x1 = float32(D)
	y1 = float32((WinX / 2) - x)
	z1 = float32((WinY / 2) - y)

	ray.vx = x1 - scene.cam.pos.x
	ray.vy = y1 - scene.cam.pos.y
	ray.vz = z1 - scene.cam.pos.z

	ray.topMesh = scene.meshs
	ray.topSpot = scene.lights

	pixel = calc(scene, &ray, x, y)
	return pixel
}

func calc(scene *Scene, ray *Ray, x, y int) color.RGBA {
	var k float32 = math.MaxFloat32
	var pixel color.RGBA
	var ret int
	var tmp *mesh = nil
	var ray Ray

	for tmp != nil{
		ret = interChoice(scene, tmp, ray)
		if (ret == 1) {

		}
	}
}
*/
