package controller

import (
	"fmt"
	"movie/api/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MoviePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Movie API made by NAYEEM",
	})
}

func GetMoviesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Movies)
}

func GetMovieByIDHandler(c *gin.Context) {
	var MovieId int
	fmt.Println("movie api", c.Param("movie_id"))

	if _, err := fmt.Sscan(c.Param("movie_id"), &MovieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.ID == MovieId {
			c.JSON(http.StatusOK, Movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}

func GetMovieByTitleHandler(c *gin.Context) {
	var MovieTitle string
	fmt.Println(c.Param("movie_title"))
	if _, err := fmt.Sscan(c.Param("movie_title"), &MovieTitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.Title == MovieTitle {
			c.JSON(http.StatusOK, Movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}
func GetMovieByYearHandler(c *gin.Context) {
	var MovieYear int
	Movies := []model.Movie{}
	fmt.Println(c.Param("movie_year"))
	if _, err := fmt.Sscan(c.Param("movie_year"), &MovieYear); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.Year == MovieYear {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}

func GetMovieByGenreHandler(c *gin.Context) {
	var MovieGenre string
	Movies := []model.Movie{}
	fmt.Println(c.Param("movie_genre"))
	if _, err := fmt.Sscan(c.Param("movie_genre"), &MovieGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if MatchIndex(Movie.Genre, MovieGenre) {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}

func MatchIndex(str string, matchingStr string) bool {
	for index, value := range strings.Split(str, ",") {
		fmt.Println(matchingStr, value)
		if matchingStr == value {
			fmt.Println(index, value)
			return true
		}
	}
	return false
}

func GetMovieByCastHandler(c *gin.Context) {

	var MovieCast string
	Movies := []model.Movie{}
	if _, err := fmt.Sscan(c.Param("movie_cast"), &MovieCast); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		fmt.Println(Movie.ID)
		if MatchIndex(Movie.Cast, MovieCast) {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}
func GetMoviesByLanguageHandler(c *gin.Context) {
	var MovieLanguage string
	Movies := []model.Movie{}
	fmt.Println(c.Param("movie_language"))
	if _, err := fmt.Sscan(c.Param("movie_language"), &MovieLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.Language == MovieLanguage {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}
func GetMovieByBoxOfficeHandler(c *gin.Context) {
	var MovieBoxOffice string
	fmt.Println(c.Param("movie_boxoffice"))
	Movies := []model.Movie{}
	if _, err := fmt.Sscan(c.Param("movie_boxoffice"), &MovieBoxOffice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.BoxOffice == MovieBoxOffice {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}

func GetMovieByRatingHandler(c *gin.Context) {
	var MovieRating string
	Movies := []model.Movie{}
	fmt.Println(c.Param("movie_rating"))
	if _, err := fmt.Sscan(c.Param("movie_rating"), &MovieRating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.Rating == MovieRating {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)
}
func GetMovieByVerdictHandler(c *gin.Context) {
	var MovieVerdict string
	Movies := []model.Movie{}
	fmt.Println(c.Param("movie_verdict"))
	if _, err := fmt.Sscan(c.Param("movie_verdict"), &MovieVerdict); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, Movie := range model.Movies {
		if Movie.Verdict == MovieVerdict {
			Movies = append(Movies, Movie)
		}
	}
	if len(Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "MOVIE not found"})
		return
	}
	c.JSON(http.StatusOK, Movies)

}
