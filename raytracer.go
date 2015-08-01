package main

import (
	"image"
	"image/color"
)

func Raytracer(scene *Scene, img *image.RGBA) {
	var c color.Color

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c = calc()
			img.Set(x, y, c)
		}
	}
}

func calc() color.Color {
	return color.RGBA{255, 0, 0, 255}
}

func init_calc() color.Color {
	return color.RGBA{255, 0, 0, 255}
}
