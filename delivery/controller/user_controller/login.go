package user_controller

import (
	"context"
	"loan-tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Login(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.Login(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
