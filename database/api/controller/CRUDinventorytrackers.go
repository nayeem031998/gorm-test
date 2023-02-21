package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"database/api/model"
)

func (server *Server) CreateInventoryTrackers(c *gin.Context) {
	var (
		inventorytracker model.InventoryTracker
	)
	if err := c.ShouldBindJSON(&inventorytracker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := inventorytracker.ValidateInventoryTracker(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := inventorytracker.SaveInventoryTracker(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "InventoryTracker Created Successfully", "data": inventorytracker})

}
func (server *Server) UpdateInventoryTrackers(c *gin.Context) {
	var (
		inventorytracker model.InventoryTracker
	)
	id := c.Param("id")
	if err := server.DB.First(&inventorytracker, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&inventorytracker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := inventorytracker.ValidateInventoryTracker(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := inventorytracker.UpdateInventoryTrackers(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "InventoryTracker Updated Successfully", "data": inventorytracker})
}
func (server *Server) DeleteInventoryTrackers(c *gin.Context) {
	var (
		inventorytracker model.InventoryTracker
	)
	id := c.Param("id")
	if err := server.DB.First(&inventorytracker, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := inventorytracker.DeleteInventoryTrackers(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "InventoryTracker Deleted Successfully", "data": inventorytracker})
}

func (server *Server) GetInventoryTrackers(c *gin.Context) {
	var (
		inventorytracker []model.InventoryTracker
	)
	id := c.Param("id")
	if err := server.DB.First(&inventorytracker, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(inventorytracker) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "InventoryTracker Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "InventoryTracker Found Successfully", "data": inventorytracker})
}
func (server *Server) GetAllInventoryTrackers(c *gin.Context) {
	var (
		inventorytracker model.InventoryTracker
	)
	if err := server.DB.Find(&inventorytracker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "InventoryTracker Found Successfully", "data": inventorytracker})
}

