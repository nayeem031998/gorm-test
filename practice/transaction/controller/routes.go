package controller

import (
	"github.com/gin-gonic/gin"
)



func (r *Server) InitRouter() {
	r.Router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my database",
		})
	})
	r.Router.GET("/transaction", r.TransactionController)
}
