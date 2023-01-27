package controller

import (
	"gorm/core/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) SignUpController(c *gin.Context) {
	userObj := &models.SignUp{}
	if err := c.Bind(userObj); err != nil {
		c.JSON(412, gin.H{"message": "Invalid data"})
		return
	}
	if err := userObj.Validate(); err != nil {
		c.JSON(412, gin.H{"message": err.Error()})
		return
	}}
