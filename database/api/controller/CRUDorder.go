package controller

import (
	"database/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateOrder(c *gin.Context) {
	var (
		order model.Orders
		err   error
	)
	if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.OrdersStuct(c)
	if err = order.ValidateOrders(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = order.SaveOrder(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Order Created Successfully", "data": order})
}
func (server *Server) UpdateOrders(c *gin.Context) {
	var (
		order model.Orders
		err   error
	)
	id := c.Param("id")
	if err = server.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.OrdersStuct(c)
	if err = order.ValidateOrders(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = order.UpdateOrder(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order Updated Successfully", "data": order})
}

func (server *Server) DeleteOrders(c *gin.Context) {
	var (
		order model.Orders
		err   error
	)
	id := c.Param("id")
	if err = server.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = order.DeleteOrder(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order Deleted Successfully", "data": order})
}
func (server *Server) GetOrders(c *gin.Context) {
	var (
		orders model.Orders
		err    error
	)
	// customerId := c.Request.Header["User-Id"][0]
	if err = server.DB.Debug().Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}
func (server *Server) GetOrder(c *gin.Context) {
	var (
		order map[string]interface{}
		err   error
	)
	id := c.Param("id")

	if err = server.DB.Debug().Table("orders ord").Joins("Left join products pro ON ord.product_id = pro.product_id").Where("ord.order_id=?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(order) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Order Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
func (server *Server) GetOrderById(c *gin.Context) {

	var (
		order []model.Orders
		err   error
	)
	id := c.Param("id")
	if err = server.DB.Debug().First(&order, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(order) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Order Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
