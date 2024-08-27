package router

import (
	"loan-tracker/delivery/controller/user_controller"
	"loan-tracker/infrastructure/auth"
	"loan-tracker/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, env *bootstrap.Env) {
	r := router.Group("/users")
	r.Use()
	{
		router.POST("/register", uc.SignUp)
		router.POST("/login", uc.Login)
		router.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.Logout)
		router.POST("/promote", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.PromoteDemote)
		router.PATCH("/updateUser", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.UpdateUser)

	}

}
