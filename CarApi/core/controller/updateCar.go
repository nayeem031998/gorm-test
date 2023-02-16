package controller

import (
	"fmt"
	"net/http"

	"car/core/model"

	"github.com/gin-gonic/gin"
)

func UpdateCarHandler(c *gin.Context) {
	var CarId int
	fmt.Println("car api", c.Param("car_id"))
	if _, err := fmt.Sscan(c.Param("car_id"), &CarId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedCar model.Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, Car := range model.Cars {
		if Car.Id == CarId {
			model.Cars[i] = updatedCar
			c.JSON(http.StatusOK, updatedCar)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}
