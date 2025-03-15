package cmd

import (
	"github.com/kaitkotak-be/internal/config"
	"github.com/kaitkotak-be/internal/database"

	"github.com/gofiber/fiber/v3"
)

// func main() {
// 	cfg := config.LoadConfig()

// 	database.ConnectDB(cfg)
// 	defer database.CloseDB()

// 	app := fiber.New()

// 	app.Get("/", func(c fiber.Ctx) error {
// 		return c.SendString("Hello, World 👋!")
// 	})

// 	fmt.Println("Server running on port 8000...")
// 	log.Fatal(app.Listen(":8000"))
// }

func NewApp() *fiber.App {
	cfg := config.LoadConfig()

	database.ConnectDB(cfg)
	defer database.CloseDB()

	app := fiber.New()

	// Fix the handler signature
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!") // Ensure this returns an `error`
	})

	return app
}
