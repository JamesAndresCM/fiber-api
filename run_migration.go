package main

import (
  "log"
  "flag"
  "github.com/JamesACM/golang-fiber-example/migrations"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generating migrations")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Start migrate process")
		migrations.Migrate()
		log.Println("migration finished")
	}
}
