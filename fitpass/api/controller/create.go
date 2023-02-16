package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
)
func CreateStudent(){
students := []string{"sagar","sourav"}
    r := gin.Default()
	//READ
	r.GET("/get", func(c *gin.Context) {
		c.String(200, strings.Join(students, ", "))
	})
	//CREATE
	r.POST("/add", func(c *gin.Context) {
		//get raw string from request body
		student := c.PostForm("student")
		//append to slice
		students = append(students, student)

		c.JSON(200, gin.H{
			"message":   "Added 1 New Student",
			"employees": strings.Join(students, ", "),
		})
	})
}