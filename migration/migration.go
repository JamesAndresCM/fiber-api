package migration

import (
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/db"
	"github.com/JamesAndresCM/golang-fiber-example/lib"
)

func Migrate() {
  db := db.DB.Db
	db.AutoMigrate(&models.Movie{})
  db.Migrator().DropTable(&models.CustomUser{})
	db.AutoMigrate(&models.User{})

	pgDB, err := db.DB()
	lib.Fatal(err)
	defer pgDB.Close()
}
