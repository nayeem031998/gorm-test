package model

import (
	"errors"

	"gorm.io/gorm"
)

type InventoryTracker struct {
	InventoryId     int     `gorm:"primary_key;auto_increment" json:"inventory_id"`
	QuantitySold    int64   `json:"quantity_sold"`
	PricePerUnit    float64 `json:"price_per_unit"`
	ProductId       int     `json:"product_id"`
	QuantityInStock int64   `json:"quantity_in_stock"`
}

func (InventoryTracker) TableName() string {
	return "inventory_trackers"
}

func (IT *InventoryTracker) ValidateInventoryTracker() error {
	
	if IT.QuantitySold == 0 {
		return errors.New("quantity Sold is required")
	}
	if IT.PricePerUnit == 0 {
		return errors.New("price Per Unit is required")
	}
	if IT.ProductId == 0 {
		return errors.New("Product Id is required")
	}
	if IT.QuantityInStock == 0 {
		return errors.New("quantity In Stock is required")
	}
	return nil
}
func (IT *InventoryTracker) SaveInventoryTracker(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Create(IT).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (IT *InventoryTracker) UpdateInventoryTrackers(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&InventoryTracker{}).Where("inventory_id = ?", IT.InventoryId).Updates(InventoryTracker{QuantitySold: IT.QuantitySold, PricePerUnit: IT.PricePerUnit, ProductId: IT.ProductId, QuantityInStock: IT.QuantityInStock}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (IT *InventoryTracker) DeleteInventoryTrackers(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&IT).Delete(&IT).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
