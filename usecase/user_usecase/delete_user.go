package user_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserUsecase) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	// ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	// defer cancel()

	err := u.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return errors.New("couldn't delete user, process is canceled")
	}
	return nil
}
