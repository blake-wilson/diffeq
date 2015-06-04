package methods

import (
	"testing"

	"github.com/blake-wilson/exparser"
	"github.com/blake-wilson/exparser/types"
)

var taylorSimpleFunc, taylorSimpleFuncDeriv types.AstNode

func init() {
	taylorSimpleFunc, _ = exparser.EvalExpression("3*x^2")
	taylorSimpleFuncDeriv, _ = exparser.EvalExpression("6*x")
}

func benchmarkTaylor(ts float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Taylor(taylorSimpleFunc, 1, 0, 4, ts, taylorSimpleFuncDeriv)
	}
}

// test various timesteps for the method.
func BenchmarkTaylorOneHundredth(b *testing.B)  { benchmarkTaylor(0.01, b) }
func BenchmarkTaylorOneThousandth(b *testing.B) { benchmarkTaylor(0.001, b) }
func BenchmarkTaylorOneMillionth(b *testing.B)  { benchmarkTaylor(0.000001, b) }
