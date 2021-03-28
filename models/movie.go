package models

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Title        string `json:"title" gorm:"not null; unique"`
	Description  string `json:"description" gorm:"not null; unique"`
	Year         int    `json:"year" gorm:"not null"`
}

func GetMovies(db *gorm.DB) {
	var movies []Movie

	if result := db.Find(&movies); result.Error != nil {
		return movies, result.Error
	}
	return movies, nil
}
