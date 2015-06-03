package diffeq

import "github.com/blake-wilson/exparser/types"

/** Euler's method - f is the derivative of the function being estimated,
* y_init is the initial value of the function at the initial time,
* t_init is the initial time, and timeStep is the time step. */
func Euler(f types.AstNode, yInit float64, tInit float64, tFinal float64, timeStep float64) ([]float64, []float64) {

	// hmmmm, parallel arrays
	var timeIntervals []float64
	var estimates []float64
	estimate := yInit
	ctx := types.NewContext()

	// Initial conditions
	ctx.AssignVariable("t", tInit)
	ctx.AssignVariable("x", yInit)

	for i := tInit; i < tFinal; i += timeStep {
		timeIntervals = append(timeIntervals, i)

		estimate = estimate + timeStep*f.Eval(ctx)
		ctx.AssignVariable("t", i)
		ctx.AssignVariable("x", estimate)

		estimates = append(estimates, estimate)
	}
	return timeIntervals, estimates
}
