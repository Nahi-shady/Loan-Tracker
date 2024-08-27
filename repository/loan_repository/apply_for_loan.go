package loan_repository

import (
	"context"
	"loan-tracker/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) ApplyForLoan(ctx context.Context, loan *domain.Loan) error {
	loan.ID = primitive.NewObjectID()
	loan.CreatedAt = time.Now()
	loan.Status = "pending"

	_, err := r.collection.InsertOne(ctx, loan)
	return err
}
