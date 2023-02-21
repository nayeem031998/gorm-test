package model

import (
	//"errors"
	"database/api/utils/commonfunction"
	"errors"

	//"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Orders struct {
	OrderId       uint      ` gorm:"primary_key;auto_increment" json:"order_id"`
	ProductId     int       `json:"product_id"`
	CustomerId    int       `json:"customer_id"`
	OrderDate     time.Time `json:"order_date"`
	RequiredDate  time.Time `json:"required_date"`
	ShippedDate   time.Time `json:"shipped_date"`
	Address       string    `json:"address"`
	OrderStatus   string    `json:"order_status"`
	OrderTotal    float64   `json:"order_total"`
	Quantity      int       `json:"quantity"`
	DeviceDetails string    `json:"device_details"`
}

func (OD *Orders) TableName() string {
	return "orders"
}
func (ORD *Orders) OrdersStuct(c *gin.Context) *Orders {
	ORD.ProductId, _ = commonfunction.ParseStringToInt(c.Request.FormValue("product_id"))
	ORD.CustomerId, _ = commonfunction.ParseStringToInt(c.Request.FormValue("customer_id"))
	ORD.OrderDate, _ = commonfunction.ParseStringToTime(c.Request.FormValue("order_date"), "2006-01-02")
	ORD.RequiredDate, _ = commonfunction.ParseStringToTime(c.Request.FormValue("required_date"), "2006-01-02")
	ORD.ShippedDate, _ = commonfunction.ParseStringToTime(c.Request.FormValue("shipped_date"), "2006-01-02")
	ORD.Address = c.Request.FormValue("address")
	ORD.OrderStatus = c.Request.FormValue("order_status")
	ORD.OrderTotal, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("order_total"))
	ORD.Quantity, _ = commonfunction.ParseStringToInt(c.Request.FormValue("quantity"))
	ORD.DeviceDetails = c.Request.FormValue("device_details")
	return ORD
}

func (OD *Orders) ValidateOrders() error {
	if OD.ProductId == 0 {
		return errors.New("Product Id is required")
	}
	if OD.CustomerId == 0 {
		return errors.New("customer Id is required")
	}
	if OD.OrderDate.IsZero() {
		return errors.New("Order Date is required")
	}
	if OD.RequiredDate.IsZero() {
		return errors.New("required Date is required")
	}
	if OD.ShippedDate.IsZero() {
		return errors.New("shipped Date is required")
	}
	if OD.Address == "" {
		return errors.New("address is required")
	}
	if OD.OrderStatus == "" {
		return errors.New("Order Status is required")
	}
	if OD.OrderTotal == 0 {
		return errors.New("Order Total is required")
	}
	if OD.Quantity == 0 {
		return errors.New("quantity is required")
	}
	if OD.DeviceDetails == "" {
		return errors.New("device Details is required")
	}
	return nil
}

func (ORD *Orders) SaveOrder(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Create(&ORD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (ORD *Orders) UpdateOrder(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&ORD).Where("order_id = ?", ORD.OrderId).Updates(&ORD).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit().Error
}
func (ORD *Orders) DeleteOrder(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&ORD).Delete(&ORD).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (ORD *Orders) FindAllOrders(db *gorm.DB) (*[]Orders, error) {
	var err error
	orders := []Orders{}
	err = db.Debug().Model(&Orders{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Orders{}, err
	}
	return &orders, err
}
func (ORD *Orders) FindOrderById(db *gorm.DB, oid uint) (*Orders, error) {
	var err error
	err = db.Debug().Model(&Orders{}).Where("order_id = ?", oid).Take(&ORD).Error
	if err != nil {
		return &Orders{}, err
	}
	return ORD, err
}

