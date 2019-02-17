package service

import (
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessSum(t *testing.T) {
	c := new(echo.Context)
	chp := make(chan types.Payload)
	chr := make(chan types.Result)

	go ProcessSum(*c, chp, chr)
	chp <- types.Payload{
		A: 2,
		B: 3,
	}
	result := <-chr
	if result.Error != "" {
		t.Fail()
	}
	assert.Equal(t, 5.0, result.Result)
}
