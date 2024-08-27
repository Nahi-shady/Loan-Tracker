package user_controller

import (
	"context"
	"loan-tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) SignUp(c *gin.Context) {
	var req domain.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.SignUp(context.Background(), req)
	r := domain.SignupResponse{}
	if resp == r {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
