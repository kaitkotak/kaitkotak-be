package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaitkotak-be/internal/config"
)

var DB *pgxpool.Pool

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	// Print to debug
	log.Println("Generated DSN:", dsn)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to parse DB config: %v", err)
	}

	// Access ConnConfig and set PreferSimpleProtocol
	connConfig := config.ConnConfig
	connConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol // Disable prepared statements
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = pool
	fmt.Println("Connected to Supabase PostgreSQL successfully!")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		DB = nil // Ensure DB is nil after closing
		fmt.Println("Database connection closed.")
	}
}
