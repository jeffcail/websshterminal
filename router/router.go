package router

import (
	"time"

	"github.com/c/websshterminal.io/middlewares"
	"github.com/c/websshterminal.io/ubzer"

	"github.com/c/websshterminal.io/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunSshTerminal
func RunSshTerminal() {
	e := echo.New()
	e.Static("/", "dist")
	e.Static("/ssh/node", "dist")
	e.Static("/ssh/dial", "dist")

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: ubzer.EchoLog,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middlewares.RequestLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.DefaultBodyDumpConfig))

	e.GET("/ssh", handler.ShellWeb)

	e.Logger.Fatal(e.Start(":5555"))
}
