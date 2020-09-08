package edsm

// Vector3 represents a 3D vector
type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// DistanceTo returns the distance to the target Vector3
func (v *Vector3) DistanceTo(target *Vector3) float64 {
	return distance(v, target)
}
