package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}

	_, err := ur.collection.DeleteOne(ctx, filter)

	return err
}
