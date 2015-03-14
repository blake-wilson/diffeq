package diffeq

// try to use the convention function( time, previous value)
type scalar_func func(float64, float64) float64
