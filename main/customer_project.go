package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}

type employees []Employee

func allEmployees(w http.ResponseWriter, r *http.Request) {
	employees := employees{
		Employee{Name: "Nayeem", Age: 25, Address: "golaghat", Salary: 25000},
		Employee{Name: "Asif", Age: 24, Address: "Jorhat", Salary: 50000},
		Employee{Name: "SANDEEP sir ", Age: 28, Address: "Delhi", Salary: 60000},
		Employee{Name: "Gyan sir ", Age: 45, Address: "up", Salary: 100000},
		Employee{Name: "Rohit sir ", Age: 35, Address: "up", Salary: 70000},
		Employee{Name: "Ravi sir ", Age: 40, Address: "punjab", Salary: 90000},
	}
	fmt.Println("Endpoint Hit: All employees Endpoint")
	json.NewEncoder(w).Encode(employees)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to FITPASS")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/employees", allEmployees)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
