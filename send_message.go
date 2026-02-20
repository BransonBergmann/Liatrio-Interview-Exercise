package main

import (
	"log"

	
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/text", func(c *fiber.Ctx) error {
		return c.SendString("My name is Branson Bergmann!")
	})

	log.Fatal(app.Listen(":8080"))
}