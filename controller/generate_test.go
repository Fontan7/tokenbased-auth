package controller_test

import (
	"testing"
	"tokenbased-auth/controller"
	"tokenbased-auth/database"
	"tokenbased-auth/internal"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	// Mock the credentials
	user := database.User{
		ID:          1,
		UserName:    "user.test",
		DisplayName: "Test User",
		Email:       "test@example.com",
	}

	// Generate the access token
	tokenString, err := controller.GenerateAccessToken(user)
	assert.Equal(t, nil, err)

	// Parse the token
	token, e := jwt.ParseWithClaims(tokenString, &internal.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(internal.TokenSignatureKey), nil
	})
	assert.NoError(t, e)

	// Assert the token is valid
	assert.True(t, token.Valid)

	// Assert the claims are correct
	claims, ok := token.Claims.(*internal.Claims)
	assert.True(t, ok)
	assert.Equal(t, user.ID, claims.ID)
	assert.Equal(t, user.UserName, claims.UserName)
	assert.Equal(t, user.DisplayName, claims.DisplayName)
	assert.Equal(t, user.Email, claims.Email)
}
