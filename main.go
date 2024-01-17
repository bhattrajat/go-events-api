package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bhattrajat/go-events-api/db"
	"github.com/bhattrajat/go-events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:eventId", getEvent)
	r.POST("/events", addEvent)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
		return
	}
	fmt.Println(events)
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func getEvent(c *gin.Context) {
	eventIdParam := c.Param("eventId")
	eventId, err := strconv.ParseInt(eventIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No event found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func addEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid/Missing fields in request body",
		})
		return
	}
	event.ID = 1
	event.UserId = 1
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error adding event",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"event": event,
	})
}
