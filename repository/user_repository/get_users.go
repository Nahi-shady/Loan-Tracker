package user_repository

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *UserRepository) GetUsers(ctx context.Context, email string) ([]domain.User, error) {
	var users []domain.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
