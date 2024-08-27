package loan_controller

import (
	"loan-tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *LoanController) ApplyForLoan(ctx *gin.Context) {
	var loanRequest domain.LoanRequest

	if err := ctx.ShouldBindJSON(&loanRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user := ctx.GetString("user_id")
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	loan := domain.Loan{
		Amount:   loanRequest.Amount,
		Term:     loanRequest.Term,
		Interest: 5.0,
	}

	createdLoan, err := c.loanUsecase.ApplyForLoan(ctx, userID, loan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for loan"})
		return
	}

	ctx.JSON(http.StatusOK, createdLoan)
}
