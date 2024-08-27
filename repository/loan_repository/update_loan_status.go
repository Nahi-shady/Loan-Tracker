package loan_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) UpdateStatus(ctx context.Context, loanID primitive.ObjectID, status string, timestamp time.Time) error {
	filter := bson.M{"_id": loanID}
	update := bson.M{"$set": bson.M{"status": status}}

	if status == "approved" {
		update["$set"].(bson.M)["approved_at"] = timestamp
	} else if status == "rejected" {
		update["$set"].(bson.M)["rejected_at"] = timestamp
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
