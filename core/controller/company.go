package controller

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) companyController(c *gin.Context) {
	var company []map[string]interface{}
	server.DB.Table("company").Scan(&company)
	c.JSON(200, company)

}
