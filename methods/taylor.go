package methods

import (
	"math"

	mTypes "github.com/blake-wilson/diffeq/types"
	"github.com/blake-wilson/exparser/types"
)

func Taylor(f types.AstNode, yInit float64, tInit float64, tFinal float64, timeStep float64, higherDerivatives ...types.AstNode) []*mTypes.EvalPoint {

	estimates := make([]*mTypes.EvalPoint, 0, int((tFinal-tInit)/timeStep+1))
	estimates = append(estimates, &mTypes.EvalPoint{
		Time:  tInit,
		Value: yInit,
	})

	ctx := types.NewContext()

	estimate := yInit

	for i := tInit + timeStep; i < tFinal; i += timeStep {
		un := estimate

		ctx.AssignVariable("t", i)
		ctx.AssignVariable("x", un)
		un += timeStep * f.Eval(ctx)

		for j := 0; j < len(higherDerivatives); j++ {
			ctx.AssignVariable("t", i/estimate)
			ctx.AssignVariable("x", Factorial(int32(j+2)))
			un += math.Pow(timeStep, float64(j+2)) * higherDerivatives[j].Eval(ctx)
		}
		estimates = append(estimates, &mTypes.EvalPoint{
			Time:  i,
			Value: un,
		})
		estimate = un

	}
	return estimates
}
