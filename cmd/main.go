package main

import (
	"fmt"
	"log"

	"github.com/kaitkotak-be/internal/config"
	"github.com/kaitkotak-be/internal/database"
	"github.com/kaitkotak-be/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func init() {
	config.LoadConfig()
}

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.SetupRoutes(app)
	fmt.Println("Server running on port 8000...")
	log.Fatal(app.Listen(":80"))
}
