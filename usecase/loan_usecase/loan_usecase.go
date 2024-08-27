package loan_usecase

import (
	"context"
	"time"

	"loan-tracker/domain"
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
