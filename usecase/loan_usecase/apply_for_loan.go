package loan_usecase

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) ApplyForLoan(ctx context.Context, userID primitive.ObjectID, loan domain.Loan) (domain.Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	loan.UserID = userID
	err := u.loanRepo.ApplyForLoan(ctx, &loan)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}
