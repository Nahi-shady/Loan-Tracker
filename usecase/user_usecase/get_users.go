package user_usecase

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserUsecase) GetByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	// defer cancel()

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
