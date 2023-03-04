package router

import (
	"gefferson.com.br/geffws/api/controller"
	health "gefferson.com.br/geffws/api/router/health"
)

type Routes interface {
	NewPingRoutes() health.PingRoutes
}

func NewRoutes(controllers controller.Controllers) Routes {
	return &routesFactory{controllers}
}

type routesFactory struct {
	controllers controller.Controllers
}

func (r *routesFactory) NewPingRoutes() health.PingRoutes {
	return health.NewPingRoutes(r.controllers.NewPingController())
}
