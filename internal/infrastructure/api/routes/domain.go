package routes

import (
	"github.com/igorralexsander-corp/echoapp/internal/app/model"
	"github.com/labstack/echo"
)

type Domain struct {
}

func NewDomainRoute() *Domain {
	return &Domain{}
}

func (r Domain) Register(e *echo.Echo) {
	e.GET("/v1/domain", r.GetDomain)
}

func (r *Domain) GetDomain(c echo.Context) error {
	domainResult := model.DomainMapping{
		OriginSystem:            "SYSTEM-A",
		OriginSystemDomainKey:   "GENERO",
		OriginSystemDomainValue: "MASCULINO",
		TargetSystem:            "SYSTEM-B",
		TargetSystemDomainValue: "MASC",
	}
	return c.JSON(200, domainResult)
}
