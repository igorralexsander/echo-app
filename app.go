package app

import (
	"github.com/igorralexsander-corp/echoapp/internal/infrastructure/api/routes"
)

type Application struct {
	Routes *Routes
}

type Routes struct {
	Domain *routes.Domain
}

func NewApplication() Application {
	instance := Application{}

	instance.Routes = instance.initRoutes()

	return instance
}

func (a *Application) initRoutes() *Routes {
	return &Routes{
		Domain: a.provideDomainRoute(),
	}
}

func (a *Application) provideDomainRoute() *routes.Domain {
	return routes.NewDomainRoute()
}
