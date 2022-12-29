package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Book struct {
	Id         string    `json:"id" gorm:"primaryKey;index"`
	TitleId    string    `json:"title_id" gorm:"not null;index"`
	Authors    []*Author `json:"authors" gorm:"many2many:book_authors;"`
	BorrowedBy string    `json:"borrowed_by" gorm:"index"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"index"`
}

func (b *Book) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *Book) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), b)
}
