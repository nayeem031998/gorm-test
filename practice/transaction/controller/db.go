package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
	
}

func (server *Server) DBInitialization() {
	dsn := "host=localhost user=postgres password=123456 dbname=transaction port=5432"
	var err error
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	server.Router = gin.New()
	server.Router.Use(gin.Logger())
	server.InitRouter()

}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
