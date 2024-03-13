package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"tokenbased-auth/database"
	i "tokenbased-auth/internal"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

func RefreshToken(c *gin.Context) (interface{}, *i.Error) {
	// Get the refresh token from the request header
	refreshToken := c.GetHeader("refresh-token")
	_, claims := i.GetRefreshTokenFromCtx(c)
	if claims == nil {
		return nil, i.DetailError(http.StatusInternalServerError, errors.New("invalid token"))
	}

	if claims.StandardClaims.Subject != "refresh" {
		return nil, i.DetailError(http.StatusBadRequest, errors.New("invalid token type"))
	}

	//go get this in the db
	user := database.User{}
	err := user.GetUserByEmail(claims.Email)
	if err != nil {
		err.Error = fmt.Sprintf("RefreshToken: %v", err)
		return nil, err
	}

	currentRefreshToken, err := user.GetUserRefreshToken(user.ID)
	if err != nil {
		err.Error = fmt.Sprintf("RefreshToken: %v", err)
		return nil, err
	}

	if currentRefreshToken != refreshToken {
		return nil, i.DetailError(http.StatusUnauthorized, errors.New("Refresh token invalid token"))
	}

	// Generate a new access token
	newAccessToken, err := GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err = GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	tokens := map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": refreshToken,
	}

	return tokens, nil
}

func GenerateRefreshToken(u database.User) (string, *i.Error) {
	// Declare the expiration time of the token
	//refresh token should be long lived
	expirationTime := time.Now().Add(152 * time.Hour)

	// Create the JWT claims
	claims := &i.Claims{
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    i.TokenIssuer,
			Subject:   "refresh",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	newRefreshToken, err := token.SignedString([]byte(i.TokenSignatureKey))
	if err != nil {
		return "", i.DetailError(500, err)
	}

	t := database.Token{
		Raw: newRefreshToken,
	}

	updateErr := t.InsertRefreshToken(u.ID)
	if updateErr != nil {
		return "", updateErr
	}

	return newRefreshToken, nil
}
