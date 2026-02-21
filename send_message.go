package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Output struct {
	Message		string `json:"message"`
	Timestamp 	int64  `json:"timestamp"`
}

func main() {
	app := fiber.New()

	out := Output {
		Message: "My name is Branson Bergmann",
		Timestamp: time.Now().Unix(),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(out)
	})

	log.Fatal(app.Listen(":8080"))
}