package test2

import "testing"

func BenchmarkTest_g0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		g0(&a)
	}
}

func BenchmarkTest_g1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		g1(&a)
	}
}
