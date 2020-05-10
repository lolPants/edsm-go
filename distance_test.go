package edsm

import (
	"testing"
)

var (
	vecA Vector3 = Vector3{3.03125, -0.09375, 3.15625}
	vecB Vector3 = Vector3{-81.78125, -149.4375, -343.375}
)

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distance(&vecA, &vecB)
	}
}
