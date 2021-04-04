package models

import (
  "github.com/jinzhu/gorm"
  "github.com/JamesAndresCM/golang-fiber-example/configuration"
)

type Movie struct {
	gorm.Model
	Title        string `json:"title" gorm:"not null; unique"`
	Description  string `json:"description" gorm:"not null; unique"`
	Year         int    `json:"year" gorm:"not null"`
}


func GetMovies() ([]*Movie, error) {
  db := configuration.GetConnection()

  movies := []*Movie{}

	if result := db.Find(&movies); result.Error != nil {
		return movies, result.Error
	}
	return movies, nil
}

