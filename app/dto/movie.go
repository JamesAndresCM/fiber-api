package dto

import (
	"github.com/JamesAndresCM/fiber-api/app/models"
)

type Response struct {
	Movies []*models.Movie
	Meta   Meta
}
