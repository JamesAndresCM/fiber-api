package models

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Title        string `json:"title" gorm:"not null; unique"`
	Description  string `json:"description" gorm:"not null; unique"`
	Year         int    `json:"year" gorm:"not null"`
}
