package main

import (
	"fmt"
	"os"
	"flag"
)

func usage() {
	fmt.Println("raygo [JSON_SCENE] [-o OUTFILE]")
}

func main() {
	var scene Scene

	if len(os.Args) > 1 {
		var sceneFile string
		flag.StringVar(&sceneFile, "i", "scene.json", "json scene file")

		var imgOut string
		flag.StringVar(&imgOut, "o", "out.png", "image output")
		
		flag.Parse()
	} else {
		scene.Width = 1080
		scene.Height = 720
		scene.D = 100
		scene.Cam.Pos.X = -300
		scene.Cam.Pos.Z = 50
		
		scene.Meshs = append(scene.Meshs, Mesh{})
	}
	Raytracer(&scene)
	fmt.Println(scene)
}
