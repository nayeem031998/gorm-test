package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMovieHandler(c *gin.Context) {
	var (
		MovieId int64
		err     error
	)
	if MovieId, err = strconv.ParseInt(c.Param("movie_id"), 0, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.S
	fmt.Println("movie api", MovieId)
	for i, Movie := range model.Movies {
		if int64(Movie.ID) == MovieId {
			model.Movies = append(model.Movies[:i], model.Movies[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully!"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}
