// Creates multiple lines
package main

// Import packages
import ("image"; "image/color"; "image/png";
		"os"
		"log")

func createImage() {
	// Declare image size
	width, height := 512, 512

	// Create a new image
	img := image.Rect(0, 0, width, height)
	c := NewCanvas(img)
	c.DrawGradient()

	for x:=0; x<width; x+=8 {
		c.DrawLine(color.RGBA{0, 0, 0, 255},
					Coordinate{0.0, 0.0},
					Coordinate{float64(x), float64(height)})
	}

	name := "line.png"
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	png.Encode(file, c)
}

func main() {
	createImage()
}