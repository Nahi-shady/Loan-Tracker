package loan_usecase

import (
	"context"
	"errors"
	"time"

	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUsecase struct {
	loanRepo       domain.LoanRepository
	contextTimeout time.Duration
}

func NewLoanUsecase(loanRepo domain.LoanRepository, timeout time.Duration) domain.LoanUsecase {
	return &LoanUsecase{
		loanRepo:       loanRepo,
		contextTimeout: timeout,
	}
}

func (u *LoanUsecase) GetAllLoans(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	loans, err := u.loanRepo.GetAll(ctx, status, order)
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (u *LoanUsecase) UpdateLoanStatus(ctx context.Context, loanID primitive.ObjectID, status string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// Ensure the status is either "approved" or "rejected"
	if status != "approved" && status != "rejected" {
		return errors.New("invalid status, must be 'approved' or 'rejected'")
	}

	// Update loan status in the repository
	err := u.loanRepo.UpdateLoanStatus(ctx, loanID, status)
	if err != nil {
		return err
	}

	return nil
}
