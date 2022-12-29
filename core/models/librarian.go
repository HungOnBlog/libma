package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Librarian struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;index"`
	Email     string         `json:"email" gorm:"not null;index"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (l *Librarian) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Librarian) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), l)
}
