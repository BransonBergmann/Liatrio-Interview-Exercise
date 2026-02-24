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

	
	
	app.Get("/", func(c *fiber.Ctx) error {
		out := Output {
		Message: "My name is Branson Bergmann",
		Timestamp: time.Now().Unix(),
		}
		return c.JSON(out)
	})

	addr := "0.0.0.0:80"
	log.Fatal(app.Listen(addr))
}