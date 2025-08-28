package main

import (
	"github.com/andersonrodrigues14/api-eventos-go/db"
	"github.com/andersonrodrigues14/api-eventos-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":5059") //localhost:5059

}
