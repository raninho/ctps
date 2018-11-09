package ctps

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CTPSBlock struct {
	ID        uint      `gorm:"primary_key"`
	CTPSId    string    `gorm:"column:ctpsid"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (CTPSBlock) TableName() string {
	return "ctps"
}

//InsertCTPS ...
func InsertCTPS(db *gorm.DB, ctps *CTPSBlock) (err error) {
	ctps.CreatedAt = time.Now()

	err = db.Create(ctps).Error
	if err != nil {
		return
	}

	return
}

type CTPSTransaction struct {
	ID            uint      `gorm:"primary_key"`
	CTPSId        string    `gorm:"column:ctpsid"`
	TransactionId string `gorm:"column:transactionid"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

func (CTPSTransaction) TableName() string {
	return "ctps_transaction"
}

//InsertCTPSTransaction ...
func InsertCTPSTransaction(db *gorm.DB, transaction *CTPSTransaction) (err error) {
	transaction.CreatedAt = time.Now()

	err = db.Create(transaction).Error
	if err != nil {
		return
	}

	return
}

//GetLastCTPSTransactionByCTPSID ...
func GetLastCTPSTransactionByCTPSID(db *gorm.DB, ctpsId string) (transaction *CTPSTransaction, err error) {
	transaction = &CTPSTransaction{}
	if err = db.First(transaction, "ctpsid = ?", ctpsId).Error; err != nil {
		return
	}
	return
}
