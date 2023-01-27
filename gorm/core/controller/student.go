package controller

import (
	"gorm/core/models"

	"github.com/gin-gonic/gin"
)

func (server *Server) StudentController(c *gin.Context) {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		c.JSON(412, gin.H{"message": "Invalid data"})
		return
	}
	if err := student.Validate(); err != nil {
		c.JSON(412, gin.H{"message": err.Error()})
		return
	}
	if err := student.SaveStudent(server.DB); err != nil {
		c.JSON(412, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, student)

}
