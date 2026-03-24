package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"os"
)

var (
	ImageTag  = "unknown"
	GitCommit = "unknown"
)

type Output struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	ImageTag  string `json:"image_tag"`
	GitCommit string `json:"git_commit"`
	Revision  string `json:"revision"`  
}

func main() {
	app := fiber.New()
	revision := os.Getenv("CONTAINER_APP_REVISION")
	if revision == "" {
		revision = "unknown"
	}
	
	
	app.Get("/", func(c *fiber.Ctx) error {
		out := Output{
			Message:   "My name is Branson Bergmann, and I got the job",
			Timestamp: time.Now().UnixMilli(),
			ImageTag:  ImageTag,
			GitCommit: GitCommit,
			Revision:  revision,
		}
		return c.JSON(out)
	})

	addr := "0.0.0.0:80"
	log.Fatal(app.Listen(addr))
}