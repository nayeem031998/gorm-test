package models
import(
	"gorm.io/gorm"
)


type Class struct {
	gorm.Model
	Id         int    `json:"Id" gorm:"primary_key"`
	Chair      int    `json:"Chair"`
	Tables     int    `json:"Tables"`
	Whiteboard int    `json:"white_board"`
	Standard   string `json:"Standard"`
}