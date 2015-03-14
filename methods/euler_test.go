package diffeq

import "testing"

func benchmarkEuler(ts float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Euler(simpleFunc, 1, 0, 4, ts)
	}
}

func BenchmarkEulerOneHundredth(b *testing.B)  { benchmarkEuler(0.01, b) }
func BenchmarkEulerOneThousandth(b *testing.B) { benchmarkEuler(0.001, b) }
func BenchmarkEulerOneMillionth(b *testing.B)  { benchmarkTaylor(0.000001, b) }
