package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Initialization() {

	server.Router = gin.New()
	server.Router.Use(gin.Logger())
	server.initializeRoutes()

}
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
