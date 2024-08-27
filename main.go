package main

import (
	"fmt"
	"loan-tracker/delivery/controller/loan_controller"
	"loan-tracker/delivery/controller/user_controller"
	"loan-tracker/delivery/router"

	"loan-tracker/infrastructure/auth"
	"loan-tracker/infrastructure/bootstrap"

	"loan-tracker/infrastructure/email"
	"loan-tracker/repository/loan_repository"
	"loan-tracker/repository/refresh_token_repository"
	"loan-tracker/repository/reset_token_repository"
	"loan-tracker/repository/user_repository"
	"loan-tracker/usecase/loan_usecase"
	"loan-tracker/usecase/user_usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	userCollection := db.Collection("users")
	refreshTokenCollection := db.Collection("refresh-tokens")
	resetTokenCollection := db.Collection("reset-tokens")
	loanCollection := db.Collection("loans")

	userRepo := user_repository.NewUserRepository(userCollection)
	refreshTokenRepo := refresh_token_repository.NewRefreshTokenRepository(refreshTokenCollection)
	resetTokenRepo := reset_token_repository.NewResetTokenRepository(resetTokenCollection)
	loanRepo := loan_repository.NewLoanRepository(loanCollection)

	authService := auth.NewAuthService(refreshTokenRepo, resetTokenRepo, env.AccessTokenSecret, env.RefreshTokenSecret, env.ResetTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.ResetTokenExpiryHour)

	emailService := email.NewEmailService(env.SMTPServer, env.SMTPPort, env.SMTPUser, env.SMTPPassword, env.FromAddress)

	userUsecase := user_usecase.NewUserUsecase(userRepo, authService, emailService, time.Duration(env.ContextTimeout))
	loanUsecase := loan_usecase.NewLoanUsecase(loanRepo, time.Duration(env.ContextTimeout))

	userController := user_controller.NewUserController(userUsecase, authService, env)
	loanController := loan_controller.NewLoanController(loanUsecase)

	r := gin.Default()
	router.SetRouter(r, userController, loanController, env)
	r.Run(env.ServerAddress)

	fmt.Println("Work correctly", env.DBName)
}
