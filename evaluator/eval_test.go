package evaluator

import (
	"fmt"
	"testing"
)

func TestEvalOverInterval(t *testing.T) {
	var expr = "x + 5"
	results := EvalOverInterval(expr, "x", 1, 4, 0.5)
	fmt.Printf("%s", results)

}
