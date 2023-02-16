package controller

import (
	"car/core/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Cars API",
	})
}

func GetCarsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Cars)
}

func GetCarHandler(c *gin.Context) {
	var CarId int
	fmt.Println(c.Param("Car_id"))
	if _, err := fmt.Sscan(c.Param("Car_id"), &CarId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, Car := range model.Cars {
		if Car.Id == CarId {
			c.JSON(http.StatusOK, Car)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
}
func GetCarByIdHandler(c *gin.Context) {
	var CarId int
	fmt.Println(c.Param("car_id"))
	if _, err := fmt.Sscan(c.Param("car_id"), &CarId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, Car := range model.Cars {
		if Car.Id == CarId {
			c.JSON(http.StatusOK, Car)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
}
func GetCarByBrandHandler(c *gin.Context) {
	var CarBrand string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_brand"))
	if _, err := fmt.Sscan(c.Param("car_brand"), &CarBrand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Brand == CarBrand {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
func GetCarByModelHandler(c *gin.Context) {
	var CarModel string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_model"))
	if _, err := fmt.Sscan(c.Param("car_model"), &CarModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Model == CarModel {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
func GetCarByYearHandler(c *gin.Context) {
	var CarYear int
	Cars := []model.Car{}
	fmt.Println(c.Param("car_year"))
	if _, err := fmt.Sscan(c.Param("car_year"), &CarYear); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Year == CarYear {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
func GetCarByMileageHandler(c *gin.Context) {
	var CarMileage int
	Cars := []model.Car{}
	fmt.Println(c.Param("car_mileage"))
	if _, err := fmt.Sscan(c.Param("car_mileage"), &CarMileage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Mileage == CarMileage {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}

func GetCarByNameHandler(c *gin.Context) {
	var CarName string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_name"))
	if _, err := fmt.Sscan(c.Param("car_name"), &CarName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Name == CarName {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}

func GetCarByPriceHandler(c *gin.Context) {
	var CarPrice int
	Cars := []model.Car{}
	fmt.Println(c.Param("car_price"))
	if _, err := fmt.Sscan(c.Param("car_price"), &CarPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Price == CarPrice {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
func GetCarByCapacityHandler(c *gin.Context) {
	var CarCapacity int
	Cars := []model.Car{}
	fmt.Println(c.Param("car_capacity"))
	if _, err := fmt.Sscan(c.Param("car_capacity"), &CarCapacity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Capacity == CarCapacity {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
func GetCarByFuelTypeHandler(c *gin.Context) {
	var CarFuelType string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_fueltype"))
	if _, err := fmt.Sscan(c.Param("car_fueltype"), &CarFuelType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.FuelType == CarFuelType {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}

func GetCarByColorHandler(c *gin.Context) {
	var CarColor string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_color"))
	if _, err := fmt.Sscan(c.Param("car_color"), &CarColor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Color == CarColor {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}




func GetCarByGearTypeHandler(c *gin.Context) {
	Cars := []model.Car{}
	var CarGearType string
	fmt.Println(c.Param("car_geartype"))
	if _, err := fmt.Sscan(c.Param("car_geartype"), &CarGearType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.GearType == CarGearType {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}

func GetCarByEngineHandler(c *gin.Context) {
	var CarEngine string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_engine"))
	if _, err := fmt.Sscan(c.Param("car_engine"), &CarEngine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.Engine == CarEngine {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}


func GetCarByCarStatusHandler(c *gin.Context) {
	var CarStatus string
	Cars := []model.Car{}
	fmt.Println(c.Param("car_status"))
	if _, err := fmt.Sscan(c.Param("car_status"), &CarStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Car := range model.Cars {
		if Car.CarStatus == CarStatus {
			Cars = append(Cars, Car)
		}
	}
	if len(Cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CAR not found"})
		return
	}
	c.JSON(http.StatusOK, Cars)
}
