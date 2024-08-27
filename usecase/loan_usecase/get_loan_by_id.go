package loan_usecase

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) GetLoanByID(ctx context.Context, loanID primitive.ObjectID) (domain.Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	loan, err := u.loanRepo.GetLoanByID(ctx, loanID)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}
