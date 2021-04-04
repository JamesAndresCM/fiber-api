package controllers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/JamesAndresCM/golang-fiber-example/models"
)

func ListAllMovies(c *fiber.Ctx) error {
  movies, _ := models.GetMovies()
  return c.JSON(movies)
}
