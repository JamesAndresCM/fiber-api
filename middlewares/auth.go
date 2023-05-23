package middlewares

import (
	"github.com/JamesAndresCM/golang-fiber-example/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		bearerPrefix := "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido 1",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		// Verificar y analizar el token JWT
		token, err := utils.ParseJWT(tokenString)
		if err != nil {
			// El token es inválido, devolver una respuesta de error
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		// Extraer los datos del token (en este caso, el user id)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			// No se pudieron extraer los datos del token, devolver una respuesta de error
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		id := claims["user_id"]

		// Establecer el ID en el contexto
		c.Locals("user_id", id)

		// Continuar con la siguiente función de middleware o controlador
		return c.Next()
	}
}
