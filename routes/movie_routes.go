package routes

import (
  "github.com/JamesAndresCM/golang-fiber-example/app/handlers"
  "github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
  api := app.Group("/api")

  v1 := api.Group("/v1")
  v1.Get("/movies", handlers.ListAllMovies)
  v1.Get("/movies/:id",handlers.GetMovie)
  v1.Post("/movies", handlers.CreateMovie)
  v1.Delete("/movies/:id", handlers.DestroyMovie)
  v1.Patch("/movies/:id", handlers.UpdateMovie)
}
