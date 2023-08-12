package handlers

import (
	"github.com/JamesAndresCM/golang-fiber-example/app/dto"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/JamesAndresCM/golang-fiber-example/app/services"
	"github.com/JamesAndresCM/golang-fiber-example/lib"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

const ObjectsPerPage = 10

func ListAllMovies(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	movies, _ := services.GetMovies(page, pageSize)
	countmovies, _ := services.CountMovies()

	totalPages := math.Ceil(float64(countmovies) / float64(pageSize))
	meta := dto.Meta{CurrentPage: page, TotalElements: int(countmovies), TotalPages: totalPages, ObjectsPerPage: ObjectsPerPage}
	data := dto.Response{Movies: movies, Meta: meta}
	return c.JSON(data)
}

func GetMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	result, err := services.GetMovie(id)
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
	movie.UserID = uint(c.Locals("user_id").(float64))
	result, err := services.CreateMovie(movie)
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	return c.JSON(result)
}

func DestroyMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(lib.Response(200, err.Error()))
	}
	userID := uint(c.Locals("user_id").(float64))
	err = services.DeleteMovie(id, userID)
	if err != nil {
		return c.JSON(lib.Response(400, err.Error()))
	}
	return c.JSON(fiber.Map{"status": 200, "message": "Movie id " + strconv.Itoa(id) + " successfully deleted"})
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
	userID := uint(c.Locals("user_id").(float64))
	result, err := services.UpdateMovie(id, movie, userID)
	if err != nil {
		return c.JSON(lib.Response(400, err.Error()))
	}
	return c.JSON(result)
}
