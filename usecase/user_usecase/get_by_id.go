package user_usecase

import (
	"context"
	"loan-tracker/domain"
)

func (uc *UserUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	// defer cancel()

	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
