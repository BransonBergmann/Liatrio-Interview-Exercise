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

	out := output {
		message: "My name is Branson Bergmann",
		timestamp: time.Now().Unix(),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(out)
	})

	log.Fatal(app.Listen(":8080"))
}