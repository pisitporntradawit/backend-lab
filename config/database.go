package config

import (
    "context"
    "fmt"
    "os"
    "time"
	"log"
    "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Database() (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
    cfg := os.Getenv("DATABASE_URL")
    if cfg == "" {
        return nil, fmt.Errorf("DATABASE_URL is not set")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    conn, err := pgxpool.New(ctx, cfg)
    if err != nil {
        return nil, fmt.Errorf("connectdb.ConnectDB new pool: %w", err)
    }

    if err := conn.Ping(ctx); err != nil {
        conn.Close()
        return nil, fmt.Errorf("connectdb.ConnectDB ping: %w", err)
    }

    log.Println("database connection established")
    return conn, nil
}  