package controller

import (
	"database/api/model"
	//"fmt"
	"net/http"

	//"strings"

	//"time"

	//"strconv"
	"github.com/gin-gonic/gin"
)

// create product in formdata and json format
func (server *Server) CreateProduct(c *gin.Context) {
	var (
		product model.Products
		err     error
	)
	if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ProductsStuct(c)
	if err = product.ValidateProducts(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = product.SaveProduct(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Created Successfully", "data": product})
}

// update product in formdata and json format
func (server *Server) GetProducts(c *gin.Context) {
	var (
		products []model.Products
		err      error
	)
	if err = server.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}
func (server *Server) GetProduct(c *gin.Context) {
	var (
		product model.Products
		err     error
	)
	id := c.Param("id")
	if err = server.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
func (server *Server) UpdateProducts(c *gin.Context) {
	var (
		product model.Products
		err     error
	)
	id := c.Param("id")
	if err = server.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ProductsStuct(c)
	if err = product.ValidateProducts(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = product.UpdateProduct(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Product Updated Successfully", "product": product})

}
func (server *Server) DeleteProducts(c *gin.Context) {
	var (
		product model.Products
		err     error
	)
	id := c.Param("id")
	if err = server.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = product.DeleteProduct(server.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Product Deleted Successfully", "product": product})
}
func (server *Server) GetProductsByName(c *gin.Context) {
	var (
		products []model.Products
		err      error
	)
	name := c.Param("name")
	if err = server.DB.Where("product_name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}












