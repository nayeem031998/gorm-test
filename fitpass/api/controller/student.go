package controller

import (
	"github.com/gin-gonic/gin"
)

func (RD *Server) GetDetails(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Where("std_id=?", 1).Scan(&Student)
	C.JSON(200, Student)

}
func (RD *Server)GetNames(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Where("name=?", "asif").Scan(&Student)
	C.JSON(200, Student)
}
func (RD *Server)GetAge(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Where("age=?", 25).Scan(&Student)
	C.JSON(200, Student)
}
func (RD *Server)GetAddress(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Where("address=?", "assam").Scan(&Student)
	C.JSON(200, Student)
}
func (RD *Server) JoinTable(C *gin.Context) {
	var Student []map[string]interface{}
	RD.DB.Table("myschema.student").Joins("left join myschema.registration on myschema.student.std_id=myschema.registration.regt_id").Scan(&Student)
	C.JSON(200, Student)
}