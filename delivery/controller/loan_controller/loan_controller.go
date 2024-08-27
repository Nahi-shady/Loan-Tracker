package loan_controller

import (
	"loan-tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
	loanUsecase domain.LoanUsecase
}

func NewLoanController(loanUsecase domain.LoanUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
	}
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
	status := ctx.DefaultQuery("status", "all")
	order := ctx.DefaultQuery("order", "asc")

	loans, err := c.loanUsecase.GetAllLoans(ctx, status, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans"})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) UpdateLoanStatus(ctx *gin.Context) {
	loanIDStr := ctx.Param("id")
	status := ctx.PostForm("status")

	loanID, err := primitive.ObjectIDFromHex(loanIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	if status != "approved" && status != "rejected" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	err = c.loanUsecase.UpdateLoanStatus(ctx, loanID, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Loan status updated successfully", "status": status})
}
