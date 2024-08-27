package user_usecase

import (
	"context"
	"errors"
	"loan-tracker/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Login(ctx context.Context, req domain.LoginRequest) (domain.LoginResponse, error) {
	// Find user by username or email
	user, err := u.GetByUsernameOrEmail(ctx, req.Identifier)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return domain.LoginResponse{}, errors.New("wrong password")
	}

	// Generate the access token
	accessToken, err := u.authService.GenerateAccessToken(ctx, user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// Generate and store the refresh token
	_, err = u.authService.GenerateAndStoreRefreshToken(ctx, user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
