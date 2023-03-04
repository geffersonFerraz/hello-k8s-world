package main

import (
	"fmt"
	"os"

	"gefferson.com.br/geffws/api/controller"
	"gefferson.com.br/geffws/api/router"
	"gefferson.com.br/geffws/api/usecase"
	"gefferson.com.br/geffws/server"
)

func main() {
	fmt.Println("Starting server... Port: ", os.Getenv("PORT"))

	useCases := usecase.NewUseCases()
	controllers := controller.NewControllers(useCases)
	routes := router.NewRoutes(controllers)

	server := server.NewGinServer(os.Getenv("PORT"), routes)

	server.Server()

}
