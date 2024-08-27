package user_usecase

import (
	"time"

	"loan-tracker/domain"
)

type UserUsecase struct {
	userRepo       domain.UserRepository
	authService    domain.AuthService
	emailService   domain.EmailService
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, authService domain.AuthService, emailService domain.EmailService, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		userRepo:       userRepository,
		emailService:   emailService,
		authService:    authService,
		contextTimeout: timeout,
	}
}
