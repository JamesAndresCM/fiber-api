package routes

import (
  "github.com/JamesAndresCM/golang-fiber-example/controllers"
  "github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
  api := app.Group("/api")

  v1 := api.Group("/v1")
  v1.Get("/movies", controllers.ListAllMovies)
  v1.Get("/movies/:id",controllers.GetMovie)
}
