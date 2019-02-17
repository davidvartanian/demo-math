package service

import (
	"github.com/davidvartanian/demo-math/helper"
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
	"math"
)

// Divide numbers provided via Payload object whenever its channel receives one
func ProcessDivide(c echo.Context, chap chan types.Payload, char chan types.Result) {
	defer helper.RecoverPanic(char)
	pp := <-chap
	op := op(pp.A, pp.B)
	if op == math.Inf(1) {
		r := types.Result{Error: "Division by zero!"}
		c.Logger().Panic(r.Error)
		char <- r
	} else {
		r := types.Result{Result: op}
		char <- r
	}
}

func op(a float64, b float64) float64 {
	return a / b
}
