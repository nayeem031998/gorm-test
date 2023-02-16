package model

import "time"
type Student struct {
	Std_id        uint64   `gorm:"primary_key;auto_increment" json:"std_id"`
	Name          string   ` json:"name"`
	Age           int      ` json:"age"`
	Address       string   ` json:"address"`
	DOB 		 time.Time   ` json:"DOB"`
	Number		int64  ` json:"number"`
	Email		string  ` json:"email"`
	Amount		int64  ` json:"amount"`
}
func (s *Student) TableName() string {
	return "myschema.student"
}

func (s *Student) GetDetails() []Student {
	var Student []Student
	//RD.db.Table("myschema.student").Where("std_id=?", 1).Scan(&Student)
	return Student
}
func (s *Student) GetNames() []Student {
	var Student []Student
	//RD.db.Table("myschema.student").Where("name=?", "nayeem").Scan(&Student)
	return Student
}
func (s *Student) GetAge() []Student {
	var Student []Student
	//RD.db.Table("myschema.student").Where("age=?", 22).Scan(&Student)
	return Student
}
func (s *Student) GetAddress() []Student {
	var Student []Student
	return Student
}
