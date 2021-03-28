package migrations

import (
  "github.com/JamesAndresCM/golang-fiber-example/configuration"
  "github.com/JamesAndresCM/golang-fiber-example/models"
)
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.Movie{})
}
