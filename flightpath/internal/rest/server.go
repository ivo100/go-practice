package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

const (
	Version       = "v2023.4.10"
	HostPort      = ":8080"
	RouteEndpoint = "/calculate"
)

func Setup() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		e.Logger.SetLevel(log.DEBUG)
	}
	return e
}

// Serve is entry point of the REST server
func Serve() {
	e := Setup()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, Version)
	})

	e.POST(RouteEndpoint, FindRoute)

	e.Logger.Fatal(e.Start(HostPort))
}
