package main

import "math"

type Point struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Rect struct {
	Point
	Size
}

func (r Rect) RectArea() int {
	return r.Width * r.Height
}

func (pt Point) Distance() float64 {
	return math.Sqrt(float64(pt.X*pt.X + pt.Y*pt.Y))
}

type Point3D struct {
	Point
	Z int
}

func Distance(pt Point3D) float64 {
	return math.Sqrt(float64(pt.X*pt.X + pt.Y*pt.Y + pt.Z*pt.Z))
}
