package diffeq

import (
	"testing"

	"github.com/blake-wilson/exparser"
	"github.com/blake-wilson/exparser/types"
	"github.com/stretchr/testify/assert"
)

var simpleFunc, linearFunc types.AstNode

func init() {
	linearFunc, _ = exparser.EvalExpression("2*x")
	simpleFunc, _ = exparser.EvalExpression("3*x^2")
}

func TestEuler(t *testing.T) {
	_, estimates := Euler(linearFunc, 0, 1, 5, 0.01)
	assert.Equal(t, float64(0), estimates[len(estimates)-1])
}

func benchmarkEuler(ts float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Euler(simpleFunc, 1, 0, 4, ts)
	}
}

func BenchmarkEulerOneHundredth(b *testing.B)  { benchmarkEuler(0.01, b) }
func BenchmarkEulerOneThousandth(b *testing.B) { benchmarkEuler(0.001, b) }
func BenchmarkEulerOneMillionth(b *testing.B)  { benchmarkEuler(0.000001, b) }
