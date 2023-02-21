package model

import (
	"database/api/utils/commonfunction"
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Products struct {
	ProductId     uint    `gorm:"primary_key;auto_increment" json:"product_id"`
	ProductName   string  `json:"product_name"`
	ProductPrice  float64 `json:"product_price"`
	ProductImage  string  `json:"product_image"`
	ProductType   int     `json:"product_type"`
	ProductWeight string  `json:"product_weight"`
	ProductBrand  string  `json:"product_brand"`
	TaxPercentage float64 `json:"tax_percentage"`
	ProductStatus string  `json:"product_status"`
}

func (PD *Products) TableName() string {
	return "products"
}
func (PRD *Products) ProductsStuct(c *gin.Context) *Products {
	// fmt.Println(c.Request.FormValue("product_name"))
	PRD.ProductName = c.Request.FormValue("product_name")
	PRD.ProductPrice, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("product_price"))
	file, err := c.FormFile("product_image")
	if err != nil {
		log.Fatal(err)
	}
	PRD.ProductImage = file.Filename
	PRD.ProductType, _ = commonfunction.ParseStringToInt(c.Request.FormValue("product_type"))
	PRD.ProductWeight = c.Request.FormValue("product_weight")
	PRD.ProductBrand = c.Request.FormValue("product_brand")
	PRD.ProductStatus = c.Request.FormValue("product_status")
	PRD.TaxPercentage, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("tax_percentage"))
	c.SaveUploadedFile(file, "/Users/nayeem/Documents/"+file.Filename)

	return PRD
}

func ParseStringToFloat64(s string) {
	panic("unimplemented")
}

func (PD *Products) ValidateProducts() error {
	if PD.ProductName == "" {
		return errors.New("Product Name is required")
	}
	if PD.ProductPrice == 0 {
		return errors.New("Product Price is required")
	}
	if PD.ProductImage == "" {
		return errors.New("Product Image is required")
	}
	if PD.ProductType == 0 {
		return errors.New("Product Type is required")
	}
	if PD.ProductWeight == "" {
		return errors.New("Product Weight is required")
	}
	if PD.ProductBrand == "" {
		return errors.New("Product Brand is required")
	}
	if PD.TaxPercentage == 0 {
		return errors.New("Tax Percentage is required")
	}
	if PD.ProductStatus == "" {
		return errors.New("Product Status is required")
	}
	return nil
}

func (PRD *Products) SaveProduct(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Create(&PRD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (PRD *Products) UpdateProduct(db *gorm.DB) error{
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&PRD).Updates(&PRD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (PRD *Products) DeleteProduct(db *gorm.DB) error{
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&PRD).Delete(&PRD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}