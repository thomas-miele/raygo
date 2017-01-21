package main

import (
	"fmt"
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
			c = calc(x, y)
			img.Set(x, y, c)
		}
	}
}

func calc(x, y int) color.RGBA {
	return init_calc(x, y)
}

func init_calc(x, y int) color.RGBA {
	//t_ray ray
	var x1, y1, z1 float32
	var pixel color.RGBA

	pixel.R = 255
	pixel.A = 255

	x1 = float32(D)
	y1 = float32((WinX / 2) - x)
	z1 = float32((WinY / 2) - y)

	fmt.Println(x1, y1, z1)
	/*
	ray.Vx = x1 - eye->pos.X;
	ray.Vy = y1 - eye->pos.Y;
	ray.Vz = z1 - eye->pos.Z;

	ray.top_mesh = llist;
	ray.top_spot = spot;
	ray.bool = 1;
	rotate_x(&ray.Vx, &ray.Vy, &ray.Vz, -eye->rot.X);
	rotate_y(&ray.Vx, &ray.Vy, &ray.Vz, -eye->rot.Y);
	rotate_z(&ray.Vx, &ray.Vy, &ray.Vz, -eye->rot.Z);
	color = calc(eye, llist, spot, &ray);
        */
	return pixel
}
