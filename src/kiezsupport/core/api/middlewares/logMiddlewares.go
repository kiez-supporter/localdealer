package middlewares

import (
	"encoding/json"

	"kiezsupport/core/modules/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// SetCompleteLogMiddlware Middleware for logging request and response
func SetCompleteLogMiddlware(e *echo.Echo) {

	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		var bodyJSON interface{}
		var bodyRESP interface{}

		json.Unmarshal(reqBody, &bodyJSON)
		json.Unmarshal(resBody, &bodyRESP)

		log.Logger.WithFields(logrus.Fields{
			"remote_ip": c.RealIP(),
			"protocol":  c.Request().Proto,
			"host":      c.Request().Host,
			"uri":       c.Request().RequestURI,
			"headers":   c.Request().Header,
			"request":   bodyJSON,
			"response":  bodyRESP,
		}).Info("REQUEST LOG")
	}))
}
