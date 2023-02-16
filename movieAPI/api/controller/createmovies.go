package controller

import (
	//"fmt"
	"movie/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMovieHandler(c *gin.Context) {
	var newMovie model.Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, str := newMovie.ValidateMovie(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": str})
		return
	}
	newMovie.ID = len(model.Movies) + 1
	model.Movies = append(model.Movies, newMovie)
	c.JSON(http.StatusOK, gin.H{"message": "Movie Created successfully!", "movie": newMovie})
}
