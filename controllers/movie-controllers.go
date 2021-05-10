package controllers

import (
  "strconv"
  "github.com/JamesAndresCM/golang-fiber-example/lib"
  "github.com/gofiber/fiber/v2"
  "github.com/JamesAndresCM/golang-fiber-example/models"
)

func ListAllMovies(c *fiber.Ctx) error {
  var movie models.Movie
  movies, _ := movie.GetMovies()
  return c.JSON(movies)
}

func GetMovie(c *fiber.Ctx) error {
  var movie models.Movie
  id, err := strconv.Atoi(c.Params("id"))
  if err != nil {
    return c.JSON(lib.Response(200, err.Error()))
  }
  result, err := movie.GetMovie(id)
  if err != nil{
    return c.JSON(lib.Response(200, err.Error()))
  }
  return c.JSON(result)
}
