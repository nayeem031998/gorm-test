package controller

import (
	"car/core/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCarHandler(c *gin.Context) {
	var CarId int
	fmt.Println(c.Param("car_id"))
	if _, err := fmt.Sscan(c.Param("car_id"), &CarId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, Car := range model.Cars {
	if CarId == Car.Id {
			fmt.Println(model.Cars[:i], model.Cars[i+1:])
			model.Cars = append(model.Cars[:i], model.Cars[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}