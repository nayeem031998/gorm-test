package controller

import (
	"movie/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)
 func CreateCinemaHallHandler(c *gin.Context) {
 	var newCinemaHall model.CinemaHall
 	if err := c.ShouldBindJSON(&newCinemaHall); err != nil {
 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
 		return
 	}
 
 	if ok, str := newCinemaHall.ValidateCinemaHall(); !ok {
 		c.JSON(http.StatusBadRequest, gin.H{"error": str})
 		return
 	}
 	newCinemaHall.Id = len(model.CinemaHalls) + 1
 	model.CinemaHalls = append(model.CinemaHalls, newCinemaHall)
 	c.JSON(http.StatusOK, gin.H{"message": "CinemaHall Created successfully!", "cinemahall": newCinemaHall})
 }