// Creates multiple spirals
package main

// Import packages
import ("image"; "image/color"; "image/png";
		"os"
		"log"
		"math/rand"
		"time"
		"fmt")

func createImage(iterations uint32, degree float64, factor float64) {
	// Declare image size
	width, height := 2048, 1024

	// Create a new image
	img := image.Rect(0, 0, width, height)
	c := NewCanvas(img)
	c.DrawGradient()

	rand.Seed(time.Now().UTC().UnixNano())
	for i:=0; i<300; i++ {
		x := float64(width) * rand.Float64()
		y := float64(height) * rand.Float64()
		color := color.RGBA{uint8(rand.Intn(255)),
							uint8(rand.Intn(255)),
							uint8(rand.Intn(255)),
							255}
		c.DrawSpiral(color, Coordinate{x, y}, iterations, degree, factor)
	}

	name := fmt.Sprintf("spiral_%d_%f_%f.png", iterations, degree, factor)
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	png.Encode(file, c)
}

func main() {
	createImage(9000, 0.04, 0.999)
	createImage(10000, 0.04, 0.9999)
}