package main

import (
	"fmt"
	"os"
	"github.com/thomas-miele/raygo/types"
	"github.com/thomas-miele/raygo/raytracer"
)

func usage() {
	fmt.Println("raygo [JSON_SCENE]...")
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Raytracer")
		//scene := types.Scene{Width: 640, Height: 480, D: 100}
		//fmt.Println(scene)
	} else {
		var scene types.Scene

		scene.Width = 1080
		scene.Height = 720
		scene.D = 100
		scene.Cam.Pos.X = -300
		scene.Cam.Pos.Z = 50
		
		scene.Meshs = append(scene.Meshs, types.Mesh{})

		raytracer.Raytracer(&scene)
		fmt.Println(scene)
	}
}
