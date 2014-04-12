// Reference: http://www.pheelicks.com/2013/11/intro-to-images-in-go-fractals/

package main

// Import packages
import ("image"; "image/color"; "image/draw";
		"os"
		"log")

// Declare a new structure
type Canvas struct {
	image.RGBA
}

func NewCanvas(r image.Rectangle) *Canvas {
	canvas := new(Canvas)
	canvas.RGBA = *image.NewRGBA(r)
	return canvas
}

func (c Canvas) Clone() *Canvas {
	clone := NewCanvas(c.Bounds())
	copy(clone.Pix, c.Pix)
	return clone
}

func (c Canvas) DrawGradient() {
	size := c.Bounds().Size()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := color.RGBA{
							uint8(255 * x / size.X),
							uint8(255 * y / size.Y),
							55,
							255}

			c.Set(x, y, color)
		}
	}
}

func (c Canvas) DrawLine(color color.RGBA, from Coordinate, to Coordinate) {
	delta := to.Sub(from)
	length := delta.Length()
	xStep, yStep := delta.X/length, delta.Y/length
	limit := int(length+0.5)

	for i:=0; i<limit; i++ {
		x := from.X + float64(i)*xStep
		y := from.Y + float64(i)*yStep
		c.Set(int(x), int(y), color)
	}
}

func (c Canvas) DrawSpiral(color color.RGBA, from Coordinate, iterations uint32, 
							degree float64, factor float64) {
	dir := Coordinate{0, 3.5}
	last := from
	var i uint32

	// Iterations defines the number of small lines drawn
	for i = 0; i<iterations; i++ {
		next := last.Add(dir)
		c.DrawLine(color, last, next)
		// Only rotation will create a circle
		dir.Rotate(degree)

		// This scaling is the one which is doing the magic
		dir.Scale(factor)
		last = next
	}
}

func (c Canvas) DrawRect(color color.RGBA, from Coordinate, to Coordinate) {
	for x := int(from.X); x<=int(to.X); x++ {
		for y := int(from.Y); y <= int(to.Y); y++ {
			c.Set(x, y, color)
		}
	}
}

func (c Canvas) SaveToFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	// png.Encode(file, c.RGBA)
}

func CreateCanvas(fileName string) *Canvas {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	c := NewCanvas(img.Bounds())
	draw.Draw(c, img.Bounds(), img, image.ZP, draw.Src)
	return c
}