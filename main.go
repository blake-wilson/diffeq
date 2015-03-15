package main

import (
	"fmt"
	"net/http"

	"github.com/blake-wilson/diffeq/methods"
)

func simpleFunc(params ...float64) float64 {
	return 3 * params[0] * params[0]
}

func simpleFuncDeriv(params ...float64) float64 {
	return 6 * params[0]
}

func handler(w http.ResponseWriter, r *http.Request) {
	times, estimates := diffeq.Taylor(simpleFunc, 1, 0, 4, 0.01, simpleFuncDeriv)
	var printStr string
	for i := 0; i < len(times); i++ {
		printStr += fmt.Sprintf("Estimate for time %f is %f\n", (times[i]), estimates[i])
	}
	fmt.Fprintf(w, "%s", printStr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
