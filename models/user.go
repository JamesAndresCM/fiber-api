package models

import (
	"github.com/JamesAndresCM/golang-fiber-example/configuration"
	"github.com/JamesAndresCM/golang-fiber-example/models/scopes"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Movies   []Movie `json:"movies"`
}
