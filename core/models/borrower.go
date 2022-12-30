package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Borrower struct {
	Id            string         `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"not null;unique"`
	Email         string         `json:"email" gorm:"not null;unique"`
	Password      string         `json:"password" gorm:"not null"`
	BorrowedBooks []*Book        `json:"borrowed_books" gorm:"foreignKey:BorrowedBy;references:Id"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}

func (b *Borrower) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *Borrower) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), b)
}
