package diffeq

import (
	"math"

	"github.com/blake-wilson/exparser/types"
)

func Taylor(f types.AstNode, yInit float64, tInit float64, tFinal float64, timeStep float64, higherDerivatives ...types.AstNode) ([]float64, []float64) {

	var timeIntervals []float64
	var estimates []float64

	ctx := types.NewContext()

	estimate := yInit
	for i := tInit; i < tFinal; i += timeStep {
		timeIntervals = append(timeIntervals, i)
		un := estimate

		ctx.AssignVariable("t", i)
		ctx.AssignVariable("x", un)
		un += timeStep * f.Eval(ctx)

		for j := 0; j < len(higherDerivatives); j++ {
			ctx.AssignVariable("t", i/estimate)
			ctx.AssignVariable("x", Factorial(int32(j+2)))
			un += math.Pow(timeStep, float64(j+2)) * higherDerivatives[j].Eval(ctx)
		}
		estimates = append(estimates, un)
		estimate = un

	}
	return timeIntervals, estimates
}
