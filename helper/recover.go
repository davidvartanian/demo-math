package helper

import (
	"github.com/davidvartanian/demo-math/types"
)

func RecoverPanic(ch chan types.Result) {
	if r := recover(); r != nil {
		ch <- types.Result{Error: "Internal error!"}
	}
}
