package edsm

import (
	"math"
)

func distance(vec1 *Vector3, vec2 *Vector3) float64 {
	a := math.Abs(vec1.X - vec2.X)
	b := math.Abs(vec1.Y - vec2.Y)
	c := math.Abs(vec1.Z - vec2.Z)

	a = math.Pow(a, 2)
	b = math.Pow(b, 2)
	c = math.Pow(c, 2)

	return math.Sqrt(a + b + c)
}
