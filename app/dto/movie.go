package dto

import (
	"github.com/JamesAndresCM/golang-fiber-example/app/models"
)

type Response struct {
	Movies []*models.Movie
	Meta   Meta
}
