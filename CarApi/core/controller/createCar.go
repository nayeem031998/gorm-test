package controller

import (
	"car/core/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCarHandler(c *gin.Context) {
	var newCar model.Car
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCar.Id = len(model.Cars) + 1
	model.Cars = append(model.Cars, newCar)
	c.JSON(http.StatusCreated, newCar)
}
