package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BookTitle struct {
	Id        string         `json:"id" gorm:"primaryKey;index"`
	Title     string         `json:"title" gorm:"not null;index"`
	Books     []*Book        `json:"books" gorm:"foreignKey:TitleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (b *BookTitle) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *BookTitle) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), b)
}
