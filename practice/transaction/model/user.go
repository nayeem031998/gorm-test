package user

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type User struct {
	gorm.Model
	Id      int32   `json:"id"`
	Name    string `json:"name"`
	Balance int  `json:"balance"`
}
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (T *User) TableName() string {
	return "bank.account"
}
 //func (T *User) GetUsers(c *gin.Context) {
	// var users []User
	 //T.DB.Table().Find(&users)
	 //c.JSON(200, users)
// }

	  
	 
