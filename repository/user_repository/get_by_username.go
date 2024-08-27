package user_repository

import (
	"context"
	"loan-tracker/domain"
	"loan-tracker/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, err
		}
		return domain.User{}, err
	}
	return user, nil
}
