package service

import (
	ctrl "token-master/controller"

	"github.com/gin-gonic/gin"
)

func Private(rg *gin.RouterGroup) {
	r := rg.Group("/private")

	r.POST("/token-master/login", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.UserLogin) })
	r.POST("/token-master/logout", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.UserLogout) })
	r.POST("/token-master/refresh", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.RefreshToken) })

	//TODO: Add the rest of the routes andf implementations

	//r.POST("/token-master/google", func(c *gin.Context) { gHandler(c, ctrl.GoogleLogin) })
	//r.POST("/token-master/spotify", func(c *gin.Context) { gHandler(c, ctrl.SpotifyLogin) })
	//r.POST("/token-master/apple", func(c *gin.Context) { gHandler(c, ctrl.AppleLogin) })

	//r.POST("/token-master/verify-email", func(c *gin.Context) { gHandler(c, ctrl.VerifyEmail) })
	//r.POST("/token-master/verify-phone", func(c *gin.Context) { gHandler(c, ctrl.VerifyPhone) })
	//r.POST("/token-master/change-email", func(c *gin.Context) { gHandler(c, ctrl.ChangeEmail) })
	//r.POST("/token-master/change-phone", func(c *gin.Context) { gHandler(c, ctrl.ChangePhone) })

	//r.POST("/token-master/register", func(c *gin.Context) { gHandler(c, ctrl.Register) })
	//r.POST("/token-master/confirm-registration", func(c *gin.Context) { gHandler(c, ctrl.ConfirmRegistration) })

	//r.POST("/token-master/password-reset/request", func(c *gin.Context) { gHandler(c, ctrl.RequestPasswordReset) })
	//r.POST("/token-master/password-reset/confirm", func(c *gin.Context) { gHandler(c, ctrl.ConfirmPasswordReset) })
	//r.POST("/token-master/change-password", func(c *gin.Context) { gHandler(c, ctrl.ChangePassword) })
}
