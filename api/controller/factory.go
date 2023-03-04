package controller

import (
	health "gefferson.com.br/geffws/api/controller/health"
	"gefferson.com.br/geffws/api/usecase"
)

type Controllers interface {
	NewPingController() health.PingController
}

func NewControllers(useCases usecase.UseCases) Controllers {
	return &controllersFactory{useCases}
}

type controllersFactory struct {
	useCases usecase.UseCases
}

func (c *controllersFactory) NewPingController() health.PingController {
	return health.NewPingController()
}
