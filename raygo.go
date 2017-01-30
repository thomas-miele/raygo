package main

import (
	"os"
	"fmt"
	"image"
	"image/png"
	"bytes"

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
	esrv.GET("/ray", raygo)

	esrv.Logger.Fatal(esrv.Start(":8000"))
}

func raygo(c echo.Context) error {
	var scene raytracer.Scene

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

	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return c.Stream(http.StatusOK, "image/png", buffer)
}
