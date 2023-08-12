package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string     `json:"title" gorm:"not null; unique"`
	Description string     `json:"description" gorm:"not null;"`
	Year        int        `json:"year" gorm:"not null"`
	UserID      uint       `json:"-"`
	User        CustomUser `gorm:"foreignkey:UserID" json:"user"`
}
