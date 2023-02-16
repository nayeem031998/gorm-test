package main

import (
	//"fmt"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	// "fmt"
)

func main() {
	employees := []string{"rohit", "sanjay", "mohit", "sanjiv"}
	r := gin.Default()
	//READ
	r.GET("/get", func(c *gin.Context) {
		c.String(200, strings.Join(employees, ", "))

	})
	//CREATE
	r.POST("/add", func(c *gin.Context) {
		//get raw string from request body
		employee := c.PostForm("employee")
		//append to slice
		employees = append(employees, employee)

		c.JSON(200, gin.H{
			"message":   "Added 1 New Employee",
			"employees": strings.Join(employees, ", "),
		})
	})
	//UPDATE
	r.PUT("/update", func(c *gin.Context) {
		//update an employee by value
		fmt.Println(c.Query("new_value"))
		employee := c.Query("old_value")
		newemployee := c.Query("new_value")
		for i, v := range employees {
			if v == employee {
				employees[i] = newemployee
			}
		}
		c.JSON(200, gin.H{
			"message":   "Updated 1 Employee",
			"employees": strings.Join(employees, ", "),
		})
	})
	//DELETE
	r.DELETE("/delete", func(c *gin.Context) {
		//delete 1 employee by name
		employee := c.Query("employee")
		for i, v := range employees {
			if v == employee {
				employees = append(employees[:i], employees[i+1:]...)
			}
		}
		c.JSON(200, gin.H{
			"message":   "Deleted 1 Employee",
			"employees": strings.Join(employees, ", "),
		})
	})
	r.Run() // listen and serve on
}
