package model
import(

)
type Registration struct {
	Regt_id		int64   `gorm:"primary_key;auto_increment" json:"regt_id"`
    Email		string  ` json:"email"`
	Password	string  ` json:"password"`
	Gendar		string  ` json:"gendar"`
	Name		string  ` json:"name"`	
}
func (r *Registration) TableName() string {
	return "myschema.registration"
}