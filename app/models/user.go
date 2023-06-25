package models

import (
	"errors"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/app/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Movies   []Movie `json:"movies"`
}

func (u *User) Register() (string, error) {
	db := db.GetConnection()
	// Verificar si el nombre de usuario ya está en uso
	if !db.Where("name = ?", u.Name).First(&User{}).RecordNotFound() {
		return "", errors.New("El nombre de usuario ya está en uso")
	}

	if !db.Where("email = ?", u.Email).First(&User{}).RecordNotFound() {
		return "", errors.New("El email ya está en uso")
	}

	// Hash de la contraseña antes de guardarla en la base de datos
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	u.Password = string(hashedPassword)

	// Crear el usuario en la base de datos
	if err := db.Create(u).Error; err != nil {
		return "", err
	}

	// Generar el token JWT
	tokenString, err := utils.GenerateJWT(u.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (user *User) Authenticate(email, password string) (string, error) {
	db := db.GetConnection()
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
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
