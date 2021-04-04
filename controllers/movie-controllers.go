package controllers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/JamesAndresCM/golang-fiber-example/models"
)

func ListAllMovies(c *fiber.Ctx) error {
  var movie models.Movie
  movies, _ := movie.GetMovies()
  return c.JSON(movies)
}
