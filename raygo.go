package main

import (
	"os"
	"fmt"
	"image"
	"image/png"

	"net/http"
	"github.com/labstack/echo"

	"github.com/thomas-miele/raygo/raytracer"
)

func main() {
	esrv := echo.New()

	esrv.Static("/", "www")

	esrv.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World !")
	})

	esrv.GET("/ray", func(c echo.Context) error {
		raygo()
		return c.HTML(http.StatusOK, "<img src=\"out.png\">")
		//return c.HTML(http.StatusOK, "<img src=\"dummy.png\">")

	})

	esrv.Logger.Fatal(esrv.Start(":8000"))
}

func raygo() {
	var scene raytracer.Scene
	var path string = "www/out.png"

	// si arguments -> utiliser une scene json
	// scene par defaut
	scene.Width = raytracer.WinX
	scene.Height = raytracer.WinY
	scene.D = raytracer.D
	scene.Cam.Pos.X = -300
	scene.Cam.Pos.Y = 50

	scene.Meshs = append(scene.Meshs, raytracer.Mesh{})

	imgRect := image.Rect(0, 0, scene.Width, scene.Height)
	img := image.NewRGBA(imgRect)

	raytracer.Raytracer(&scene, img)

	// create and populate file
	var err1 = os.Remove(path)
	if err1 != nil {
		fmt.Println(err1)
	}
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
