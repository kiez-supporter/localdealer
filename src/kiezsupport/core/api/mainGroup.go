package api

import (
	"kiezsupport/core/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {

	//preflight
	e.OPTIONS("/v1/health", handlers.Hello)
	e.GET("/v1/health", handlers.Hello)

	e.GET("/v1/locations", handlers.GetLocations)
	e.GET("/v1/location", handlers.GetLocationById)

	//e.POST("/v1/locations", handlers.AddLocation)
	//e.PUT("/v1/locations", handlers.EditLocation)
	//e.DELETE("/v1/locations", handlers.DeleteLocation)
}
