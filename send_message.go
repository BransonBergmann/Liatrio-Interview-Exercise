package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type output struct {
	message		string
	timestamp 	int64
}

func main() {
	app := fiber.New()

	output := output {
		message: "My name is Branson Bergmann",
		timestamp: time.Now().Unix(),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(output)
	})

	log.Fatal(app.Listen(":8080"))
}