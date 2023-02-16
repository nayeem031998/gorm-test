package controller

//import //("movie/api/middleware"

//import// "movie/api/middleware"

//"movie/api/model"
//"net/http"
//import
 //"github.com/gin-gonic/gin"
//)

func (s *Server) initializeRoutes() {
	s.Router.GET("/", IndexController)
	s.Router.GET("/home", MoviePage)
	s.Router.GET("/movies", GetMoviesHandler)
	s.Router.POST("/add", CreateMovieHandler)
	s.Router.PUT("/movies/update/:movie_id", UpdateMovieHandler)
	s.Router.DELETE("/movies/delete/:movie_id",DeleteMovieHandler)
	s.Router.GET("/movies/get/:movie_id", GetMovieByIDHandler)
	s.Router.GET("/movies/title/:movie_title", GetMovieByTitleHandler)
	s.Router.GET("/movies/year/:movie_year", GetMovieByYearHandler)
	s.Router.GET("/movies/genre/:movie_genre", GetMovieByGenreHandler)
	s.Router.GET("/movies/language/:movie_language", GetMoviesByLanguageHandler)
	s.Router.GET("/movies/boxoofice/:movie_boxoffice", GetMovieByBoxOfficeHandler)
	s.Router.GET("/movies/rating/:movie_rating", GetMovieByRatingHandler)
	s.Router.GET("/movies/verdict/:movie_verdict", GetMovieByVerdictHandler)
	s.Router.GET("/movies/cast/:movie_cast", GetMovieByCastHandler)
	s.Router.GET("/cinemahall/id/:cinemahall_id", GetCinemaHallsByIdHandler)
	s.Router.GET("/cinemahall/name/:cinemahall_name", GetCinemaHallByNameHandler)
	s.Router.GET("/cinemahall/city/:cinemahall_city", GetCinemHallByCityHandler)
	s.Router.GET("/cinemahall/movie_name/:cinemahall_movie_name", GetCinemaHallMovieNameByHandler)
	s.Router.GET("/cinemahall/parking/:cinemahall_parking", GetCinemHallByParkingHandler)
	s.Router.GET("/cinemahall/type/:cinemahall_type", GetCinemHallByTypeHandler)
	s.Router.GET("/cinemahall/capacity/:cinemahall_capacity", GetCinemHallByCapacityHandler)
	s.Router.GET("/cinemahall/time/:cinemahall_time",GetCinemaHallsByTimeHandler)
	s.Router.GET("/cinemahall/facilities/:facilities",GetCinemaHallByFacilitiesHandler)
	s.Router.POST("/create",CreateCinemaHallHandler)
	s.Router.PUT("/update/:cinemahall_id", UpdateCinemaHallHandler)
	s.Router.DELETE("/delete/:cinemahall_id", DeleteCinemaHallHandler)
	s.Router.GET("/page", HomePageHandler)
	s.Router.GET("/all", GetCinemaHallsHandler)
}
