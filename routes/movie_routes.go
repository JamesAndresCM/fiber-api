package routes

import (
	"github.com/JamesAndresCM/fiber-api/app/handlers"
	"github.com/JamesAndresCM/fiber-api/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/movies", handlers.ListAllMovies)
	v1.Get("/movies/:id", handlers.GetMovie)
	v1.Post("/movies", middlewares.JWTMiddleware(), handlers.CreateMovie)
	v1.Delete("/movies/:id", middlewares.JWTMiddleware(), handlers.DestroyMovie)
	v1.Patch("/movies/:id", middlewares.JWTMiddleware(), handlers.UpdateMovie)
	v1.Put("/movies/:id", middlewares.JWTMiddleware(), handlers.UpdateMovie)
}
