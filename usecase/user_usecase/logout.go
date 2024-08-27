package user_usecase

import "context"

func (u *UserUsecase) Logout(ctx context.Context, userID string) error {
	// Delete the refresh token associated with the user
	return u.authService.DeleteRefreshToken(ctx, userID)
}
