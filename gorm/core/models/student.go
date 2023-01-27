package models

import (
	"errors"

	"gorm.io/gorm"
)

type Student struct {
	Id      int    `json:"id" gorm:"primary_key;column:std_id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
	Number  int    `json:"Number"`
	Email   string `json:"Email"`
	Date    int    `json:"Date"`
}

func (RD *Student) TableName() string {
	return "myschema.student"
}

func (RD *Student) GetStudentDetails(DB *gorm.DB) {
	DB.Debug().Model(RD).Find(&RD)
}

func (RD *Student) Validate() error {
	if RD.Name == " " {
		return errors.New("name is required")
	}
	if RD.Age == 0 {
		return errors.New("age is required")
	}
	if RD.Address == " " {
		return errors.New("address is required")
	}
	if RD.Number == 0 {
		return errors.New("number is required")
	}
	return nil
}

func (RD *Student) SaveStudent(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Create(&RD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
