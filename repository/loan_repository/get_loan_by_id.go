package loan_repository

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) GetLoanByID(ctx context.Context, loanID primitive.ObjectID) (domain.Loan, error) {
	var loan domain.Loan
	filter := bson.M{"_id": loanID}
	err := r.collection.FindOne(ctx, filter).Decode(&loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}
