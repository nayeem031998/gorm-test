package constant

import (

)
func GetSchema() (string, string, string) {
const (
DB_DRIVER   ="postgres"
DB_NAME     ="school"
DB_USER     ="postgres"
DB_PASSWORD ="123456"
DB_PORT ="5432"
DB_SSLMODE  ="disable"
SCHEMA      ="myschema"
student     = "myschema.student"
registration  = "myschema.registration"
	)
	return student, registration, SCHEMA
} 