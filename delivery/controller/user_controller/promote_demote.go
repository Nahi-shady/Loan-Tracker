package user_controller

import (
	"context"
	"loan-tracker/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) PromoteDemote(c *gin.Context) {
	var request domain.PromoteDemoteRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.Identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "identifier must be specified (i.e username or email)"})
		return
	}
	var action string
	if request.Action == "promote" {
		action = "promote"
	} else if request.Action == "demote" {
		action = "demote"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action must be specified (i.e promote or demote)"})
		return
	}

	if err := uc.userUsecase.PromoteDemote(context.Background(), request.Identifier, action); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
