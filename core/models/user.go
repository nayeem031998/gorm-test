package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      int    `json:"Id" gorm:"primary_key"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
	Salary  int    `json:"Salary"`
	Branch  string `json:"Branch"`
	Date    int    `json:"Date"`
}
