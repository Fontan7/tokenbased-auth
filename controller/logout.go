package controller

import (
	"token-master/database"
	"token-master/internal"

	"github.com/gin-gonic/gin"
)

func UserLogout(c *gin.Context) (response interface{}, err *internal.Error) {
	// Get the refresh token from the request header
	refreshToken := c.GetHeader("refresh-token")
	_, claims := internal.GetRefreshTokenFromCtx(c)

	//we replace the token for an invalid string if its equal on db, in this case empty
	u := database.User{}
	u.GetUserByEmail(claims.Email)
	if err := u.UpdateUserRefreshToken(refreshToken, ""); err != nil {
		return nil, err
	}

	response = "accepted"
	return response, nil
}
