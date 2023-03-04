package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	health "gefferson.com.br/geffws/api/usecase/health"
)

func NewPingController() PingController {
	return &pingController{}
}

type PingController interface {
	Ping(c *gin.Context)
}

type pingController struct{}

// Health godoc
// @Tags         health
// @Summary Get api health
// @Produce      json
// @Success      200
// @Router /ping [get]
// @Security JWT
func (pc *pingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, health.GetPingResponse())
}
