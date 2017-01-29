package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/thomas-miele/raygo/raytracer"
)

func usage() {
	fmt.Println("raygo [JSON_SCENE] [-o OUTFILE]")
}

func main() {
	var scene raytracer.Scene
	var path string = "out.png"

	// si arguments -> utiliser une scene json
	if len(os.Args) > 1 {
		var sceneFile string
		flag.StringVar(&sceneFile, "i", "scene.json", "json scene file")

		var imgOut string
		flag.StringVar(&imgOut, "o", "out.png", "image output")

		flag.Parse()

		if flag.Parsed() {
			path = imgOut
		}
	} else {
		// scene par defaut
		scene.width = raytracer.WinX
		scene.height = raytracer.WinY
		scene.d = raytracer.D
		scene.cam.pos.x = -300
		scene.cam.pos.y = 50

		scene.meshs = append(scene.meshs, mesh{})
	}

	imgRect := image.Rect(0, 0, scene.width, scene.height)
	img := image.NewRGBA(imgRect)

	raytracer.Raytracer(&scene, img)

	// create en populate file
	outfd, err := os.Create(path)
	defer outfd.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(outfd, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
