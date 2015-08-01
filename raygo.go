package main

import (
	"fmt"
	"os"
	"flag"
	"image"
	"image/png"
)

func usage() {
	fmt.Println("raygo [JSON_SCENE] [-o OUTFILE]")
}

func main() {
	var scene Scene
	var path string = "out.png"

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
		scene.Width = 1080
		scene.Height = 720
		scene.D = 100
		scene.Cam.Pos.X = -300
		scene.Cam.Pos.Z = 50
		
		scene.Meshs = append(scene.Meshs, Mesh{})
	}

	imgRect := image.Rect(0, 0, 100, 100)
	img := image.NewRGBA(imgRect)

	Raytracer(&scene, img)

	// create en populate file
	outfd, err := os.Create(path)
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
