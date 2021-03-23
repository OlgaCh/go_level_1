package prime

import "testing"

func BenchmarkCustomPrimeImpl1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CustomPrimeImpl(1000)
	}
}

func BenchmarkGoPrimeImpl1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GoPrimeImpl(1000)
	}
}
