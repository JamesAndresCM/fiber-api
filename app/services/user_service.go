package services

import (
	"errors"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/app/utils"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) (string, error) {
	db := db.GetConnection()
	// Verificar si el nombre de usuario ya está en uso
	if !db.Where("name = ?", user.Name).First(&models.User{}).RecordNotFound() {
		return "", errors.New("El nombre de usuario ya está en uso")
	}

	if !db.Where("email = ?", user.Email).First(&models.User{}).RecordNotFound() {
		return "", errors.New("El email ya está en uso")
	}

	// Hash de la contraseña antes de guardarla en la base de datos
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	// Crear el usuario en la base de datos
	if err := db.Create(user).Error; err != nil {
		return "", err
	}

	// Generar el token JWT
	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthenticateUser(email, password string) (string, error) {
	db := db.GetConnection()
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("Usuario no encontrado")
		}
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Credenciales inválidas")
	}

	// Generar el token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("Error al generar el token JWT")
	}

	return token, nil
}
