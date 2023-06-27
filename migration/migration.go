package migration

import (
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/lib"
)

func Migrate() {
	db := db.GetConnection()
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.User{})

	pgDB, err := db.DB()
	lib.Fatal(err)
	defer pgDB.Close()
}
