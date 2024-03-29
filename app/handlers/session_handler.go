package handlers

import (
	"github.com/JamesAndresCM/fiber-api/app/models"
  "github.com/JamesAndresCM/fiber-api/app/services"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	// Obtener datos de registro del cuerpo de la solicitud
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al analizar los datos de registro"})
	}

	tokenString, err := services.RegisterUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func SignIn(c *fiber.Ctx) error {
	// Obtener los datos de autenticación del cuerpo de la solicitud
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error al analizar los datos de autenticación",
		})
	}

	// Autenticar al usuario y generar el token JWT
	tokenString, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Devolver el token JWT al cliente
	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
