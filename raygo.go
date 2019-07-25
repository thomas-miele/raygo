package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/tmiele/raygo/raytracer"
)

func main() {
	var scene raytracer.Scene

	scene.Width = 640
	scene.Height = 480
	scene.D = 100
	scene.Cam.Pos.X = -300
	scene.Cam.Pos.Y = 50

	scene.Meshs = append(scene.Meshs, raytracer.Mesh{})

	// RGBA image use for pixel by pixel raytracer
	imgRect := image.Rect(0, 0, scene.Width, scene.Height)
	img := image.NewRGBA(imgRect)

	raytracer.Raytracer(&scene, img)

	// encoding into png buffer
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, buffer)
}
