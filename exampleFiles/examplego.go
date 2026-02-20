package test

import (
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func test() {
    e := echo.New()

    // Middleware
    e.Use(middleware.RequestLogger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.GET("/health", func(c echo.Context) error {
        return c.String(http.StatusOK, "Health is OK!!")
    })

    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "8080"
    }

    e.Logger.Fatal(e.Start(":" + httpPort))
}