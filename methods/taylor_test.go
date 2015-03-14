package diffeq

import "testing"

func benchmarkTaylor(ts float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Taylor(simpleFunc, 1, 0, 4, ts, simpleFuncDeriv)
	}
}

// test various timesteps for the method.
func BenchmarkTaylorOneHundredth(b *testing.B)  { benchmarkTaylor(0.01, b) }
func BenchmarkTaylorOneThousandth(b *testing.B) { benchmarkTaylor(0.001, b) }
func BenchmarkTaylorOneMillionth(b *testing.B)  { benchmarkTaylor(0.000001, b) }
