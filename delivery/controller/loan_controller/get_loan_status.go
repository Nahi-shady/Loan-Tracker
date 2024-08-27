package loan_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *LoanController) GetLoanStatus(ctx *gin.Context) {
	loanIDParam := ctx.Param("id")
	loanID, err := primitive.ObjectIDFromHex(loanIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := c.loanUsecase.GetLoanByID(ctx, loanID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loan status"})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}
