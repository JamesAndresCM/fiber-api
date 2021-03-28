package migrations

import (
  "github.com/JamesAndresCM/go-fiber/configuration"
  "github.com/JamesAndresCM/go-fiber/models"
)
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.Movie{})
}
