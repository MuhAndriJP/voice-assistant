package routes

import (
	speak "github.com/MuhAndriJP/voice-assistant/htgo-tts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("."))
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token", "app-key"},
		ExposeHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token", "app-key"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	e.GET("/speak", speak.Handler)
	return e
}
