package api

import (
	app "github.com/igorralexsander-corp/echoapp"
	"github.com/labstack/echo"
)

func NewHttpServer(application *app.Application) *echo.Echo {
	e := echo.New()
	//e.HideBanner = true
	//e.HidePort = true

	registerRoutes(application, e)
	return e
}

func registerRoutes(application *app.Application, e *echo.Echo) {
	application.Routes.Domain.Register(e)
}

func Start(e *echo.Echo, host string) {
	e.Start(host)
}
