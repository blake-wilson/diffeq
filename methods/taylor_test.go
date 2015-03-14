package diffeq

import (
	"fmt"
	"testing"
)

func BenchmarkTaylor(b *testing.B) {

	simpleFunc := func(params ...float64) float64 { return 3 * params[0] * params[1] }
	simpleFuncDeriv := func(params ...float64) float64 { return 6 * params[0] }
	times, estimates := Taylor(simpleFunc, 1, 0, 4, 0.01, simpleFuncDeriv)

	for i := 0; i < len(times); i++ {
		fmt.Printf("Value at t=%f is %f\n", times[i], estimates[i])
	}
}
