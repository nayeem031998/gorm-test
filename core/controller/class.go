package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) ClassController(c *gin.Context) {
	var class []map[string]interface{}
	server.DB.Table("myschema.class").Scan(&class)
	c.JSON(200, class)

}
