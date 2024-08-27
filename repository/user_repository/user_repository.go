package user_repository

import (
	"loan-tracker/domain"
	"loan-tracker/mongo"
)

type UserRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) domain.UserRepository {
	return &UserRepository{
		collection: collection,
	}
}
