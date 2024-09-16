package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// find the event
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event"})
		return
	}

	// Register for the event, probably come back to this later and check if the user is already registered
	// TODO: Check if user is already registered
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event"})
		return
	}

	// Send success message
	context.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully"})
}

func unregisterFromEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err := event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel registration."})
		return
	}

	// Send success message
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully"})
}
