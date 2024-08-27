package user_repository

import (
	"context"
	"loan-tracker/domain"
)

func (r *UserRepository) Signup(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
