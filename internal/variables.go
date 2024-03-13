package internal

import (
	"database/sql"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type HandlerResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Error represents the data returned by an aborted HTTP request.
type Error struct {
	Time   string `json:"time_RFC3339"`
	Path   string `json:"path"`
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// Credentials represents the data needed to generate a token.
type Credentials struct {
	Id          int    `json:"id"`
	UserName    string `json:"user_name" binding:"required"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password"`
}

// Define a struct that will be encoded to a token
type Claims struct {
	ID          int    `json:"id"`
	UserName    string `json:"user_name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	jwt.StandardClaims
}

// DetailError returns an Error struct with the given status and message.
func DetailError(status int, err error) *Error {
	return &Error{
		Time:   time.Now().Format(time.RFC3339),
		Path:   "",
		Status: status,
		Error:  error.Error(err),
	}
}

const (
	Port           = ":8080"
	TokenIssuer    = "https://.com"
	Issuer2        = "accounts.google.com"
	CAccessToken   = "access-token"
	CAccessClaims  = "access-claims"
	CRefreshToken  = "refresh-token"
	CRefreshClaims = "refresh-claims"
)

// ENV and SECRETS
var (
	_   = godotenv.Load(".env")
	Env = os.Getenv("ENV")

	TokenSignatureKey string = "secret overwritten by init.go"
	ClientKey         string = os.Getenv("ANDROID_KEY")
	SwaggerUrl        string = "https://www.google.com/"

	//Database
	DB     *sql.DB = nil
	DB_URL string  = os.Getenv("DB_URL")
)
