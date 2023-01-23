package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) UserController(c *gin.Context) {
	var student []map[string]interface{}
	server.DB.Table("nayeem.student").Scan(&student)
	c.JSON(200, student)

}
