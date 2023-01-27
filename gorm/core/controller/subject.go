package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) SubjectController(c *gin.Context) {
	var subject []map[string]interface{}
	server.DB.Table("myschema.subject").Scan(&subject)
	c.JSON(500, subject)

}
