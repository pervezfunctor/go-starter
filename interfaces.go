package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return float64(r.width * r.height)
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * float64(math.Pi)
}

func TotalArea(s []Shape) float64 {
	r := s[0].(Rectangle)
	r.width *= 0
	var sum float64 = 0

	for _, shp := range s {
		sum += shp.Area()
	}

	return sum
}
