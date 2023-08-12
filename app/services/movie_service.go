package services

import (
	"errors"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/app/models/scopes"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"gorm.io/gorm"
)

func GetMovies(page, pageSize int) ([]*models.Movie, error) {
	db := db.DB.Db
	movies := []*models.Movie{}
	if result := db.Scopes(scopes.Paginate(page, pageSize)).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}).Joins("JOIN users on movies.user_id = users.id").Find(&movies); result.Error != nil {
		return movies, result.Error
	}
	return movies, nil
}

func CountMovies() (int64, error) {
	db := db.DB.Db
	movies := []*models.Movie{}
	var count int64
	if result := db.Find(&movies).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func GetMovie(id int) (*models.Movie, error) {
	db := db.DB.Db
	movie := &models.Movie{}
	if result := db.First(movie, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Movie not found")
		}
		return nil, result.Error
	}
	return movie, nil
}

func CreateMovie(movie *models.Movie) (*models.Movie, error) {
	db := db.DB.Db
	if result := db.Create(&movie); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func DeleteMovie(id int, userID uint) error {
	db := db.DB.Db
	movie := &models.Movie{}
	if result := db.First(movie, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("Movie not found")
		}
		return result.Error
	}
	if movie.UserID != userID {
		return errors.New("unauthorized: the movie does not belong to the user")
	}

	if result := db.Delete(movie); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateMovie(id int, updatedMovie *models.Movie, userID uint) (*models.Movie, error) {
	db := db.DB.Db
	movie := &models.Movie{}
	if result := db.First(movie, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Movie not found")
		}
		return nil, result.Error
	}

	if movie.UserID != userID {
		return nil, errors.New("unauthorized: the movie does not belong to the user")
	}

	mov := models.Movie{Title: updatedMovie.Title, Description: updatedMovie.Description, Year: updatedMovie.Year}
	if result := db.Model(movie).UpdateColumns(mov); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}
