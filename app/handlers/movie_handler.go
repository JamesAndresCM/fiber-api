package handlers

import (
	"github.com/JamesAndresCM/golang-fiber-example/lib"
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

const ObjectsPerPage = 10

type Meta struct {
	//TODO:include next page, current_page, etc
	CurrentPage    int     `json:"current_page"`
	TotalElements  int     `json:"total_elements"`
	TotalPages     float64 `json:"total_pages"`
	ObjectsPerPage int     `json:"objects_per_page"`
}

type Data struct {
	Movies []*models.Movie
	Meta   Meta
}

func ListAllMovies(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	var movie models.Movie
	movies, _ := movie.GetMovies(page, pageSize)
	countmovies, _ := movie.CountMovies()
	totalPages := math.Ceil(float64(countmovies) / float64(pageSize))
	meta := Meta{CurrentPage: page, TotalElements: countmovies, TotalPages: totalPages, ObjectsPerPage: ObjectsPerPage}
	data := Data{Movies: movies, Meta: meta}
	return c.JSON(data)
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
