package diffeq

import (
	"fmt"
	"testing"
)

func TestEuler(t *testing.T) {

	simpleFunc := func(params ...float64) float64 { return 3 * params[0] * params[1] }
	times, estimates := Euler(simpleFunc, 1, 0, 4, 0.01)

	for i := 0; i < len(times); i++ {
		fmt.Printf("Value at t=%f is %f\n", times[i], estimates[i])
	}
}
