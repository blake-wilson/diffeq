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
