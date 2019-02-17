// Package handler connects router, extracts request data and executes services in go routines using channels.
package handler

import (
	"github.com/davidvartanian/demo-math/service"
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
	"net/http"
)

// Handler for divide endpoint
// @Summary Handler for divide endpoint
// @Accept  json
// @Produce  json
// @Param   body     body    types.Payload     true        "Payload"
// @Success 200 {object} types.Result	""
// @Failure 400 {object} types.Result "Bad Request"
// @Router /divide [post]
func HandleDivide(c echo.Context) error {
	payload := new(types.Payload)
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := c.Validate(payload); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	chp := make(chan types.Payload)
	chr := make(chan types.Result)

	go service.ProcessDivide(c, chp, chr)
	chp <- *payload
	result := <-chr
	st := http.StatusOK
	if result.Error != "" {
		st = http.StatusBadRequest
	}
	return c.JSON(st, result)
}
