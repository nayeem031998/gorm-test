package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id      int    `json:"Id" gorm:"primary_key"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
	Number  int    `json:"Number"`
	Email   string `json:"Email"`
	Date    int    `json:"Date"`
}

func (RD *Student) GetStudentDetails(DB *gorm.DB) {
	var Student []map[string]interface{}
	DB.Debug().Table("myschema.student").Select("name", "age").Where("id = ?", 1).Scan(&Student)

}
