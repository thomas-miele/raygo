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
		scene.width = WinX
		scene.height = WinY
		scene.d = D
		scene.cam.pos.x = -300
		scene.cam.pos.y = 50
		
		scene.meshs = append(scene.meshs, mesh{})
	}

	imgRect := image.Rect(0, 0, scene.width, scene.height)
	img := image.NewRGBA(imgRect)

	Raytracer(&scene, img)

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
