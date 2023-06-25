package migration

import (
  "github.com/JamesAndresCM/golang-fiber-example/db"
  "github.com/JamesAndresCM/golang-fiber-example/app/models"
)
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.Movie{})
	db.CreateTable(&models.User{})
}
