package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IntIdtoUint(num int) (uint, *Error) {
	if num >= 0 {
		return uint(num), nil
	}

	return 0, DetailError(400, fmt.Errorf("id conversion to unsigned integer, negative id"))
}

func ValidateToken(c *gin.Context, tokenString string) (*jwt.Token, *Claims, *Error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ValidateToken unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TokenSignatureKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, nil, DetailError(http.StatusBadRequest, fmt.Errorf("ValidateToken invalid token signature: "+err.Error()))
		}

		return nil, nil, DetailError(http.StatusInternalServerError, fmt.Errorf("ValidateToken could not parse token: "+err.Error()))
	}
	if !token.Valid {
		return nil, nil, DetailError(http.StatusBadRequest, fmt.Errorf("ValidateToken invalid token"))
	}

	return token, claims, nil
}

func GetAccessTokenFromCtx(c *gin.Context) (*jwt.Token, *Claims) {
	token, _ := c.Get(CAccessToken)
	claims, _ := c.Get(CAccessClaims)

	return token.(*jwt.Token), claims.(*Claims)
}

func GetRefreshTokenFromCtx(c *gin.Context) (*jwt.Token, *Claims) {
	token, _ := c.Get(CRefreshToken)
	claims, _ := c.Get(CRefreshClaims)

	return token.(*jwt.Token), claims.(*Claims)
}
