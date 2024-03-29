package main

import (
	"flag"
	"github.com/JamesAndresCM/fiber-api/db"
	"github.com/JamesAndresCM/fiber-api/migration"
	"github.com/JamesAndresCM/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	db.GetConnection()
	RunMigrations()
	app := fiber.New()
	app.Use(logger.New())

	routes.MovieRoutes(app)
	routes.UserRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "First Endpoint",
		})
	})
	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}

func RunMigrations() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generating migrations")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Start migrate process")
		migration.Migrate()
    log.Fatal("Migration finished")
	}
}
