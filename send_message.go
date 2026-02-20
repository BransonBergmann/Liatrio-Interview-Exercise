package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Output struct {
	message		string
	timestamp 	int64
}

func main() {
	app := fiber.New()

	out := Output {
		"My name is Branson Bergmann",
		time.Now().Unix(),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(out)
	})

	log.Fatal(app.Listen(":8080"))
}