package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null,index"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Books     []*Book        `json:"books" gorm:"many2many:book_authors;"`
}

func (a *Author) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Author) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), a)
}
