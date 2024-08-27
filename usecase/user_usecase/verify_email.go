package user_usecase

import (
	"context"
	"errors"
	"time"
)

func (u *UserUsecase) VerifyEmail(ctx context.Context, email, token string) error {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("invalid token or email")
	}

	if user.VerificationToken != token || time.Now().After(user.TokenExpiry) {
		return errors.New("invalid or expired token")
	}

	user.Active = true
	user.VerificationToken = ""
	user.TokenExpiry = time.Time{} // Clear the token expiry

	err = u.userRepo.UpdateUser(ctx, user.ID, user)
	if err != nil {
		return errors.New("failed to update user status")
	}

	return nil
}
