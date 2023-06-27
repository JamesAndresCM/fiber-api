package db

import (
	"fmt"
	"log"
	"os"

	"github.com/JamesAndresCM/golang-fiber-example/lib"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
	"github.com/subosito/gotenv"
  "gorm.io/gorm/logger"
)

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getConfiguration() database {
	var db database
	file, err := os.Open("./.env")
	lib.Fatal(err)
	defer file.Close()

	gotenv.Load()
	db.Database = os.Getenv("DB_NAME")
	db.Password = os.Getenv("DB_PASS")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")

	if db.Database == "" || db.Password == "" || db.Host == "" || db.Port == "" || db.User == "" {
		log.Fatal("env vars not defined")
	}

	return db
}

func GetConnection() *gorm.DB {
    c := getConfiguration()
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
        c.Host, c.Port, c.User, c.Database, c.Password)
    
    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
      Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        panic(err)
    }

    return db
}
