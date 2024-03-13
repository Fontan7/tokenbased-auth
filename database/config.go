package database

import (
	"database/sql"
	"fmt"
	"log"
	"tokenbased-auth/internal"

	_ "github.com/lib/pq"
)

// Config is the database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func DatabaseConnect() (*sql.DB, error) {
	// Define the connection string
	connStr := internal.DB_URL
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
	return db, nil
}

func GetTokenSignatureKey() (string, error) {
	query := GetTokenSignatureKeyKeyQuery
	row := internal.DB.QueryRow(query)

	var signature string
	err := row.Scan(&signature)
	if err != nil {
		return "", fmt.Errorf("GetTokenSigningKey: %v", err)
	}

	return signature, nil
}

// Define your queries here
const (
	// User queries
	GetUserByIDQuery             = "SELECT * FROM users WHERE id = $1"
	GetTokenSignatureKeyKeyQuery = "SELECT s.signature_string FROM tokenbased-auth.signature s  WHERE id = (SELECT MAX(id) FROM tokenbased-auth.signature)"
	GetUserByEmailQuery          = "SELECT id, user_name, display_name, email, spotify_id, apple_id FROM users.user WHERE email = $1"
	GetUserByIdQuery             = "SELECT id, user_name, display_name, email, spotify_id, apple_id FROM users.user WHERE id = $1"
	InsertUserQuery              = "INSERT INTO users (user_name, email) VALUES ($1, $2) RETURNING id"

	// Token queries
	GetUserRefreshTokenQuery      = "SELECT refresh_token FROM tokenbased-auth.token WHERE user_id = $1"
	GetEqualUserRefreshTokenQuery = "SELECT refresh_token FROM tokenbased-auth.token WHERE user_id = $1 AND refresh_token = $2"
	UpdateUserRefreshTokenQuery   = "UPDATE tokenbased-auth.token SET refresh_token = $1, update = NOW() WHERE user_id = $2"
	InsertRefreshTokenQuery       = "INSERT INTO tokenbased-auth.token (user_id, refresh_token) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET refresh_token = EXCLUDED.refresh_token"
)
