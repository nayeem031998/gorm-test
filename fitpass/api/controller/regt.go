package controller

import (
	"github.com/gin-gonic/gin"
)

func (RD *Server) GetRegistartion(C *gin.Context) {
	var Result []map[string]interface{}
	RD.DB.Table("myschema.").Scan(&Result)
	C.JSON(200, Result)
}
