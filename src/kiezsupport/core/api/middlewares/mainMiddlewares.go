package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"level":"debug","time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}))

	log.Trace().Msg("Set echo CORS config")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://kiezsupport.org",
			"http://www.kiezsupport.org",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAllow,
			echo.HeaderAuthorization,
			echo.HeaderCookie,
			echo.HeaderSetCookie,
			echo.HeaderXCSRFToken,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderAcceptEncoding,
			echo.HeaderXRequestedWith},
		ExposeHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAllow,
			echo.HeaderAuthorization,
			echo.HeaderCookie,
			echo.HeaderSetCookie,
			echo.HeaderXCSRFToken,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderAcceptEncoding,
			echo.HeaderXRequestedWith}}))
}
