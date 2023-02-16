package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) IndexController(c *gin.Context) {
	AddForm := []byte(`
	<html>
	<style>
	body{
		background-image:"/Users/nayeem/Downloads/IMG_20211116_225003_561.webp"
	}
	</style>
	<body>
	<img src=image:"Users/nayeem/Downloads/IMG_20211116_225003_561.webp">

	</body>
	<h3 style="top: 50%;
		text-align: center;
		margin-top: 10%;
		font-size: 90px;"><a href="https://journeyblog99.blogspot.com/" style="text-decoration: none"><span style="    color: #ef4c4f;">WELCOME TO SCHOOL</span></a></h3>
     
		<a href="http://localhost:8084/student"><button style="text-decoration: none">student</botton></a>
		<a href="http://localhost:8084/class"><button style="text-decoration: none">class</botton></a>
		<a href="http://localhost:8084/subject"><button style="text-decoration: none">subject</botton></a>
		</html>`)

	ContentTypeHTML := "text/html; charset=utf-8"
	c.Data(http.StatusOK, ContentTypeHTML, AddForm)
	return
}
