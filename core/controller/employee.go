package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) employeeController(c *gin.Context) {
	var employee []map[string]interface{}
	server.DB.Table("employee").Scan(&employee)
	c.JSON(500, employee)
	
}
