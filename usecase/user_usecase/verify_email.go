package user_usecase

import (
	"context"
	"errors"
	"loan-tracker/domain"
	"time"
)

func (u *UserUsecase) VerifyEmail(ctx context.Context, email, token string) (domain.LoginResponse, error) {
	user_, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid token or email")
	}

	if user_.VerificationToken != token || time.Now().After(user_.TokenExpiry) {
		return domain.LoginResponse{}, errors.New("invalid or expired token")
	}

	var user domain.UpdateRequest

	user.Active = true
	user.VerificationToken = ""
	user.TokenExpiry = time.Time{} // Clear the token expiry

	err = u.userRepo.UpdateUser(ctx, user_.ID, user)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to update user status")
	}
	// Generate the access token
	accessToken, err := u.authService.GenerateAccessToken(ctx, user_)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// Generate and store the refresh token
	_, err = u.authService.GenerateAndStoreRefreshToken(ctx, user_)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
