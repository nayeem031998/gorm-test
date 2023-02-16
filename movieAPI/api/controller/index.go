package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func IndexController(c *gin.Context) {
	AddForm := []byte(`
	<html>
	<h3 style="top: 50%;
		text-align: center;
		margin-top: 10%;
		font-size: 90px;"><a href="https://youtube.com/@nayeembagpacker/" style="text-decoration: none"><span style="    color: #ef4c4f;">WELCOME TO NETFLIX</span></a></h3>

		<a href="http://localhost:1234/movies"><button style="text-decoration: none">movies</botton></a>
		`)

	ContentTypeHTML := "text/html; charset=utf-8"
	c.Data(http.StatusOK, ContentTypeHTML, AddForm)
	
}