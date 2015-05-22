package evaluator

import (
	"strconv"

	"github.com/blake-wilson/exparser"
)

// evaluate the function over the given interval and return the values at each
// specified step
func EvalOverInterval(expr string, varName string, leftBound, rightBound, stepSize float64) []float64 {
	tokens := exparser.ParseExpr(expr)

	// Identify indices of variables in the tokens
	var varIndices []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == varName {
			varIndices = append(varIndices, i)
		}
	}

	var results []float64
	var res float64
	// replace variable with values between boundaries
	for i := leftBound; i <= rightBound; i += stepSize {
		insertValues(tokens, varIndices, i)
		res, _ = exparser.EvaluatePostfix(tokens)
		results = append(results, res)
	}

	return results
}

func insertValues(tokens []string, indices []int, value float64) {
	for i := 0; i < len(indices); i++ {
		tokens[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}
}
