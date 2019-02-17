package helper

import (
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
)

func RecoverPanic(c echo.Context, ch chan types.Result) {
	if r := recover(); r != nil {
		c.Logger().Panic(r.(string))
		r := types.Result{Error: r.(string)}
		ch <- r
	}
}
