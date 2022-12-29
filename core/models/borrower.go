package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Borrower struct {
	Id          string         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;index"`
	Email       string         `json:"email" gorm:"not null;index"`
	BorrowBooks []*Book        `json:"borrow_books" gorm:"foreignKey:BorrowedBy;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (b *Borrower) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *Borrower) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), b)
}
