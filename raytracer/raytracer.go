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
			c = initCalc(scene, x, y)
			img.Set(x, y, c)
		}
	}
}

func initCalc(scene *Scene, x, y int) color.RGBA {
	var ray Ray
	var x1, y1, z1 float32
	var pixel color.RGBA

	pixel.R = 63
	pixel.G = 63
	pixel.B = 255
	pixel.A = 255

	x1 = float32(D)
	y1 = float32((WinX / 2) - x)
	z1 = float32((WinY / 2) - y)

	ray.V.X = x1 - scene.Cam.Pos.X
	ray.V.Y = y1 - scene.Cam.Pos.Y
	ray.V.Z = z1 - scene.Cam.Pos.Z

	// ray.TopMesh = scene.Meshs
	// ray.TopSpot = scene.Lights
	ray.Is = true

	// pixel = calc(scene, &ray, x, y)
	return pixel
}

// func calc(scene *Scene, ray *Ray, x, y int) color.RGBA {
// 	var k float32 = math.MaxFloat32
// 	var pixel color.RGBA

// 	var ret int
// 	var tmp *mesh = nil

// 	for tmp != nil {
// 		ret = interChoice(scene, tmp, ray)
// 		if (ret == 1) {
// 			if (ray.K <= k) {
// 				k = ray.K
// 				pixel = luminosite()
// 			}
// 		}
// 		tmp = scene.Meshs[0]
// 	}
// 	return pixel
// }
