package services

import (
	"errors"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/app/utils"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(user *models.User) (string, error) {
	db := db.GetConnection()

	var existingUser models.User

	if err := db.Where("name = ?", user.Name).First(&existingUser).Error; err == nil {
		return "", errors.New("name is already in use")
	} else if err != gorm.ErrRecordNotFound {
		return "", err
	}

	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return "", errors.New("email is already in use")
	} else if err != gorm.ErrRecordNotFound {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	if err := db.Create(user).Error; err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthenticateUser(email, password string) (string, error) {
	db := db.GetConnection()
	var user models.User

	var err error
	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("User not found")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Not valid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("Error to create JWT")
	}

	return token, nil
}
