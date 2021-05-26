package controllers

import (
	"github.com/JamesAndresCM/golang-fiber-example/lib"
	"github.com/JamesAndresCM/golang-fiber-example/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
	if err != nil {
		return c.JSON(lib.Response(400, err.Error()))
	}
	return c.JSON(result)
}

func CreateMovie(c *fiber.Ctx) error {
	movie := new(models.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	result, err := movie.CreateMovie()
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	return c.JSON(result)
}

func DestroyMovie(c *fiber.Ctx) error {
	var movie models.Movie
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	result, err := movie.Delete(id)
	if err != nil {
		return c.JSON(lib.Response(400, err.Error()))
	}
	return c.JSON(fiber.Map{"status": 200, "message": "Movie id " + strconv.Itoa(result) + " successfully deleted"})
}

func UpdateMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}

	movie := new(models.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}

	result, err := movie.Update(id)
	if err != nil {
		return c.JSON(lib.Response(400, err.Error()))
	}
	return c.JSON(result)
}
