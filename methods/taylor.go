package diffeq

import "math"

func Taylor(f scalar_func, yInit float64, tInit float64, tFinal float64, timeStep float64, higherDerivatives ...scalar_func) ([]float64, []float64) {

	var timeIntervals []float64
	var estimates []float64

	estimate := yInit
	for i := tInit; i < tFinal; i += timeStep {
		timeIntervals = append(timeIntervals, i)
		un := estimate
		un += timeStep * f(i, estimate)

		for j := 0; j < len(higherDerivatives); j++ {
			un += math.Pow(timeStep, float64(j+2)) * higherDerivatives[j](i, estimate) / Factorial(int32(j+2))
		}
		estimates = append(estimates, un)
		estimate = un

	}
	return timeIntervals, estimates
}
