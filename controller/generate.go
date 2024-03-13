package controller

import (
	"time"
	"tokenbased-auth/database"
	i "tokenbased-auth/internal"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateAccessToken(user database.User) (string, *i.Error) {
	// Declare the expiration time of the token
	// access token should be short-lived
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	claims := &i.Claims{
		ID:          user.ID,
		UserName:    user.UserName,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    i.TokenIssuer,
			Subject:   "access",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(i.TokenSignatureKey))
	if err != nil {
		return "", i.DetailError(500, err)
	}

	return tokenString, nil
}
