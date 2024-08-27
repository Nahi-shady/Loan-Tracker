package router

import (
	"loan-tracker/delivery/controller/user_controller"
	"loan-tracker/infrastructure/auth"
	"loan-tracker/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, env *bootstrap.Env) {
	// User routes
	router.POST("/signup", uc.SignUp)
	router.POST("/login", uc.Login)
	router.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.Logout)
	router.POST("/promote", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.PromoteDemote)
	router.PUT("/updateUser", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.UpdateUser)

	// router.POST("/forgot-password", uc.ForgotPassword)
	// router.POST("/reset-password", uc.ResetPassword)
	// router.POST("/refresh", uc.RefreshTokens)

}
