package models

func GetSchema() (string, string, string) {
	const (
		student = "myschema.student"
		class   = "myschema.class"
		subject = "myschema.subject"
	)
	return student, class, subject
} 
