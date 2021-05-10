package lib

import (
  "log"
  "github.com/gofiber/fiber/v2"
)

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Response(code int, err string) map[string]interface{} {
  data := make(map[string]interface{})
  data["response"] = fiber.NewError(code, err)
  return data
}
