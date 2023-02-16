package controller

import (
	//"github.com/gin-gonic/gin"
	//"net/http"
	//"fmt"
)
func (s *Server) initializeRoutes() {
	r:=s.Router
	r.GET("/", HomeHandler)
	r.GET("/cars", GetCarsHandler)
	r.GET("/cars/:car_id", GetCarHandler)
	r.POST("/add", CreateCarHandler)
	r.DELETE("/cars/:car_id", DeleteCarHandler)
	r.PUT("/cars/:car_id", UpdateCarHandler)
	r.GET("/cars/name/:car_name",GetCarByNameHandler)
	r.GET("/cars/brand/:car_brand", GetCarByBrandHandler)
	r.GET("/cars/model/:car_model", GetCarByModelHandler)
	r.GET("/cars/year/:car_year", GetCarByYearHandler)
	r.GET("/cars/price/:car_price", GetCarByPriceHandler)
	r.GET("/cars/id/:car_id", GetCarByIdHandler)
	r.GET("/cars/mileage/:car_mileage", GetCarByMileageHandler)
	r.GET("/cars/color/:car_color", GetCarByColorHandler)	
	r.GET("/cars/engine/:car_engine", GetCarByEngineHandler)
	r.GET("/cars/geartype/:car_geartype", GetCarByGearTypeHandler)
	r.GET("/cars/capacity/:car_capacity", GetCarByCapacityHandler)
	r.GET("/cars/fueltype/:car_fueltype", GetCarByFuelTypeHandler)
	r.GET("/cars/carstatus/:car_status", GetCarByCarStatusHandler)
}