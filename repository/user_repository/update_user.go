package user_repository

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser domain.UpdateRequest) error {
	collection := ur.collection
	filter := bson.M{"_id": userID}
	var update = bson.M{"$set": bson.M{}}
	if updatedUser.Firstname != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["firstname"] = updatedUser.Firstname
	}
	if updatedUser.Lastname != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["lastname"] = updatedUser.Lastname
	}
	if updatedUser.Username != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["username"] = updatedUser.Username
	}
	if updatedUser.Bio != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["bio"] = updatedUser.Bio
	}
	if updatedUser.ProfilePicture != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["profile_picture"] = updatedUser.ProfilePicture
	}
	if updatedUser.ContactInformation != "" {
		updateSet := update["$set"].(bson.M)
		updateSet["contract_information"] = updatedUser.ContactInformation
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
