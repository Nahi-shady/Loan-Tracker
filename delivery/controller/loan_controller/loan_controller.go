package loan_controller

import (
	"loan-tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
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
