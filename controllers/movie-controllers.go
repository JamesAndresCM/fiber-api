package controllers

import "github.com/gofiber/fiber/v2"

func ListAllMovies(c *fiber.Ctx) {
  c.Send("All movies")
}
