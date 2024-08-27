package user_usecase

import (
	"context"
	"loan-tracker/domain"
	"strings"
)

func (uc *UserUsecase) GetByUsernameOrEmail(ctx context.Context, identifier string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	var user domain.User
	var err error
	if strings.Contains(identifier, "@") {
		user, err = uc.userRepo.GetByEmail(ctx, identifier)
	} else {
		user, err = uc.userRepo.GetByUsername(ctx, identifier)
	}

	return user, err
}
