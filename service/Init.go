package service

import (
	"fmt"
	"log"
	"net/http"
	"token-master/internal"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsConfig(router))
	router.Use(checkAPIKey())
	getRoutes(router)

	fmt.Println("Serving", internal.Port)
	return router
}

func corsConfig(*gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTION")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-API-Key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		//fmt.Println(c.Request.Method)
		//fmt.Println(c.Request.Response)
		//fmt.Println(c.Request.WithContext(c))
	}
}

func checkAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if internal.ClientKey != c.GetHeader("X-API-Key") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid api key")
		}

		c.Next()
	}
}

func parseAndSetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(internal.CAccessToken)
		refreshToken := c.GetHeader(internal.CRefreshToken)

		if accessToken != "" {
			token, claims, err := internal.ValidateToken(c, accessToken)
			if err != nil {
				c.AbortWithStatusJSON(err.Status, err)
			}

			if token == nil || claims == nil {
				log.Println("token or claims is nil")
				c.AbortWithStatusJSON(http.StatusInternalServerError, "token or claims is nil")
			}

			c.Set(internal.CAccessToken, token)
			c.Set(internal.CAccessClaims, claims)
		}

		if refreshToken != "" {
			token, claims, err := internal.ValidateToken(c, refreshToken)
			if err != nil {
				c.AbortWithStatusJSON(err.Status, err)
			}

			if token == nil || claims == nil {
				log.Println("token or claims is nil")
				c.AbortWithStatusJSON(http.StatusInternalServerError, "token or claims is nil")
			}
			
			c.Set(internal.CRefreshToken, token)
			c.Set(internal.CRefreshClaims, claims)
		}

		c.Next()
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
func getRoutes(router *gin.Engine) {
	V := router.Group("/")

	Private(V)
}

type ControllerFunctions func(c *gin.Context) (response interface{}, err *internal.Error)

func gHandler(c *gin.Context, fn ControllerFunctions) {
	response, err := fn(c)
	if err != nil {
		path := c.Request.URL.Path
		err.Path = path
		c.AbortWithStatusJSON(err.Status, err)
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
