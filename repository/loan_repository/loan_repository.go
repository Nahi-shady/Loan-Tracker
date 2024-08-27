package loan_repository

import (
	"context"
	"loan-tracker/domain"
	"time"

	"loan-tracker/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	collection mongo.Collection
}

func NewLoanRepository(collection mongo.Collection) domain.LoanRepository {
	return &LoanRepository{
		collection: collection,
	}
}

func (r *LoanRepository) GetAll(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	var loans []domain.Loan

	filter := bson.M{}
	if status != "" && status != "all" {
		filter["status"] = status
	}

	sortOrder := 1 // ascending by default
	if order == "desc" || (order == "" && status != "pending") {
		sortOrder = -1
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: sortOrder}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *LoanRepository) UpdateLoanStatus(ctx context.Context, loanID primitive.ObjectID, status string) error {
	filter := bson.M{"_id": loanID}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
