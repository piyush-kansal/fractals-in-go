// Creates a node network
// Reference: http://www.pheelicks.com/2013/11/intro-to-images-in-go-fractals/

package main

// Import packages
import ("image"; "image/color"; "image/png";
		"os"
		"log"
		"math"; "math/cmplx";
		"fmt")

// Converts a point on a canvas to a complex number
func toComplex(x, y int, zoom float64, center complex128) complex128 {
	return center + complex(float64(x)/zoom, float64(y)/zoom)
}

// Iterate as per Mandelbrot algo and returns the magnitude
func iterate(c complex128, iter int) float64 {
	z := complex(0, 0)
	for i := 1; i < iter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2000 {
			return 2000
		}
	}

	return cmplx.Abs(z)
}

// Converts magnitude into a color based on the image gradient
func createColorizer(fileName string) func(float64) color.Color {
	gradient := CreateCanvas(fileName)
	yLimit := gradient.Bounds().Size().Y - 1
	return func(magnitude float64) color.Color {
		m := int(math.Max(math.Min(300*magnitude, float64(yLimit)), 1))
		return gradient.At(0, m)
	}
}

func drawFractal(canvas *Canvas, zoom float64, center complex128, colorizer func(float64) color.Color) {
	size := canvas.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			c := toComplex(x-size.X/2, y-size.Y/2, zoom, center)
			mag := iterate(c, 150)
			color := colorizer(mag)
			canvas.Set(x, y, color)
		}
	}
}

func createFractal(zoom float64, real float64, imag float64, gradFile string) {
	width, height := 2048, 1024
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	center := complex(real, imag)
	colorizer := createColorizer(gradFile)
	drawFractal(canvas, zoom, center, colorizer)

	name := fmt.Sprintf("fractal_%f_%f_%f.png", zoom, real, imag)
	outFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	png.Encode(outFile, canvas)
}

func main() {
	createFractal(100, 0, 0, "gradients/gradient1.png")
	createFractal(1000, 0, 0, "gradients/gradient2.png")
	createFractal(16000, 0, 0, "gradients/gradient2.png")
	createFractal(6000, 0.75, 0.25, "gradients/gradient3.png")
	createFractal(16000.0, -0.71, -0.25, "gradients/gradient3.png")
	createFractal(30000.0, -0.71, -0.25, "gradients/gradient3.png")
	createFractal(100000.0, -0.71, -0.25, "gradients/gradient1.png")
}