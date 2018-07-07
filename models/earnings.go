package models

import "github.com/jinzhu/gorm"

type Earnings struct {
	gorm.Model
	IDI      int `gorm:"size:100" json:"idi"`
	Download int `gorm:"size:100" json:"download"`
	Category string `gorm:"size:100" json:"category"`
	Country  string `gorm:"size:100" json:"country"`
	City     string `gorm:"size:100" json:"city"`
}
