package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)
func UpdateCinemaHallHandler(c *gin.Context) {
	var CinemaHallId int
	if _, err := fmt.Sscan(c.Param("cinemahall_id"), &CinemaHallId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updatedCinemaHall model.CinemaHall
	if err := c.ShouldBindJSON(&updatedCinemaHall); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, str := updatedCinemaHall.ValidateCinemaHall(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": str})
		return
	}
	for i, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Id == CinemaHallId {
			model.CinemaHalls[i] = updatedCinemaHall
			c.JSON(http.StatusOK, updatedCinemaHall)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
}
