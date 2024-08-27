package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserController) Profile(c *gin.Context) {
	user := c.GetString("user_id")
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	resp, err := uc.userUsecase.GetByID(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
