package routes

import (
  "github.com/JamesAndresCM/golang-fiber-example/app/handlers"
	"github.com/JamesAndresCM/golang-fiber-example/app/middlewares"
  "github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
  api := app.Group("/api")

  v1 := api.Group("/v1")
  v1.Post("/sign-up", handlers.SignUp)
	v1.Post("/sign-in", handlers.SignIn)

	v1.Get("/current_user", middlewares.JWTMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"message": "Ruta protegida",
		})
	})
}
