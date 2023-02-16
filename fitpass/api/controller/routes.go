package controller

import "fitpass/api/middleware"

//import( "github.com/gin-gonic/gin"
// "strings"

func (s *Server) DBRoutes() {
	s.Router.GET("/students",middleware.SetMiddlewareJSON(), s.StudentController)
	s.Router.GET("/student",middleware.SetMiddlewareJSON(), s.GetDetails)
	s.Router.GET("/studentname",middleware.SetMiddlewareJSON(),s.GetNames)
	s.Router.GET("/studentage",middleware.SetMiddlewareJSON(), s.GetAge)
	s.Router.GET("/studentaddress",middleware.SetMiddlewareJSON(), s.GetAddress)	
	s.Router.GET("/join",middleware.SetMiddlewareJSON(),s.JoinTable)
	}
