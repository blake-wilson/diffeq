package methods

import (
	"fmt"
	"math"

	mTypes "github.com/blake-wilson/diffeq/types"
	"github.com/blake-wilson/exparser/types"
)

/** Euler's method - f is the derivative of the function being estimated,
* y_init is the initial value of the function at the initial time,
* t_init is the initial time, and timeStep is the time step. */
func Euler(f types.AstNode, yInit float64, tInit float64, tFinal float64, timeStep float64) ([]*mTypes.EvalPoint, error) {

	estimates := make([]*mTypes.EvalPoint, 0, int((tFinal-tInit)/timeStep+1))
	ctx := types.NewContext()

	// Initial conditions
	ctx.AssignVariable("t", tInit)
	ctx.AssignVariable("x", yInit)

	estimate := yInit
	estimates = append(estimates, &mTypes.EvalPoint{
		Time:  tInit,
		Value: yInit,
	})

	for i := tInit + timeStep; i < tFinal; i += timeStep {
		estimate = estimate + timeStep*f.Eval(ctx)

		// Error check estimate
		if math.IsInf(estimate, 0) {
			// cannot continue approximation - value too large
			return estimates, fmt.Errorf("numerical method diverged")
		}

		estimates = append(estimates, &mTypes.EvalPoint{
			Time:  i,
			Value: estimate,
		})

		ctx.AssignVariable("t", i)
		ctx.AssignVariable("x", estimate)

	}
	return estimates, nil
}
