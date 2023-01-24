package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) StudentController(c *gin.Context) {
	var student []map[string]interface{}
	server.DB.Table("myschema.student").Scan(&student)
	c.JSON(200, student)
}
