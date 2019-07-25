package raytracer

import (
	//"math"
	"image"
	"image/color"
)

const D = 100
const WinX = 720
const WinY = 480

type Raytracer struct {
	Image *image.RGBA
	Scene *Scene
}

func (self *Raytracer) Algo() {
	var pixel color.RGBA

	bounds := self.Image.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel = self.Pixel(x, y)
			self.Image.Set(x, y, pixel)
		}
	}

}

func (self *Raytracer) Pixel(x int, y int) color.RGBA {
	var ray Ray
	var x1, y1, z1 float32
	var pixel color.RGBA

	pixel.R = 255
	pixel.G = 99
	pixel.B = 71
	pixel.A = 255

	x1 = float32(D)
	y1 = float32((WinX / 2) - x)
	z1 = float32((WinY / 2) - y)

	ray.V.X = x1 - self.Scene.Cam.Pos.X
	ray.V.Y = y1 - self.Scene.Cam.Pos.Y
	ray.V.Z = z1 - self.Scene.Cam.Pos.Z

	// ray.TopMesh = scene.Meshs
	// ray.TopSpot = scene.Lights
	ray.Is = true

	//pixel = self.Calcul(&ray, x, y)
	return pixel
}

// func (self *Raytracer) Calcul(ray *Ray, x, y int) color.RGBA {
// 	var k float32 = math.MaxFloat32
// 	var pixel color.RGBA

// 	var ret int
// 	var tmp *mesh = nil

// 	for tmp != nil {
// 		ret = interChoice(self.scene, tmp, ray)
// 		if ret == 1 {
// 			if ray.K <= k {
// 				k = ray.K
// 				pixel = luminosite()
// 			}
// 		}
// 		tmp = scene.Meshs[0]
// 	}
// 	return pixel
// }
