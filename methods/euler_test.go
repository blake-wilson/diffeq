package diffeq

import (
	"fmt"
	"testing"
)

func TestEuler(t *testing.T) {

	simpleFunc := func(time float64, val float64) float64 { return 3 * time * time }
	times, estimates := Euler(simpleFunc, 1, 0, 4, 0.01)

	fmt.Printf("SDFDFSDFS")
	for i := 0; i < len(times); i++ {
		fmt.Printf("Value at t=%f is %f\n", times[i], estimates[i])
	}
}
