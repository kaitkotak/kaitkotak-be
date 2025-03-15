package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

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

// func init() {
// 	err := godotenv.Load() // Load .env file
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	log.Println("✅ .env file loaded successfully")
// }

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	JobTitle  string    `json:"job_title"`
	CreatedAt time.Time `json:"created_at"`
}

func NewApp() *fiber.App {
	cfg := config.LoadConfig()

	database.ConnectDB(cfg)
	defer database.CloseDB()

	app := fiber.New()

	// Fix the handler signature
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!") // Ensure this returns an `error`
	// })

	app.Get("/", func(c fiber.Ctx) error {
		users, err := getUsers()
		if err != nil {
			log.Println("Failed to fetch users:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
		}
		return c.JSON(users)
	})

	return app
}

func getUsers() ([]User, error) {
	ctx := context.Background()

	// Ensure the database connection is available
	if database.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	// Execute the query
	rows, err := database.DB.Query(ctx, "SELECT id, name, job_title, created_at FROM users")
	if err != nil {
		log.Println("Database query error:", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.JobTitle, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
