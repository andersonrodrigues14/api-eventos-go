package main

import (
	"net/http"

	"github.com/andersonrodrigues14/api-eventos-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.POST("/events", createEvent)

	server.Run(":5059") //localhost:5059

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could nor parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event created! ", "event": event})
}
