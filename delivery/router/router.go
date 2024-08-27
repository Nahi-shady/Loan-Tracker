package router

import (
	"loan-tracker/delivery/controller/loan_controller"
	"loan-tracker/delivery/controller/user_controller"
	"loan-tracker/infrastructure/auth"
	"loan-tracker/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, lc *loan_controller.LoanController, env *bootstrap.Env) {
	a := router.Group("/admin")
	a.Use()
	{
		a.GET("/users/:id", uc.Profile)
		a.DELETE("/users/:id", uc.Deleteuser)
		a.GET("/loans", lc.GetAllLoans)
	}

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

	l := router.Group("/loans")
	l.Use()
	{
		l.POST("", auth.JwtAuthMiddleware(env.AccessTokenSecret), lc.ApplyForLoan)
		a.GET("/loans", lc.GetLoanStatus)
	}

}
