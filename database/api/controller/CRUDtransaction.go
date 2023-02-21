package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"database/api/model"
)

func (server *Server) CreateTransactions(c *gin.Context) {
	var ( 
		transaction model.Transaction
	)
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := transaction.ValidateTransaction(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := transaction.SaveTransaction(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Created Successfully", "data": transaction})

}
func (server *Server) UpdateTransactions(c *gin.Context) {
	var (
		transaction model.Transaction
	)
	id := c.Param("id")
	if err := server.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := transaction.ValidateTransaction(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := transaction.UpdateTransaction(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Updated Successfully", "data": transaction})
}
func (server *Server) DeleteTransactions(c *gin.Context) {
	var (
		transaction model.Transaction
	)
	id := c.Param("id")
	if err := server.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := transaction.DeleteTransaction(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Deleted Successfully", "data": transaction})
}
func (server *Server) GetTransactions(c *gin.Context) {
	var (
		transaction model.Transaction
	)
	id := c.Param("id")
	if err := server.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Retrieved Successfully", "data": transaction})
}
func (server *Server) GetAllTransactions(c *gin.Context) {
	var (
		transaction []model.Transaction
	)
	if err := server.DB.Scan(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transactions Retrieved Successfully", "data": transaction})
}



