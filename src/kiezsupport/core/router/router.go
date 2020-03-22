package router

import (
	"kiezsupport/core/api"
	"kiezsupport/core/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetCompleteLogMiddlware(e)

	// set main routes
	api.MainGroup(e)

	return e
}
