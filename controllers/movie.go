package controllers

import (
  "github.com/gofiber/fiber/v2"
)

type Movie struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Year int `json:"year"`
}
