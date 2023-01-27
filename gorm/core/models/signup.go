package models

import "errors"

type SignUp struct {
	RegtId    int    `json:"regt_id" gorm:"primary_key"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	PhoneNo   int    `json:"phone_no"`
	Name      string `json:"Name"`
	Address   string `json:"address"`
	Course    string `json:"Course"`
}

func (RD *SignUp) Validate() error {
	if RD.Email == " " {
		return errors.New("email is required")
	}
	if RD.Password == " " {
		return errors.New("password is required")
	}
	if RD.PhoneNo == 0 {
		return errors.New("phone number is required")
	}
	if RD.Name == " " {
		return errors.New("name is required")
	}
	if RD.Address == " " {
		return errors.New("address is required") 
	}
	if RD.Course == " " {
		return errors.New("course is required")
	}
	if RD.Gender == " " {
		return errors.New("gender is required")
	}
	return nil
}

