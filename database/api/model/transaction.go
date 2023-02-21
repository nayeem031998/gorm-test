package model

import (
	"database/api/utils/commonfunction"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, _ := time.ParseInLocation("2006-01-02", strInput, time.Local)
	t.Time = newTime
	return nil
}

type Transaction struct {
	TransactionId     int64     `gorm:"primary_key;auto_increment" json:"transaction_id"`
	CustomerId        int       `json:"customer_id"`
	TransactionAmount float64   `json:"transaction_amount"`
	TxnDate           Datetime  `gorm:"-" json:"txn_date"`
	TransactionDate   time.Time `json:"transaction_date"`
	MobileNumber      int64     `json:"mobile_number"`
	ProductId         int       `json:"product_id"`
	DiscountAmount    float64   `json:"discount_amount"`
	TotalAmount       float64   `json:"total_amount"`
}

func (TRN *Transaction) TableName() string {
	return "transactions"
}
func (TR *Transaction) TransactionStuct(c *gin.Context) *Transaction {
	TR.TransactionId, _ = commonfunction.ParseStringToInt64(c.Request.FormValue("transaction_id"))
	TR.CustomerId, _ = commonfunction.ParseStringToInt(c.Request.FormValue("customer_id"))
	TR.TransactionAmount, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("transaction_amount"))
	TR.MobileNumber, _ = commonfunction.ParseStringToInt64(c.Request.FormValue("mobile_number"))
	TR.ProductId, _ = commonfunction.ParseStringToInt(c.Request.FormValue("product_id"))
	TR.DiscountAmount, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("discount_amount"))
	TR.TotalAmount, _ = commonfunction.ParseStringToFloat64(c.Request.FormValue("total_amount"))
	return TR
}
func (TR *Transaction) ValidateTransaction() error {
	TR.TransactionDate = TR.TxnDate.Time
	if TR.CustomerId == 0 {
		return errors.New("customer Id is required")
	}
	if TR.TransactionAmount == 0 {
		return errors.New("transaction Amount is required")
	}
	if TR.TransactionDate.IsZero() {
		return errors.New("transaction Date is required")
	}
	if TR.MobileNumber == 0 {
		return errors.New("mobile Number is required")
	}
	if TR.ProductId == 0 {
		return errors.New("product Id is required")
	}
	if TR.DiscountAmount == 0 {
		return errors.New("discount Amount is required")
	}
	if TR.TotalAmount == 0 {
		return errors.New("total Amount is required")
	}
	return nil
}

func (TRN *Transaction) SaveTransaction(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Create(&TRN).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (TRN *Transaction) UpdateTransaction(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Model(&TRN).Updates(&TRN).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (TRN *Transaction) DeleteTransaction(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Debug().Delete(&TRN).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
