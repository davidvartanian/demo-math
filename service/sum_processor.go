package service

import (
	"github.com/davidvartanian/demo-math/helper"
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
)

func ProcessSum(c echo.Context, chap chan types.Payload, char chan types.Result) {
	defer helper.RecoverPanic(c, char)
	pp := <-chap
	op := pp.A + pp.B
	r := types.Result{Result: op}
	char <- r
}
