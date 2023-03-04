package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "gefferson.com.br/geffws/api/router"
)

type GinServer interface {
	GetServer() *gin.Engine
	Server()
}

func NewGinServer(port string, routes route.Routes) *ginServer {
	server := gin.New()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		MaxAge:           12 * time.Hour,
	}))

	return &ginServer{server, port, routes}
}

type ginServer struct {
	server *gin.Engine
	port   string
	routes route.Routes
}

func (g *ginServer) GetServer() *gin.Engine {
	return g.server
}

func (g *ginServer) Server() {
	v1 := g.server.Group("/v1")
	{
		g.routes.NewPingRoutes().MakeGroup(v1)
	}

	err := g.server.Run(fmt.Sprintf(":%s", g.port))
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
