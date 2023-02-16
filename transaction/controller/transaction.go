package controller

import (
	"fmt"
	//"os/user"

	//"transaction/model"

	"github.com/gin-gonic/gin"
)

func (server *Server) TransactionController(c *gin.Context) {
	fmt.Println("TransactionController")
	

var account []map[string]interface{}
server.DB.Table("bank.account").Order(" id asc,name").Scan(&account)
c.JSON(200,account)


}


