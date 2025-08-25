package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.Run(":5059") //localhost:5059

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "HELLO!"})
}
