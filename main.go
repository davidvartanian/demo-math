package main

import (
	_ "github.com/davidvartanian/demo-math/docs"
	"github.com/davidvartanian/demo-math/handler"
	"github.com/davidvartanian/demo-math/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// @title Swagger Demo Math
// @version 1.0
// @description This is a demo micro service
// @termsOfService http://swagger.io/terms/

// @contact.name David Vartanian
// @contact.url http://www.swagger.io/support
// @contact.email davidvartanian@posteo.de

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, user_agent=${user_agent}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Validator = &types.PayloadValidator{Validator: validator.New()}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/sum", handler.HandleSum)
	e.POST("/subtract", handler.HandleSubtract)
	e.POST("/multiply", handler.HandleMultiply)
	e.POST("/divide", handler.HandleDivide)
	e.Logger.Fatal(e.Start(":1323"))
}
