package router

import (
	"loan-tracker/delivery/controller/user_controller"
	"loan-tracker/infrastructure/auth"
	"loan-tracker/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, env *bootstrap.Env) {
	router.GET("admin/users/{id}", uc.Profile)
	router.DELETE("admin/users/{id}", uc.Deleteuser)
	r := router.Group("/users")
	r.Use()
	{
		r.POST("/register", uc.SignUp)
		r.POST("/login", uc.Login)
		r.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.Logout)
		r.POST("/promote", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.PromoteDemote)
		r.PATCH("/updateUser", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.UpdateUser)
		r.GET("/verify-email", uc.VerifyEmail)
		r.GET("/token/refresh", uc.RefreshTokens)
		r.GET("/profile", uc.Profile)
	}

}
