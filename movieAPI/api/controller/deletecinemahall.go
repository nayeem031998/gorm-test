package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCinemaHallHandler(c *gin.Context){
	var (
		CinemaHallId int64
		err error
    )

	if CinemaHallId, err = strconv.ParseInt(c.Param("cinemahall_id"), 0, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("CinemaHall api", CinemaHallId)
	for i, cinemahall := range model.CinemaHalls {
		if int64(cinemahall.Id) == int64(CinemaHallId){
			model.CinemaHalls = append(model.CinemaHalls[:i], model.CinemaHalls[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "CinemaHall deleted successfully!"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "CinemaHall not found"})
}




