package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)
func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Tech Company API made by NAYEEM"})
}
type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	CEO     string `json:"ceo"`
	Revenue string `json:"revenue"`
}
var companies = []Company{
	{
		ID:      "1",
		Name:    "FITPASS",
		CEO:     "AKSHAY VERMA SIR",
		Revenue: "2 billion",
	},
	{
		ID:      "2",
		Name:    "FITFEAST",
		CEO:     "ARUSHI MAM",
		Revenue: "1 billion",
	},
	{
		ID:      "3",
		Name:    "FITCASH",
		CEO:     "GYAN SIR",
		Revenue: "90 million",
	},
}
func  IndexController(c *gin.Context) {
	AddForm := []byte(`
	<html>
	<h3 style="top: 50%;
		text-align: center;
		margin-top: 10%;
		font-size: 90px;"><a href="https://journeyblog99.blogspot.com/" style="text-decoration: none"><span style="    color: #ef4c4f;">WELCOME TO FITPASS</span></a></h3>

		<a href="http://localhost:8080/companies"><button style="text-decoration: none">companies</botton></a>
	`)
	ContentTypeHTML := "text/html; charset=utf-8"
	c.Data(http.StatusOK, ContentTypeHTML, AddForm)
}
func ValidateCompany(company Company) bool {
	if company.ID== "" || company.Name == "" || company.CEO == "" || company.Revenue == "" {
		return false
	}
	return true
}
func ValidateCompanyID(companyID string) bool {
	for _, company := range companies {
		if company.ID == companyID {
			return true
		}
	}
	return false
}
func Validate(c *gin.Context) {
	companyID := c.Param("company_id")
	if !ValidateCompanyID(companyID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}
	c.Next()
}

func ValidateCompanyData(c *gin.Context) {
	var company Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !ValidateCompany(company) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company data"})
		return
	}
	c.Next()
}

func GetCompaniesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, companies)
}
func CreateCompanyHandler(c *gin.Context) {
	var newCompany Company
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCompany.ID = fmt.Sprint(len(companies) + 1)
	companies = append(companies, newCompany)
	c.JSON(http.StatusCreated, newCompany)
}
func UpdateCompanyHandler(c *gin.Context) {
	fmt.Println(c.Param("company_id"))
	companyID := c.Param("company_id")

	var updatedCompany Company
	if err := c.ShouldBindJSON(&updatedCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, company := range companies {
		if company.ID == companyID {
			companies[i] = updatedCompany
			c.JSON(http.StatusOK, updatedCompany)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
}
func DeleteCompanyHandler(c *gin.Context) {
	companyID := c.Param("company_id")
	for i, company := range companies {
		if company.ID == companyID {
			fmt.Println(companies[:i], companies[i+1:])
			companies = append(companies[:i], companies[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Company not found"})
}
func main() {
	r := gin.Default()
	r.GET("/home", IndexController)
	r.GET("/", HomepageHandler)
	r.GET("/companies", GetCompaniesHandler)
	r.POST("/add", CreateCompanyHandler)
	r.PUT("/companies/:company_id", UpdateCompanyHandler)
	r.DELETE("/companies/:company_id", DeleteCompanyHandler)
	r.Run()
}
