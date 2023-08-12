package services

import (
	"errors"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/app/utils"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/app/models/scopes"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetMovies(page, pageSize int) ([]*models.Movie, error) {
    movies := []*models.Movie{}
    if result := s.db.Scopes(scopes.Paginate(page, pageSize)).Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Table("users")
    }).Joins("JOIN users on movies.user_id = users.id").Find(&movies); result.Error != nil {
        return movies, result.Error
    }
    return movies, nil
}

func (s *MovieService) CountMovies() (int64, error) {
	movies := []*models.Movie{}
	var count int64
	if result := s.db.Find(&movies).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (s *MovieService) GetMovie(id int) (*models.Movie, error) {
    movie := &models.Movie{}
	if result := s.db.Find(movie, id); result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (s *MovieService) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	if result := s.db.Create(&movie); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func (s *MovieService) DeleteMovie(id int) error {
    movie := &models.Movie{}
	if result := s.db.Find(movie, id); result.Error != nil {
		return result.Error
	}
	if result := s.db.Delete(movie); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MovieService) UpdateMovie(id int, updatedMovie *models.Movie) (*models.Movie, error) {
    movie := &models.Movie{}
	if result := s.db.First(movie, id); result.Error != nil {
		return nil, result.Error
	}
	mov := models.Movie{Title: updatedMovie.Title, Description: updatedMovie.Description, Year: updatedMovie.Year}
	if result := s.db.Model(movie).UpdateColumns(mov); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}