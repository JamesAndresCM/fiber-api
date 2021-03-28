package main

import (
  "log"
  "flag"
  "github.com/JamesAndresCM/golang-fiber-example/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generating migrations")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Start migrate process")
		migration.Migrate()
		log.Println("migration finished")
	}
}
