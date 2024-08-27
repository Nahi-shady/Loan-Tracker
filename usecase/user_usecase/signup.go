package user_usecase

import (
	"context"
	"errors"
	"loan-tracker/domain"
	"loan-tracker/infrastructure/validation"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) SignUp(ctx context.Context, req domain.SignupRequest) (domain.SignupResponse, error) {
	err := validation.ValidateEmail(req.Email)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	if _, err := u.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return domain.SignupResponse{}, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	user := &domain.User{
		ID:                 primitive.NewObjectID(),
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Username:           req.Username,
		Password:           string(hashedPassword),
		Email:              req.Email,
		Bio:                "",
		ProfilePicture:     "",
		ContactInformation: "",
		IsAdmin:            false,
		Active:             true,
		CreatedAt:          time.Now(),
	}
	if err := u.userRepo.Signup(ctx, user); err != nil {
		return domain.SignupResponse{}, err
	}

	accessToken, err := u.authService.GenerateAccessToken(ctx, *user)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	_, err = u.authService.GenerateAndStoreRefreshToken(ctx, *user)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	return domain.SignupResponse{
		AccessToken: accessToken,
	}, nil
}
