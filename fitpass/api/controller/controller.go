package controller

import (
	"github.com/gin-gonic/gin"
)

// studentcontroller
func (RD *Server) StudentController(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Scan(&Student)
	C.JSON(200, Student)
}
