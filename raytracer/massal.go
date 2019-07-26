package raytracer

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
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

type massalMaterial struct {
	color      rgb
	reflection float64
}

type massalSphere struct {
	position Point
	size     float64
	material int
}

type massalLight struct {
	position Point
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
			pixel = massalPixel(scene, x, y)
			Image.Set(x, y, pixel)
		}
	}
}

func massalPixel(scene *MassalScene, x, y int) color.RGBA {
	//pixel := color.RGBA{255, 99, 71, 255}
	pixel := color.RGBA{0, 0, 0, 255}
	var red, green, blue float64
	var coef = 1.0
	var level int
	var ray RayLight
	ray.Position = Point{float64(x), float64(y), -10000.0}
	ray.Direction = Point{0.0, 0.0, 1.0}

	for ok := true; ok; ok = (coef > 0.0 && level < 10) {
		// recherche de l'intersection la plus proche
		var t = 20000.0
		var currentSphere = -1

		for i, elm := range scene.spheres {
			var isHit = false
			t, isHit = hitSphere(ray, elm, t)
			if isHit {
				currentSphere = i
			}
		}

		if currentSphere == -1 {
			break
		}

		var start = ray.Position.plus(ray.Direction.times(t))
		// la normale au point d'intersection
		var n = start.minus(scene.spheres[currentSphere].position)
		var temp = n.dot(n)
		if temp == 0.0 {
			break
		}

		temp = 1.0 / math.Sqrt(temp)
		n = n.times(temp)

		var currentMaterial = scene.materials[scene.spheres[currentSphere].material]
		// calcul de la valeur d'éclairement au point
		for _, current := range scene.lights {
			var dist = current.position.minus(start)
			if n.dot(dist) <= 0.0 {
				continue
			}
			var t = math.Sqrt(dist.dot(dist))
			if t <= 0.0 {
				continue
			}
			var lightRay RayLight
			lightRay.Position = start
			lightRay.Direction = dist.times(1 / t)
			// calcul des ombres
			var inShadow = false
			for _, sphere := range scene.spheres {
				var isHit = false
				t, isHit = hitSphere(lightRay, sphere, t)
				if isHit {
					inShadow = true
					break
				}
			}
			if !inShadow {
				// lambert
				var lambert = lightRay.Direction.dot(n) * coef
				red += lambert * current.color.red * currentMaterial.color.red
				green += lambert * current.color.green * currentMaterial.color.green
				blue += lambert * current.color.blue * currentMaterial.color.blue
			}
			// on itère sur la prochaine réflexion
			coef *= currentMaterial.reflection
			var reflet = 2.0 * ray.Direction.dot(n)
			ray.Position = start
			ray.Direction = ray.Direction.minus(n.times(reflet))

			level++
		}
	}
	pixel.R = uint8(red)
	pixel.G = uint8(green)
	pixel.B = uint8(blue)
	return pixel
}

func hitSphere(ray RayLight, sphere massalSphere, t float64) (float64, bool) {
	dist := ray.Position.minus(ray.Direction)
	var B = ray.Direction.dot(dist)
	var D = (B * B) - dist.dot(dist) + (sphere.size * sphere.size)
	if D < 0.0 {
		return 0, false
	}
	var t0 = B - math.Sqrt(D)
	var t1 = B + math.Sqrt(D)
	var retvalue = false
	if (t0 > 0.1) && (t0 < t) {
		t = t0
		retvalue = true
	}
	if (t1 > 0.1) && (t1 < t) {
		t = t1
		retvalue = true
	}
	return t, retvalue
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

		sphere.position.X, err = strconv.ParseFloat(elements[l+offset][0], 64)
		sphere.position.Y, err = strconv.ParseFloat(elements[l+offset][1], 64)
		sphere.position.Z, err = strconv.ParseFloat(elements[l+offset][2], 64)
		sphere.size, err = strconv.ParseFloat(elements[l+offset][3], 64)
		sphere.material, err = strconv.Atoi(elements[l+offset][4])

		scene.spheres = append(scene.spheres, sphere)
	}
	// Populate lights
	offset = 2 + nbMat + nbSphere
	for l := 0; l < nbLight; l++ {
		var light massalLight

		light.position.X, err = strconv.ParseFloat(elements[l+offset][0], 64)
		light.position.Y, err = strconv.ParseFloat(elements[l+offset][1], 64)
		light.position.Z, err = strconv.ParseFloat(elements[l+offset][2], 64)
		light.color.red, err = strconv.ParseFloat(elements[l+offset][3], 64)
		light.color.green, err = strconv.ParseFloat(elements[l+offset][4], 64)
		light.color.blue, err = strconv.ParseFloat(elements[l+offset][5], 64)

		scene.lights = append(scene.lights, light)
	}
	return elements, err
}
