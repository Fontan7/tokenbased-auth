package service

import (
	ctrl "tokenbased-auth/controller"

	"github.com/gin-gonic/gin"
)

func Private(rg *gin.RouterGroup) {
	r := rg.Group("/private")

	r.POST("/tokenbased-auth/login", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.UserLogin) })
	r.POST("/tokenbased-auth/logout", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.UserLogout) })
	r.POST("/tokenbased-auth/refresh", parseAndSetToken(), func(c *gin.Context) { gHandler(c, ctrl.RefreshToken) })

	//TODO: Add the rest of the routes andf implementations

	//r.POST("/tokenbased-auth/verify-email", func(c *gin.Context) { gHandler(c, ctrl.VerifyEmail) })
	//r.POST("/tokenbased-auth/verify-phone", func(c *gin.Context) { gHandler(c, ctrl.VerifyPhone) })
	//r.POST("/tokenbased-auth/change-email", func(c *gin.Context) { gHandler(c, ctrl.ChangeEmail) })
	//r.POST("/tokenbased-auth/change-phone", func(c *gin.Context) { gHandler(c, ctrl.ChangePhone) })

	//r.POST("/tokenbased-auth/register", func(c *gin.Context) { gHandler(c, ctrl.Register) })
	//r.POST("/tokenbased-auth/confirm-registration", func(c *gin.Context) { gHandler(c, ctrl.ConfirmRegistration) })

	//r.POST("/tokenbased-auth/password-reset/request", func(c *gin.Context) { gHandler(c, ctrl.RequestPasswordReset) })
	//r.POST("/tokenbased-auth/password-reset/confirm", func(c *gin.Context) { gHandler(c, ctrl.ConfirmPasswordReset) })
	//r.POST("/tokenbased-auth/change-password", func(c *gin.Context) { gHandler(c, ctrl.ChangePassword) })
}
