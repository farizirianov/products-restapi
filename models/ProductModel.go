package models

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type Product struct {
	Sku            string    `gorm:"column:sku;unique;primaryKey;not null" json:"sku"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Brand          string    `gorm:"column:brand;not null" json:"brand"`
	Size           string    `gorm:"column:size;not null" json:"size"`
	Price          float64   `gorm:"column:price;not null" json:"price"`
	MainImageUrl   string    `gorm:"column:mainImageUrl;not null" json:"mainImageUrl"`
	OtherImagesUrl SliceUrl  `gorm:"embedded"`
	CreatedAt      time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;autoUpdateTime:milli" json:"updatedAt"`
}

type SliceUrl []string

// Implement the Scan method that receives a value from the database and converts it to a slice of strings
func (su *SliceUrl) Scan(src interface{}) error {
	if b, ok := src.([]byte); ok {
		s := string(b)
		urls := strings.Split(s, ",")
		*su = urls
		return nil
	}
	return errors.New("invalid data type for SliceUrl")
}

// Implement the Value method that receives a slice of strings and converts it into a value that can be stored in the database.
func (su SliceUrl) Value() (driver.Value, error) {
	s := strings.Join(su, ",")
	return s, nil
}

func (Product) TableName() string {
	return "products"
}
