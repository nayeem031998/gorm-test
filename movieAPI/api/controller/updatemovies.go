package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

	func UpdateMovieHandler(c *gin.Context) {
		var MovieId int
		fmt.Println("movie api", c.Param("movie_id"))
		if _, err := fmt.Sscan(c.Param("movie_id"), &MovieId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		var updatedMovie model.Movie
		if err := c.ShouldBindJSON(&updatedMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if ok, str := updatedMovie.ValidateMovie(); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": str})
			return
		}
		for i, Movie := range model.Movies {
			if Movie.ID == MovieId {
				model.Movies[i] = updatedMovie
				c.JSON(http.StatusOK, updatedMovie)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	}