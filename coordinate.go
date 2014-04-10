package main

import ("image"
		"math")

type Coordinate struct {
	X, Y float64
}

func (c *Coordinate) Rotate(angle float64) {
	cos, sin := math.Cos(angle), math.Sin(angle)
	c.X, c.Y = c.X*cos + c.Y*sin, c.Y*cos - c.X*sin
}

func (c *Coordinate) Scale(factor float64) {
	c.X, c.Y = c.X*factor, c.Y*factor
}

func (c *Coordinate) Add(c2 Coordinate) Coordinate {
	return Coordinate{c.X+c2.X, c.Y+c2.Y}
}

func (c *Coordinate) Sub(c2 Coordinate) Coordinate {
	return Coordinate{c.X-c2.X, c.Y-c2.Y}
}

func (c *Coordinate) Length() float64 {
	return math.Hypot(c.X, c.Y)
}

func (c *Coordinate) toPoint() image.Point {
	return image.Point{int(c.X), int(c.Y)}
}
