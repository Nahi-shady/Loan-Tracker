package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLoan = "loans"
)

type Loan struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount     float64            `json:"amount" bson:"amount"`
	Term       int                `json:"term" bson:"term"` // Loan term in months
	Interest   float64            `json:"interest" bson:"interest"`
	Status     string             `json:"status" bson:"status"` // e.g., "pending", "approved", "rejected"
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	ApprovedAt *time.Time         `json:"approved_at,omitempty" bson:"approved_at,omitempty"`
	RejectedAt *time.Time         `json:"rejected_at,omitempty" bson:"rejected_at,omitempty"`
}

type LoanUsecase interface {
	ApplyForLoan(ctx context.Context, userID primitive.ObjectID, loan Loan) (Loan, error)
	GetLoanByID(ctx context.Context, loanID primitive.ObjectID) (Loan, error)
	GetAllLoans(ctx context.Context, status string, order string) ([]Loan, error)
}

type LoanRepository interface {
	ApplyForLoan(ctx context.Context, loan *Loan) error
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]Loan, error)
	UpdateStatus(ctx context.Context, loanID primitive.ObjectID, status string, timestamp time.Time) error
	GetLoanByID(ctx context.Context, loanID primitive.ObjectID) (Loan, error)
	GetAll(ctx context.Context, status string, order string) ([]Loan, error)
}

type LoanRequest struct {
	Amount float64 `json:"amount" binding:"required"`
	Term   int     `json:"term" binding:"required"`
}
