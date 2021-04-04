package routes

import (
  "github.com/JamesAndresCM/golang-fiber-example/controllers"
  "github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
  app.Get("api/v1/movies", controllers.ListAllMovies)
}
