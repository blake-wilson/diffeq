package diffeq

func Factorial(value int32) float64 {
	if value == 0 || value == 1 {
		return 1
	}

	retval := float64(1)

	for i := 2; int32(i) <= value; i++ {
		retval *= float64(i)
	}

	return retval
}

// Newton's root-finding method.  Generally a part of estimation of terms for
// implicit numerical methods
func Newton(f scalar_func, fprime scalar_func, tolerance float64, maxIterations int64) float64 {
	// random first guess is left to be 0
	var estimate float64

	for difference := tolerance + 1; difference > tolerance && numIters <= maxIterations {
		estimate -= f(estimate) / fprime(estimate)
	}
	return estimate
}
