package main

import "math"

// Vertex represents a 2D point
type Vertex struct {
	X, Y float64
}

// Abs calculates absolute distance from origin
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Add adds n to both coordinates
func (v *Vertex) add(n float64) {
	v.X += n
	v.Y += n
}
