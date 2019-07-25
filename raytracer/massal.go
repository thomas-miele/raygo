package raytracer

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Gregory Massal as insipiration : http://www.massal.net/article/raytrace/page1.html
 * Massal scene txt format
 * width height
 * nbMat nbSphere nbLight;
 * the next nbMat lines are materials
 * the next nbSphere lines are spheres
 * the next nbLight lines are lights
**/

type rgb struct {
	red, green, blue float64
}

type point struct {
	x, y, z float64
}

type massalMaterial struct {
	color      rgb
	reflection float64
}

type massalSphere struct {
	position point
	size     float64
	material int
}

type massalLight struct {
	position point
	color    rgb
}

// MassalScene ...
type MassalScene struct {
	materials    []massalMaterial
	spheres      []massalSphere
	lights       []massalLight
	SizeX, SizeY int
}

// MassalAlgo implementation
func MassalAlgo(Image *image.RGBA, scene *MassalScene) {
	var pixel color.RGBA

	bounds := Image.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel.R = 255
			pixel.G = 99
			pixel.B = 71
			pixel.A = 255
			Image.Set(x, y, pixel)
		}
	}
}

// ReadMassalText ouvre un fichier au format de Gregory Massal
func ReadMassalText(scene *MassalScene) error {
	file, err := os.Open("scene_massal.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	elements := make([][]string, 0, len(text))
	for _, line := range text {
		elm := strings.Split(line, " ")
		elements = append(elements, elm)
	}
	//fmt.Println(elements)
	elements, err = populateScene(elements, scene, len(text))
	return err
}

func populateScene(elements [][]string, scene *MassalScene, length int) ([][]string, error) {
	var err error
	var nbMat, nbSphere, nbLight int

	scene.SizeX, err = strconv.Atoi(elements[0][0])
	scene.SizeY, _ = strconv.Atoi(elements[0][1])

	nbMat, err = strconv.Atoi(elements[1][0])

	nbMat, err = strconv.Atoi(elements[1][1])
	nbSphere, err = strconv.Atoi(elements[1][1])
	nbLight, err = strconv.Atoi(elements[1][2])

	lencmp := nbMat + nbSphere + nbLight + 2
	if err != nil || length != lencmp {
		fmt.Println("Massal scene format incorrect:", err)
		os.Exit(1)
	}
	scene.materials = make([]massalMaterial, 0, nbMat)
	scene.spheres = make([]massalSphere, 0, nbSphere)
	scene.lights = make([]massalLight, 0, nbLight)

	// Populate materials
	offset := 2
	for l := 0; l < nbMat; l++ {
		var material massalMaterial

		material.color.red, err = strconv.ParseFloat(elements[l+offset][0], 64)
		material.color.green, err = strconv.ParseFloat(elements[l+offset][1], 64)
		material.color.blue, err = strconv.ParseFloat(elements[l+offset][2], 64)
		material.reflection, err = strconv.ParseFloat(elements[l+offset][3], 64)

		scene.materials = append(scene.materials, material)
	}
	// Populate spheres
	offset = 2 + nbMat
	for l := 0; l < nbSphere; l++ {
		var sphere massalSphere

		sphere.position.x, err = strconv.ParseFloat(elements[l+offset][0], 64)
		sphere.position.y, err = strconv.ParseFloat(elements[l+offset][1], 64)
		sphere.position.z, err = strconv.ParseFloat(elements[l+offset][2], 64)
		sphere.size, err = strconv.ParseFloat(elements[l+offset][3], 64)
		sphere.material, err = strconv.Atoi(elements[l+offset][4])

		scene.spheres = append(scene.spheres, sphere)
	}
	// Populate lights
	offset = 2 + nbMat + nbSphere
	for l := 0; l < nbLight; l++ {
		var light massalLight

		light.position.x, err = strconv.ParseFloat(elements[l+offset][0], 64)
		light.position.y, err = strconv.ParseFloat(elements[l+offset][1], 64)
		light.position.z, err = strconv.ParseFloat(elements[l+offset][2], 64)
		light.color.red, err = strconv.ParseFloat(elements[l+offset][3], 64)
		light.color.green, err = strconv.ParseFloat(elements[l+offset][4], 64)
		light.color.blue, err = strconv.ParseFloat(elements[l+offset][5], 64)

		scene.lights = append(scene.lights, light)
	}
	return elements, err
}
