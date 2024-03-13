package controller

import (
	"errors"
	"net/http"
	"token-master/database"
	"token-master/internal"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) (interface{}, *internal.Error) {
	var creds internal.Credentials
	// Decode the JSON request to the credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		return nil, internal.DetailError(http.StatusBadRequest, err)
	}

	// In a real application, you'd validate the credentials here
	// If they're valid, continue, if they're not, stop the request
	if creds.Email != "bart@bart.com" || creds.SpotifyId != "bart777" {
		return nil, internal.DetailError(http.StatusUnauthorized, errors.New("invalid credentials"))
	}

	//go get user from db
	user := database.User{}
	err := user.GetUserByEmail(creds.Email)
	if err != nil {
		return nil, err
	}

	//is user logged in already
	ok, err := user.IsLoggedIn(user.ID)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, internal.DetailError(http.StatusUnauthorized, errors.New("user already logged in"))
	}

	// Generate an access token and a refresh token
	accessToken, err := GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	// Finally, send the token to the client
	tokenMap := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	return tokenMap, nil
}
