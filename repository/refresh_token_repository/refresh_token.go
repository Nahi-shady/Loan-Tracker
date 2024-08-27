package refresh_token_repository

import "loan-tracker/mongo"

type refreshTokenRepository struct {
	collection mongo.Collection
}

func NewRefreshTokenRepository(collection mongo.Collection) *refreshTokenRepository {
	return &refreshTokenRepository{
		collection: collection,
	}
}
