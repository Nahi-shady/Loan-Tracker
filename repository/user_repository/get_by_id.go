package user_repository

import (
	"context"
	"loan-tracker/domain"
	"loan-tracker/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *UserRepository) GetByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error) {
	var user domain.User

	filter := bson.M{"_id": userID}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}
