package user_repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *UserRepository) SetVerificationToken(ctx context.Context, email, token string, expiry time.Time) error {
	filter := bson.M{"email": email}

	update := bson.M{
		"$set": bson.M{"verification_token": token, "token_expiry": expiry},
	}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *UserRepository) VerifyEmailToken(ctx context.Context, email, token string) (bool, error) {
	filter := bson.M{
		"email":              email,
		"verification_token": token,
		"token_expiry":       bson.M{"$gte": time.Now()}}

	update := bson.M{
		"$set": bson.M{"active": true},
		"$unset": bson.M{"verification_token": "",
			"token_expiry": ""}}

	result, err := repo.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return false, err
	}
	if result.ModifiedCount == 0 {
		return false, errors.New("invalid or expired token")
	}
	return true, nil
}
