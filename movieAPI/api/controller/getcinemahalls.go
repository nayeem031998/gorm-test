package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)
func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to the Cinema Hall API"})
}
func GetCinemaHallsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cinemahalls": model.CinemaHalls})
}
func GetCinemaHallsByIdHandler(c *gin.Context) {
	var CinemaHallId int
	if _, err := fmt.Sscan(c.Param("cinemahall_id"), &CinemaHallId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Id == CinemaHallId {
			c.JSON(http.StatusOK, CinemaHall)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
}
func GetCinemaHallByNameHandler(c *gin.Context) {
	var CinemaHallName string
	if _, err := fmt.Sscan(c.Param("cinemahall_name"), &CinemaHallName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Name == CinemaHallName {
			c.JSON(http.StatusOK, CinemaHall)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
}
func GetCinemHallByCityHandler(c *gin.Context) {
	var CinemaHallCity string
	Halls := []model.CinemaHall{}
	fmt.Println("CinemaHallCity", CinemaHallCity)
	if _, err := fmt.Sscan(c.Param("cinemahall_city"), &CinemaHallCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.City == CinemaHallCity {
            Halls = append(Halls, CinemaHall)
        }
    }
if len(Halls) == 0 {
	c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
    return
}	
	c.JSON(http.StatusOK, Halls)
} 
func GetCinemaHallMovieNameByHandler(c *gin.Context) {
	var CinemaHallMovieName string
	Halls := []model.CinemaHall{}
	fmt.Println("CinemaHallMovieName", CinemaHallMovieName)
	if _, err := fmt.Sscan(c.Param("cinemahall_movie_name"), &CinemaHallMovieName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Movie_Name == CinemaHallMovieName {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
func GetCinemaHallByFacilitiesHandler(c *gin.Context) {
	var CinemaHallFacilities string
	Halls := []model.CinemaHall{}
	fmt.Println("CinemaHallFacilities", CinemaHallFacilities)
	if _, err := fmt.Sscan(c.Param("facilities"), &CinemaHallFacilities); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Facilities == CinemaHallFacilities {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
func GetCinemHallByParkingHandler(c *gin.Context) {
	var CinemaHallParking string
	Halls := []model.CinemaHall{}
	fmt.Println("CinemaHallParking", CinemaHallParking)
	if _, err := fmt.Sscan(c.Param("cinemahall_parking"), &CinemaHallParking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Parking == CinemaHallParking {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
func GetCinemHallByTypeHandler(c *gin.Context) {
	var CinemaHallType string
	Halls := []model.CinemaHall{}
	if _, err := fmt.Sscan(c.Param("cinemahall_type"), &CinemaHallType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Type == CinemaHallType {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
func GetCinemHallByCapacityHandler(c *gin.Context) {
	var CinemaHallCapacity int
	Halls := []model.CinemaHall{}
	if _, err := fmt.Sscan(c.Param("cinemahall_capacity"), &CinemaHallCapacity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Capacity == CinemaHallCapacity {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
func GetCinemaHallsByTimeHandler(c *gin.Context) {
	var CinemaHallTime string
	Halls := []model.CinemaHall{}
	if _, err := fmt.Sscan(c.Param("cinemahall_time"), &CinemaHallTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	for _, CinemaHall := range model.CinemaHalls {
		if CinemaHall.Time == CinemaHallTime {
			Halls = append(Halls, CinemaHall)
		}
	}
	if len(Halls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
		return
	}
	c.JSON(http.StatusOK, Halls)
}
