package models
import(
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Id    int    `json:"Id" gorm:"primary_key"`
	Name  string `json:"Name"`
	Class int    `json:"class"`
}