package service

import (
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessDivideSuccess(t *testing.T) {
	c := new(echo.Context)
	chp := make(chan types.Payload)
	chr := make(chan types.Result)

	go ProcessDivide(*c, chp, chr)
	chp <- types.Payload{
		A: 1,
		B: 2,
	}
	result := <-chr
	if result.Error != "" {
		t.Fail()
	}
	assert.Equal(t, result.Result, 0.5)
}

func TestProcessDivideByZeroFail(t *testing.T) {
	c := new(echo.Context)
	chp := make(chan types.Payload)
	chr := make(chan types.Result)

	go ProcessDivide(*c, chp, chr)
	chp <- types.Payload{
		A: 1,
		B: 0,
	}
	result := <-chr
	if result.Error == "" {
		t.Fail()
	}
}
