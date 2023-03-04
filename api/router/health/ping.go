package router

import (
	controller "gefferson.com.br/geffws/api/controller/health"
	"github.com/gin-gonic/gin"
)

func NewPingRoutes(controller controller.PingController) PingRoutes {
	return &pingRoutes{controller}
}

type PingRoutes interface {
	MakeGroup(router *gin.RouterGroup)
}

type pingRoutes struct {
	controller controller.PingController
}

func (c *pingRoutes) MakeGroup(router *gin.RouterGroup) {
	group := router.Group("ping")
	{
		group.GET("", c.controller.Ping)
	}
}
