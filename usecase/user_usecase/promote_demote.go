package user_usecase

import (
	"context"
)

func (uc *UserUsecase) PromoteDemote(ctx context.Context, identifier, action string) error {
	// ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	// defer cancel()

	user, err := uc.GetByUsernameOrEmail(ctx, identifier)
	if err != nil {
		return err
	}
	err = uc.userRepo.PromoteDemote(ctx, user.ID, action)

	return err
}
