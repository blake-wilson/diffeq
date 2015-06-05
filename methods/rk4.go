package methods

import (
	"fmt"
	"math"

	mTypes "github.com/blake-wilson/diffeq/types"
	"github.com/blake-wilson/exparser/types"
)

const one_sixth = float64(1) / float64(6)

/** Runge-Kutta 4th order method - f is the derivative of the function being estimated,
y_init is the initial value of the function at the initial time,
t_init is the initial time, and timestep is the time step */
func RK4(f types.AstNode, yInit float64, tInit float64, tFinal float64, timestep float64) ([]*mTypes.EvalPoint, error) {

	estimates := make([]*mTypes.EvalPoint, 0, int((tFinal-tInit)/timestep+1))
	ctx := types.NewContext()

	// Initial conditions
	ctx.AssignVariable("t", tInit)
	ctx.AssignVariable("x", yInit)

	estimate := yInit
	estimates = append(estimates, &mTypes.EvalPoint{
		Time:  tInit,
		Value: yInit,
	})

	for i := tInit + timestep; i < tFinal; i += timestep {
		k1 := f.Eval(ctx)

		ctx.AssignVariable("t", i+timestep/2)
		ctx.AssignVariable("x", estimate+k1*timestep/2)
		k2 := f.Eval(ctx)

		ctx.AssignVariable("x", estimate+k2*timestep/2)
		k3 := f.Eval(ctx)

		ctx.AssignVariable("t", i+timestep)
		ctx.AssignVariable("x", estimate+timestep*k3)
		k4 := f.Eval(ctx)

		estimate += one_sixth * timestep * (k1 + 2*k2 + 2*k3 + k4)

		if math.IsInf(estimate, 0) {
			// solution diverges
			return estimates, fmt.Errorf("numerical method diverged")
		}

		estimates = append(estimates, &mTypes.EvalPoint{
			Time:  i,
			Value: estimate,
		})

		// increment by timestep and define x_n+1
		ctx.AssignVariable("t", i+timestep)
		ctx.AssignVariable("x", estimate)
	}

	return estimates, nil
}
