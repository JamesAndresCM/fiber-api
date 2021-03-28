package migration

import (
  "github.com/JamesAndresCM/golang-fiber-example/configuration"
  "github.com/JamesAndresCM/golang-fiber-example/models"
)
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()
  fmt.Println(db)

	db.CreateTable(&models.Movie{})
}
