package migration

import (
	"github.com/JamesAndresCM/fiber-api/app/models"
	"github.com/JamesAndresCM/fiber-api/db"
	"github.com/JamesAndresCM/fiber-api/lib"
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
