package server

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/yaml.v3"

	"github.com/tmiele/raygo/raytracer"
)

func server() {
	esrv := echo.New()

	esrv.Static("/", "www")

	// esrv.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World !")
	// })
	esrv.GET("/ray", raygo)
	esrv.GET("/ray/:scene", raygo)

	esrv.Logger.Fatal(esrv.Start(":8000"))
}

func raygo(c echo.Context) error {
	var scene raytracer.Scene
	var ymld SceneYml

	name := c.Param("scene")
	if name != "" {
		fileScene(&ymld, name)
	} else {
		defaultScene(&ymld)
	}

	scene.Width = int(ymld.Width)
	scene.Height = int(ymld.Height)
	scene.D = int(ymld.D)
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

	// png bffer as http response with content-type
	return c.Stream(http.StatusOK, "image/png", buffer)
}

func fileScene(ymld *SceneYml, name string) {

	filename := "www/scenes/" + name + ".yaml"
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(src, &ymld)
	if err != nil {
		panic(err)
	}
	fmt.Println(ymld)
}

func defaultScene(ymld *SceneYml) {

	data := []byte(dataScene)
	err := yaml.Unmarshal(data, &ymld)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ymld)
}

// Scene Yaml like www/scenes/new.yaml
type SceneYml struct {
	Width  float32 `yaml:width`
	Height float32 `yaml:height`
	D      float32 `yaml:d`
	Camera struct {
		Forward []float32 `yaml:forward`
		Right   []float32 `yaml:right`
		Up      []float32 `yaml:up`
	} `yaml:camera`
	Meshs []struct {
		Name     string    `yaml:name`
		Position []float32 `yaml:position`
		Surface  string    `yaml:surface`
	} `yaml:things`
	Lights []struct {
		Name     string    `yaml:name`
		Position []float32 `yaml:position`
		Color    []float32 `yaml:color`
	} `yaml:lights`
}

// equivalent of www/scenes/new.yaml
var dataScene = `
width: 640
height: 480
d: 100
camera:
  forward: [3.0, 2.0, 4.0]
  right: [-1.0, 0.5, 0.0]
  up: [1, 1, 1]
things:
  - name: plane
    position: [0.0, 1.0, 0.0]
    surface: checkerboard
  - name: sphere
    position: [0.0, 1.0, -0.25]
    surface: shiny
  - name: sphere
    position: [-1.0, 0.5, 1.5]
    surface: shiny
lights:
  - name: spot
    position: [-2.0, 2.5, 0.0]
    color: [0.49, 0.07, 0.07]
  - name: sun
    position: [1.5, 2.5, 1.5]
    color: [0.07, 0.07, 0.49]
  - name: point
    position: [1.5, 2.5, -1.5]
    color: [0.07, 0.49, 0.071]
  - name: area
    position: [0, 3.5, 0]
    color: [0.21, 0.21, 0.35]
`
