package model

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
)
 
type Car struct {
	Id int
	Name string
	GearType string
	Model string
	Brand string
	Capacity int
	FuelType string
	Mileage int
	Engine string
	Year int
	Color string
	Price int
	CarStatus string
}
var Cars = [] Car{
	{
		Id : 1,
		Name: "civic",
		GearType: "manual",   
		Model: "BS4",
		Brand: "honda",
		Capacity: 5,
		FuelType: "Petrol",
		Mileage: 17,
		Engine: "1200cc",
		Price: 1000000,
		Color: "Red",
		Year: 2015,
		CarStatus: "available",
	},
	{
		Id : 2,
		Name: "City",
		GearType: "manual",
		Model: "BS6",
		Brand: "honda",
		Capacity: 5,
		FuelType: "diesel",
		Mileage: 18,
		Engine: "1500cc",
		Price: 1500000,
		Color: "blue",
		CarStatus: "available",
		Year: 2016,
	},
	{
		Id : 3,
		Name: "swift",
		GearType: "manual",
		Model: "BS6",
		Brand: "suzuki",
		Capacity: 5,
		FuelType: "petrol",
		Mileage: 17,
		Engine: "1000cc",
		Price: 800000,
		Color: "gold",
		CarStatus: "unavailable",
		Year: 2022,

	},
	{
		Id : 4,
		Name: "Baleno",
		GearType: "automatic",
		Model: "BS6",
		Brand: "Nexa",
		Capacity: 6,
		FuelType: "petrol",
		Mileage: 20,
		Engine: "1000cc",
		Price: 800000,
		Color: "black",
		CarStatus: "unavailable",
		Year: 2023,
	},
	{
		Id : 5,
		Name: "verna",
		GearType: "manual",
		Model: "BS6",
		Brand: "hyundai",
		Capacity: 6,
		FuelType: "petrol",
		Mileage: 20,
		Engine: "1000cc",
		Color: "white",
		Price: 1200000,
		CarStatus: "available",
		Year: 2021,
	},
	{
		Id : 6,
		Name: "innova",
		GearType: "automatic",
		Model: "BS6",
		Brand: "toyota",
		Capacity: 10,
		FuelType: "petrol",
		Mileage: 13,
		Engine: "1000cc",
		Price: 1200000,
		CarStatus: "available",
		Color: "silver",
		Year: 2020,
	},
	{
		Id : 7,
		Name: "fortuner",
		GearType: "automatic_and_manual",
		Model: "BS6",
		Brand: "toyota",
		Capacity: 12,
		FuelType: "diesel",
		CarStatus: "waiting",
		Mileage: 12,
		Engine: "2000cc",
		Color: "green",
		Price: 4000000,
		Year: 2022,
	},
	{
		Id : 8,
		Name: "scorpio",
		Model: "BS6",
		GearType: "manual",
		Brand: "mahindra",
		Capacity: 8,
		FuelType: "petrol",
		Mileage: 15,
		Engine: "1600cc",
		Price: 2400000,
		CarStatus: "available_in_2_days",
		Color: "yellow",
		Year: 2023,
	},
	{
		Id : 9,
		Name: "XUV",
		GearType: "manual",
		Model: "BS6",
		Brand: "mahindra",
		Capacity: 9,
		FuelType: "diesel",
		Mileage: 14,
		Engine: "1800cc",
		Price: 2400000,
		CarStatus: "available_in_next_month",
		Color: "Grey",
		Year: 2019,
	},
	{
		Id : 10,
		Name: "ertiga",
		GearType: "electric",
	    Color: "Red",
		Model: "BS6",
		Brand: "maruti",
		Capacity: 7,
		FuelType: "petrol",
		Mileage: 20,
		Engine: "1300cc",
		CarStatus: "out_of_stock",
		Price: 1100000,
		Year: 2023,
	},

}
func ValidateCar(car Car) bool {
	if car.Name == "" || car.Model == "" || car.Brand == "" || car.Capacity == 0 || car.FuelType == "" || car.Mileage == 0 || car.Engine == "" || car.Year == 0 {
		return false
	}
	return true
}
func GetCarById(id int) *Car {
	for _, car := range Cars {
		if car.Id == id {
			return &car
		}
	}
	return nil
}
func ValidateMovie(car Car) (bool,int, string) {
    if car.Id == 0 {
		return false, 400, "Id is required"
	}
	if car.Name == "" {
		return false, 400, "Name is required"
	}
	if car.Model == "" {
		return false, 400, "Model is required"
	}
	if car.Brand == "" {
		return false, 400, "Brand is required"
	}
	if car.Capacity == 0 {
		return false, 400, "Capacity is required"
	}
	if car.FuelType == "" {
		return false, 400, "FuelType is required"
	}
	if car.Mileage == 0 {
		return false, 400, "Mileage is required"
	}
	if car.Engine == "" {
		return false, 400, "Engine is required"
	}
	if car.Year == 0 {
		return false, 400, "Year is required"
	}
   if car.Price == 0 {
		return false, 400, "Price is required"
   }
   if car.Color == "" {
		return false, 400, "Color is required"
		   }

	return true, 200, "success"
}





