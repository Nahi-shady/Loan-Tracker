package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) VerifyEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	token := ctx.Query("token")

	if email == "" || token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and token are required"})
		return
	}

	err := uc.userUsecase.VerifyEmail(ctx.Request.Context(), email, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verified successfully!"})
}
