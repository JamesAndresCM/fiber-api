package models

import (
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/app/models/scopes"
  "gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null; unique"`
	Description string `json:"description" gorm:"not null; unique"`
	Year        int    `json:"year" gorm:"not null"`
	UserID      uint   `json:"-"`
	User        CustomUser   `gorm:"foreignkey:UserID" json:"user"`
}

func (movie *Movie) GetMovies(page, pageSize int) ([]*Movie, error) {
    db := db.GetConnection()
    movies := []*Movie{}
    if result := db.Scopes(scopes.Paginate(page, pageSize)).Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Table("users")
    }).Joins("JOIN users on movies.user_id = users.id").Find(&movies); result.Error != nil {
        return movies, result.Error
    }
    return movies, nil
}

func (movie *Movie) CountMovies() (int64, error) {
	db := db.GetConnection()
	movies := []*Movie{}
	var count int64
	if result := db.Find(&movies).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (movie *Movie) GetMovie(id int) (*Movie, error) {
	db := db.GetConnection()
	if result := db.Find(&movie, id); result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (movie *Movie) CreateMovie() (*Movie, error) {
	db := db.GetConnection()
	if result := db.Create(&movie); result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (movie *Movie) Delete(id int) (int, error) {
	db := db.GetConnection()
	if result := db.Find(&movie, id); result.Error != nil {
		return id, result.Error
	}
	db.Delete(&movie)
	return id, nil
}

func (movie *Movie) Update(id int) (*Movie, error) {
	db := db.GetConnection()
	mov := Movie{Title: movie.Title, Description: movie.Description, Year: movie.Year}
	if result := db.First(&movie, id).UpdateColumns(mov); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}
