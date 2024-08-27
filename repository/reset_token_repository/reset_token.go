package reset_token_repository

import (
	"loan-tracker/domain"
	"loan-tracker/mongo"
)

type resetTokenRepository struct {
	collection mongo.Collection
}

func NewResetTokenRepository(collection mongo.Collection) domain.ResetTokenRepository {
	return &resetTokenRepository{
		collection: collection,
	}
}
