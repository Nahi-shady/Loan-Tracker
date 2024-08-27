package loan_repository

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Loan, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var loans []domain.Loan
	if err := cursor.All(ctx, &loans); err != nil {
		return nil, err
	}
	return loans, nil
}
