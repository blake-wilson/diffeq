package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalOverInterval(t *testing.T) {
	var expr = "x + 5"
	results := EvalOverInterval(expr, "x", 1, 4, 0.5)
	assert.Equal(t, []float64{6, 6.5, 7, 7.5, 8, 8.5, 9},
		results)
}
